import { scramble as wasmScramble } from "./pkg/solver_wasm";

export function scramble(
    alg: string,
    moves: string,
    onlyOrientation: number[],
    disregard: number[],
): Promise<string> {
    return new Promise(resolve => {
        execNonBlocking(() => {
            const result = wasmScramble(
                alg,
                moves,
                new Uint32Array(onlyOrientation),
                new Uint32Array(disregard),
            );
            resolve(result);
        }); 
    });
}

function execNonBlocking(func: () => void): void {
    setTimeout(func, 0);
}
