import * as vscode from "vscode";
import { execFile } from "child_process";
import * as path from "path";
import * as fs from "fs";
import * as os from "os";

let timeout : NodeJS.Timeout | undefined;

export function activate(context: vscode.ExtensionContext) {
    console.log("¡Compiscript extension activated!");

    const disposable = vscode.commands.registerCommand("extension.runAnalyzer", () => {
        const editor = vscode.window.activeTextEditor;
        if (!editor) {
            vscode.window.showErrorMessage("No active editor!");
            return;
        }

        const file = editor.document.fileName;

        const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");

        let terminal = vscode.window.terminals.find(t => t.name === "Compiscript Analyzer");
        if (!terminal) {
            terminal = vscode.window.createTerminal("Compiscript Analyzer");
        }
        terminal.show();

        terminal.sendText(`"${analyzePath}" "${file}"`);
    })

    context.subscriptions.push(disposable);

    const diagnostic = vscode.languages.createDiagnosticCollection("compiscript");
    context.subscriptions.push(diagnostic);

    const activeDoc = vscode.window.activeTextEditor?.document;
    if (activeDoc && activeDoc.languageId === "compiscript") {
        runSemanticAnalyzer(activeDoc, diagnostic);
    }

    vscode.workspace.onDidOpenTextDocument((document: vscode.TextDocument) => {
        if (document.languageId === "compiscript") {
            runSemanticAnalyzer(document, diagnostic);
        }
    });

    vscode.workspace.onDidChangeTextDocument((event: vscode.TextDocumentChangeEvent) => {
        if (event.document.languageId === "compiscript") {
            runWithDebounce(event.document, diagnostic);
        }
    });
}

function runWithDebounce(document: vscode.TextDocument, diagnostics: vscode.DiagnosticCollection) {
    if (timeout) {
        clearTimeout(timeout);
    }
    timeout = setTimeout(() => {
        runSemanticAnalyzer(document, diagnostics);
    }, 500);
}

function runSemanticAnalyzer(
    document: vscode.TextDocument,
    diagnostics: vscode.DiagnosticCollection
) {
    const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");

    const tmpFile = path.join(os.tmpdir(), `compiscript_${Date.now()}.cps`);
    fs.writeFileSync(tmpFile, document.getText(), "utf8");

    execFile(analyzePath, [tmpFile], (error, stdout, stderr) => {
        const diags: vscode.Diagnostic[] = [];
        const output = stdout + "\n" + stderr;
        const lines = output.split("\n");

        for (const line of lines) {
            const cleanLine = line.replace(/\x1b\[[0-9;]*m/g, "");
            const semanticMatch = cleanLine.match(/^\s*\*?\s*Error: \(line: (\d+), column: (\d+)-(\d+)\)\s+(.+)/);            
            const syntaxMatch = line.match(/^\s*line (\d+):(\d+)\s+(.+)$/);

            if (semanticMatch) {
                const lineNum = parseInt(semanticMatch[1], 10) - 1;
                const colStart = parseInt(semanticMatch[2], 10);
                const colEnd = parseInt(semanticMatch[3], 10);
                const message = semanticMatch[4];
                if (lineNum >= 0 && lineNum < document.lineCount) {
                    const textLine = document.lineAt(lineNum);
                    const range = new vscode.Range(
                        lineNum, 
                        Math.max(colStart, 0), 
                        lineNum, 
                        Math.min(colEnd, textLine.text.length)
                    );
                    diags.push(new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error));
                }
            } else if (syntaxMatch) {
                const lineNum = parseInt(syntaxMatch[1], 10) - 1;
                const colNum = parseInt(syntaxMatch[2], 10) - 1;
                const message = syntaxMatch[3];
                if (lineNum >= 0 && lineNum < document.lineCount) {
                    const textLine = document.lineAt(lineNum);
                    const range = new vscode.Range(
                        lineNum,
                        Math.max(colNum, 0),
                        lineNum,
                        Math.min(colNum + 1, textLine.text.length)
                    );
                    diags.push(new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error));
                }
            }
        }

        diagnostics.set(document.uri, diags);

        if (error) {
            console.error("Error executing the analyzer:", error);
        }

        fs.unlink(tmpFile, (err) => {
            if (err) {
                console.warn("Couldn't delete the temporal file:", tmpFile);
            }
        });
    });
}

export function deactivate() {}