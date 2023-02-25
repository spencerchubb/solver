import { scramble } from "./dist/index";

const alg = "F R U' R' U' R U R' F' R U R' U' R' F R F'";
const moves = "U,U',F,F',R,R'";
const onlyOrientation: number[] = [];
const disregard: number[] = [];

const root = document.getElementById("root");

scramble(alg, moves, onlyOrientation, disregard).then((scrambleFound) => {
    root?.appendChild(createP("scrambleFound: " + scrambleFound));
});

root?.appendChild(createP("searching for scramble - this should appear first because it's async"));

function createP(text: string): HTMLParagraphElement {
    const p = document.createElement("p");
    p.innerHTML = text;
    return p;
}