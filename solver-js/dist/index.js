var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
function solveOptionsDefaults(method, opts) {
    var _a, _b, _c;
    if (opts.alg === null || opts.alg === undefined) {
        throw new Error("alg must be defined");
    }
    if (opts.moves === null || opts.moves === undefined) {
        throw new Error("moves must be defined");
    }
    return {
        method: method,
        alg: opts.alg,
        moves: opts.moves,
        onlyOrientation: (_a = opts.onlyOrientation) !== null && _a !== void 0 ? _a : [],
        disregard: (_b = opts.disregard) !== null && _b !== void 0 ? _b : [],
        maxSolutions: (_c = opts.maxSolutions) !== null && _c !== void 0 ? _c : 10
    };
}
export function scramble(opts) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            return [2 /*return*/, setupWorker("scramble", opts)];
        });
    });
}
export function solve(opts) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            return [2 /*return*/, setupWorker("solve", opts)];
        });
    });
}
function setupWorker(method, opts) {
    return __awaiter(this, void 0, void 0, function () {
        var optsWithDefaults;
        return __generator(this, function (_a) {
            optsWithDefaults = solveOptionsDefaults(method, opts);
            return [2 /*return*/, new Promise(function (resolve) {
                    var worker = new Worker(new URL("./worker", import.meta.url));
                    worker.onmessage = function (event) {
                        // data should be a string of comma-separated solutions
                        var data = event.data;
                        var solutions = data.split(",");
                        resolve(solutions);
                        worker.terminate();
                    };
                    worker.onerror = function (event) {
                        console.log("Caught error");
                        console.error(event);
                        worker.terminate();
                    };
                    worker.postMessage(optsWithDefaults);
                })];
        });
    });
}
