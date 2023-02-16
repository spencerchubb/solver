import { scramble as wasmScramble } from "./pkg/solver_wasm";

export function scramble (alg: string, moves: string): string {
    let scramblesAsStr: string = wasmScramble(alg, moves);
    let scrambles: string[] = scramblesAsStr.split(",");
    
    return randElement(scrambles);
}

function randElement<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}
