"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
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
    const runMIPS = vscode.commands.registerCommand("extension.runMIPS", (asmPath) => {
        const editor = vscode.window.activeTextEditor;
        if (!editor) {
            vscode.window.showErrorMessage("No active editor!");
            return;
        }
        let asmFile = asmPath;
        if (!asmFile) {
            const editor = vscode.window.activeTextEditor;
            if (!editor) {
                vscode.window.showErrorMessage("No active editor!");
                return;
            }
            asmFile = editor.document.fileName;
        }
        if (!asmFile.endsWith(".asm")) {
            vscode.window.showErrorMessage("The file must be .asm for executing in MARS");
            return;
        }
        const marsPath = path.join(context.extensionPath, "lib", "Mars4_5.jar");
        const cmd = `java -jar "${marsPath}" nc "${asmFile}"`;
        let terminal = vscode.window.terminals.find(t => t.name === "MIPS Runner");
        if (!terminal) {
            terminal = vscode.window.createTerminal("MIPS Runner");
        }
        terminal.show(true);
        terminal.sendText(cmd);
    });
    context.subscriptions.push(runMIPS);
    const disposable = vscode.commands.registerCommand("extension.runAnalyzer", () => __awaiter(this, void 0, void 0, function* () {
        const editor = vscode.window.activeTextEditor;
        if (!editor) {
            vscode.window.showErrorMessage("No active editor!");
            return;
        }
        const cpsFile = editor.document.fileName;
        const cpsDir = path.dirname(cpsFile);
        const baseName = path.basename(cpsFile, path.extname(cpsFile));
        const analyzePath = path.join(__dirname, "..", "bin", "compiscript-analyzer");
        let terminal = vscode.window.terminals.find(t => t.name === "Compiscript Analyzer");
        if (!terminal) {
            terminal = vscode.window.createTerminal("Compiscript Analyzer");
        }
        terminal.show();
        terminal.sendText(`"${analyzePath}" "${cpsFile}"`);
        yield new Promise(res => setTimeout(res, 500));
        const asmFile = path.join(cpsDir, "out.asm");
        if (!fs.existsSync(asmFile)) {
            vscode.window.showErrorMessage("The analyzer did not generate the .asm file.");
            return;
        }
        yield vscode.commands.executeCommand("extension.runMIPS", asmFile);
    }));
    context.subscriptions.push(disposable);
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
            const cleanLine = line.replace(/\x1b\[[0-9;]*m/g, "");
            const semanticMatch = cleanLine.match(/^\s*\*?\s*Error: \(line: (\d+), column: (\d+)-(\d+)\)\s+(.+)/);
            const syntaxMatch = line.match(/^\s*line (\d+):(\d+)\s+(.+)$/);
            if (semanticMatch) {
                const lineNum = int_to_str(semanticMatch[1], 10) - 1;
                const colStart = int_to_str(semanticMatch[2], 10);
                const colEnd = int_to_str(semanticMatch[3], 10);
                const message = semanticMatch[4];
                if (lineNum >= 0 && lineNum < document.lineCount) {
                    const textLine = document.lineAt(lineNum);
                    const range = new vscode.Range(lineNum, Math.max(colStart, 0), lineNum, Math.min(colEnd, textLine.text.length));
                    diags.push(new vscode.Diagnostic(range, message, vscode.DiagnosticSeverity.Error));
                }
            }
            else if (syntaxMatch) {
                const lineNum = int_to_str(syntaxMatch[1], 10) - 1;
                const colNum = int_to_str(syntaxMatch[2], 10) - 1;
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