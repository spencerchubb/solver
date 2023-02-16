import { scramble as wasmScramble } from "./pkg/solver_wasm";
export function scramble(alg, moves) {
    var scramblesAsStr = wasmScramble(alg, moves);
    var scrambles = scramblesAsStr.split(",");
    return randElement(scrambles);
}
function randElement(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}
