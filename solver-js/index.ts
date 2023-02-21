import { scramble as wasmScramble } from "./pkg/solver_wasm";

export function scramble(
    alg: string,
    moves: string,
    onlyOrientation: number[],
    disregard: number[],
): string {
    return wasmScramble(
        alg,
        moves,
        new Uint32Array(onlyOrientation),
        new Uint32Array(disregard),
    );
}
