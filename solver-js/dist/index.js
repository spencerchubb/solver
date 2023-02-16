import { scramble as wasmScramble } from "./pkg/solver_wasm";
export function scramble(alg) {
    var scramblesAsStr = wasmScramble(alg);
    var scrambles = scramblesAsStr.split(",");
    return randElement(scrambles);
}
function randElement(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}
