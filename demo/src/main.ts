import { scramble } from "@spencerchubb/solver";

document.querySelector("#button-generate")?.addEventListener("click", () => {
    const alg = (document.querySelector("#input-alg") as HTMLInputElement).value;
    const moves = (document.querySelector("#input-move-set") as HTMLInputElement).value;
    const onlyOrientation: number[] = [];
    const disregard: number[] = [];
    const maxSolutions = parseInt((document.querySelector("#input-max-solutions") as HTMLInputElement).value);

    if (!validateMoves(alg)) {
        alert("alg should be separated by spaces, and the valid moves are: U, U', U2, D, D', D2, L, L', L2, R, R', R2, F, F', F2, B, B', B2");
        return;
    }
    if (!validateMoves(moves)) {
        alert("move set should be separated by spaces, and the valid moves are: U, U', U2, D, D', D2, L, L', L2, R, R', R2, F, F', F2, B, B', B2");
        return;
    }

    const root = document.getElementById("root") as HTMLElement;
    root.innerHTML = "";

    const startTime = Date.now();

    const scrambleOpts = {
        alg,
        moves,
        onlyOrientation,
        disregard,
        maxSolutions,
    };
    scramble(scrambleOpts).then((scrambles) => {
        const elapsedTime = Date.now() - startTime;
        root?.appendChild(createP("time: " + (elapsedTime / 1000) + " seconds"));
        scrambles.forEach((scramble, i) => {
            root?.appendChild(createP(`${i + 1}. ${scramble}`));
        });
    });
    root?.appendChild(createP("searching for scrambles..."));
});

function validateMoves(moves: string): boolean {
    const moveSet = moves.split(" ");
    const validMoves = new Set(["U", "U'", "U2", "D", "D'", "D2", "L", "L'", "L2", "R", "R'", "R2", "F", "F'", "F2", "B", "B'", "B2"]);
    return moveSet.every((move) => validMoves.has(move));
}

function createP(text: string): HTMLParagraphElement {
    const p = document.createElement("p");
    p.innerHTML = text;
    return p;
}