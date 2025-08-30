import * as vscode from "vscode";
import { execFile } from "child_process";
import * as path from "path";

export function activate(context: vscode.ExtensionContext) {    
    console.log("¡Compiscript extension activated!");
    const diagnostic = vscode.languages.createDiagnosticCollection("compiscript");
    context.subscriptions.push(diagnostic);

    vscode.workspace.onDidSaveTextDocument((document : vscode.TextDocument) => {
        if (document.languageId == "compiscript") {
            runSemanticAnalyzer(document, diagnostic);
        }
    })
}

function runSemanticAnalyzer(
    document: vscode.TextDocument,
    diagnostics: vscode.DiagnosticCollection
) {
    const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");

    execFile(analyzePath, [document.fileName], (error, stdout, stderr) => {
        const diags: vscode.Diagnostic[] = [];

        const output = stdout + "\n" + stderr;
        const lines = output.split("\n");
        for (const line of lines) {
            let match = line.match(/^\s*\*? ?Error: \(line: (\d+)\) (.+)/);
            if (!match) {
               match = line.match(/^\s*line (\d+):\d+ (.+)$/); 
            }

            if (match) {
                const lineNum = parseInt(match[1], 10) - 1;
                const message = match[2];
                
                console.log("Match found:", match);
                

                const textLine = document.lineAt(lineNum);
                const range = new vscode.Range(lineNum, 0, lineNum, textLine.text.length);
                diags.push(
                    new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error)
                );
            }
        }

        diagnostics.set(document.uri, diags);

        if (error) {
            console.error("Error executing the analyzer:", error);
        }
    });
}

export function deactivate() {}