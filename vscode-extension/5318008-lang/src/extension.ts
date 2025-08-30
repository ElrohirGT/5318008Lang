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
    const diags: vscode.Diagnostic[] = [];

    // Diagnóstico de prueba: línea 0, columna 0
    const range = new vscode.Range(0, 0, 0, 5); // desde la columna 0 hasta 5
    diags.push(
        new vscode.Diagnostic(range, "¡Prueba: esto es un error de ejemplo!", vscode.DiagnosticSeverity.Error)
    );

    // Actualiza la colección de diagnósticos
    diagnostics.set(document.uri, diags);
}

export function deactivate() {}