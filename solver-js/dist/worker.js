import { scramble as wasmScramble } from "./pkg/solver_wasm";
onmessage = function (e) {
    var _a = e.data, alg = _a.alg, moves = _a.moves, onlyOrientation = _a.onlyOrientation, disregard = _a.disregard;
    if (typeof alg !== "string") {
        throw new Error("'alg' is not a string");
    }
    if (typeof moves !== "string") {
        throw new Error("'moves' is not a string");
    }
    if (!Array.isArray(onlyOrientation)) {
        throw new Error("'onlyOrientation' is not an array");
    }
    if (!Array.isArray(disregard)) {
        throw new Error("'disregard' is not an array");
    }
    var result = wasmScramble(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard));
    postMessage(result);
};
