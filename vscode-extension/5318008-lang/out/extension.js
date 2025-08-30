"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = require("vscode");
const child_process_1 = require("child_process");
const path = require("path");
function activate(context) {
    console.log("¡Compiscript extension activated!");
    const diagnostic = vscode.languages.createDiagnosticCollection("compiscript");
    context.subscriptions.push(diagnostic);
    vscode.workspace.onDidSaveTextDocument((document) => {
        if (document.languageId == "compiscript") {
            runSemanticAnalyzer(document, diagnostic);
        }
    });
}
function runSemanticAnalyzer(document, diagnostics) {
    const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");
    (0, child_process_1.execFile)(analyzePath, [document.fileName], (error, stdout, stderr) => {
        if (error) {
            console.error("Error executing the analyzer:", error);
            return;
        }
        const diags = [];
        const lines = stdout.split("\n");
        for (const line of lines) {
            const match = line.match(/(.+):(\d+):(\d+): (.+)/);
            if (match) {
                const [, , lineStr, colStr, message] = match;
                const lineNum = parseInt(lineStr) - 1;
                const colNum = parseInt(colStr) - 1;
                const range = new vscode.Range(lineNum, colNum, lineNum, colNum + 1);
                diags.push(new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error));
            }
        }
        diagnostics.set(document.uri, diags);
    });
}
function deactivate() { }
//# sourceMappingURL=extension.js.map