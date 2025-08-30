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
        if (error) {
            console.error("Error executing the analyzer:", error);
            return;
        }

        const diags: vscode.Diagnostic[] = [];

        const lines = stdout.split("\n");
        for (const line of lines) {
            const match = line.match(/(.+):(\d+):(\d+): (.+)/);
            if (match) {
                const [, , lineStr, colStr, message] = match;
                const lineNum = parseInt(lineStr) - 1;
                const colNum = parseInt(colStr) - 1;

                const range = new vscode.Range(lineNum, colNum, lineNum, colNum + 1);
                diags.push(
                    new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error)
                );
            }
        }

        diagnostics.set(document.uri, diags);
    });
}

export function deactivate() {}