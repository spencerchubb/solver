import { scramble as wasmScramble } from "./pkg/solver_wasm";
export function scramble(alg, moves, onlyOrientation, disregard) {
    return new Promise(function (resolve) {
        execNonBlocking(function () {
            var result = wasmScramble(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard));
            resolve(result);
        });
    });
}
function execNonBlocking(func) {
    setTimeout(func, 0);
}
