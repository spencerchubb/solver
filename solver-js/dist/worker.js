import init, { scramble } from "./pkg/solver_wasm";

onmessage = function (e) {
    console.log(e);
    init(new URL("./pkg/solver_wasm_bg.wasm", import.meta.url)).then(() => {
        const { alg, moves, onlyOrientation, disregard } = e.data;
        const result = scramble(alg, moves, new Uint32Array(onlyOrientation), new Uint32Array(disregard));
        postMessage(result);
    });
}
