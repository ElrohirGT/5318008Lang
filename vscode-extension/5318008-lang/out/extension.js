"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = require("vscode");
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
    const diags = [];
    // Diagnóstico de prueba: línea 0, columna 0
    const range = new vscode.Range(0, 0, 0, 5); // desde la columna 0 hasta 5
    diags.push(new vscode.Diagnostic(range, "¡Prueba: esto es un error de ejemplo!", vscode.DiagnosticSeverity.Error));
    // Actualiza la colección de diagnósticos
    diagnostics.set(document.uri, diags);
}
function deactivate() { }
//# sourceMappingURL=extension.js.map