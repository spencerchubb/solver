import init, { scramble, solve } from "./pkg/solver_wasm";

onmessage = function (e) {
    const data = e.data;
    const { method, alg, moves, onlyOrientation, disregard, maxSolutions } = e.data;
    init(new URL("./pkg/solver_wasm_bg.wasm", import.meta.url)).then(() => {
        const { alg, moves, onlyOrientation, disregard, maxSolutions } = e.data;
        if (method === "scramble") {
            const result = scramble(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard), maxSolutions);
            postMessage(result);
        } else if (method === "solve") {
            const result = solve(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard), maxSolutions);
            postMessage(result);
        } else {
            console.error("Unknown method: " + e.method);
        }
    });
}
