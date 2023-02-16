import { scramble as wasmScramble } from "@spencerchubb/solver";

const scramble = wasmScramble("R U R' F' R U R' U' R' F R2 U' R' U'");
console.log({ scramble });
