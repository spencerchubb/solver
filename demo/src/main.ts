import { scramble } from "@spencerchubb/solver";

const alg = "R U R' F' R U R' U' R' F R2 U' R' U'";
const moves = "U,U',F,F',R,R'";
const onlyOrientation: number[] = [];
const disregard: number[] = [];
const scrambleFound = scramble(alg, moves, onlyOrientation, disregard);
console.log({ scrambleFound });
