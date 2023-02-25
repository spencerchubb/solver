import init, { scramble } from "./pkg/solver_wasm";

onmessage = function (e) {
    init(new URL("./pkg/solver_wasm_bg.wasm", import.meta.url)).then(() => {
        const { alg, moves, onlyOrientation, disregard, maxSolutions } = e.data;
        const result = scramble(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard), maxSolutions);
        postMessage(result);
    });
}
