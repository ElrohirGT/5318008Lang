"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.activate = activate;
exports.deactivate = deactivate;
const vscode = require("vscode");
const child_process_1 = require("child_process");
const path = require("path");
const fs = require("fs");
const os = require("os");
let timeout;
function activate(context) {
    var _a;
    console.log("¡Compiscript extension activated!");
    const diagnostic = vscode.languages.createDiagnosticCollection("compiscript");
    context.subscriptions.push(diagnostic);
    const activeDoc = (_a = vscode.window.activeTextEditor) === null || _a === void 0 ? void 0 : _a.document;
    if (activeDoc && activeDoc.languageId === "compiscript") {
        runSemanticAnalyzer(activeDoc, diagnostic);
    }
    vscode.workspace.onDidOpenTextDocument((document) => {
        if (document.languageId === "compiscript") {
            runSemanticAnalyzer(document, diagnostic);
        }
    });
    vscode.workspace.onDidChangeTextDocument((event) => {
        if (event.document.languageId === "compiscript") {
            runWithDebounce(event.document, diagnostic);
        }
    });
}
function runWithDebounce(document, diagnostics) {
    if (timeout) {
        clearTimeout(timeout);
    }
    timeout = setTimeout(() => {
        runSemanticAnalyzer(document, diagnostics);
    }, 500);
}
function runSemanticAnalyzer(document, diagnostics) {
    const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");
    const tmpFile = path.join(os.tmpdir(), `compiscript_${Date.now()}.cps`);
    fs.writeFileSync(tmpFile, document.getText(), "utf8");
    (0, child_process_1.execFile)(analyzePath, [tmpFile], (error, stdout, stderr) => {
        const diags = [];
        const output = stdout + "\n" + stderr;
        const lines = output.split("\n");
        for (const line of lines) {
            const semanticMatch = line.match(/^\s*\*?\s*Error: \(line: (\d+), column: (\d+)-(\d+)\)\s+(.+)/);
            const syntaxMatch = line.match(/^\s*line (\d+):(\d+)\s+(.+)$/);
            if (semanticMatch) {
                const lineNum = parseInt(semanticMatch[1], 10) - 1;
                const colStart = parseInt(semanticMatch[2], 10);
                const colEnd = parseInt(semanticMatch[3], 10);
                const message = semanticMatch[4];
                if (lineNum >= 0 && lineNum < document.lineCount) {
                    const textLine = document.lineAt(lineNum);
                    const range = new vscode.Range(lineNum, Math.max(colStart, 0), lineNum, Math.min(colEnd, textLine.text.length));
                    diags.push(new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error));
                }
            }
            else if (syntaxMatch) {
                const lineNum = parseInt(syntaxMatch[1], 10) - 1;
                const colNum = parseInt(syntaxMatch[2], 10) - 1;
                const message = syntaxMatch[3];
                if (lineNum >= 0 && lineNum < document.lineCount) {
                    const textLine = document.lineAt(lineNum);
                    const range = new vscode.Range(lineNum, Math.max(colNum, 0), lineNum, Math.min(colNum + 1, textLine.text.length));
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
function deactivate() { }
//# sourceMappingURL=extension.js.map