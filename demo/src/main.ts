// import { init } from "@spencerchubb/solver";

// init("./pkg/solver_wasm.js").then(wasmObj => {
//     const scramble = wasmObj.scramble("R U R' U' R' F R2 U' R' U' R U R' F'");
//     console.log({ scramble });
// });

import { scramble as wasmScramble } from "solver-wasm";

const scramble = wasmScramble("R U R' F' R U R' U' R' F R2 U' R' U'");
console.log({ scramble });
