import { scramble as wasmScramble } from "./pkg/solver_wasm";

const scramble = wasmScramble("R U R' F' R U R' U' R' F R2 U' R' U'");
console.log({ scramble });
