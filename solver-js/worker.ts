import { scramble as wasmScramble } from "./pkg/solver_wasm";

type MessageData = {
    alg: string,
    moves: string,
    onlyOrientation: number[],
    disregard: number[],
};

onmessage = function (e) {
    const { alg, moves, onlyOrientation, disregard } = (e.data as MessageData);

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

    const result = wasmScramble(
        alg,
        moves,
        new Uint32Array(onlyOrientation),
        new Uint32Array(disregard),
    );

    postMessage(result);
}