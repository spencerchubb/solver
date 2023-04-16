type Method = "scramble" | "solve";

type SolveOptions = {
    alg: string;
    moves: string;
    onlyOrientation?: number[];
    disregard?: number[];
    maxSolutions?: number;
};

type WorkerOptions = {
    method: Method;
} & SolveOptions;

function solveOptionsDefaults(method: Method, opts: SolveOptions): WorkerOptions {
    if (opts.alg === null || opts.alg === undefined) {
        throw new Error("alg must be defined");
    }
    if (opts.moves === null || opts.moves === undefined) {
        throw new Error("moves must be defined");
    }
    return {
        method,
        alg: opts.alg,
        moves: opts.moves,
        onlyOrientation: opts.onlyOrientation ?? [],
        disregard: opts.disregard ?? [],
        maxSolutions: opts.maxSolutions ?? 10,
    };
}

export async function scramble(opts: SolveOptions): Promise<string[]> {
    return setupWorker("scramble", opts);
}

export async function solve(opts: SolveOptions): Promise<string[]> {
    return setupWorker("solve", opts);
}

async function setupWorker(method: Method, opts: SolveOptions): Promise<string[]> {
    const optsWithDefaults = solveOptionsDefaults(method, opts);

    return new Promise(resolve => {
        const worker = new Worker(new URL("./worker", import.meta.url));
        worker.onmessage = (event) => {
            // data should be a string of comma-separated solutions
            const data: string = event.data;
            const solutions = data.split(",");
            resolve(solutions);
            worker.terminate();
        }
        worker.onerror = (event) => {
            console.log("Caught error");
            console.error(event);
            worker.terminate();
        }

        worker.postMessage(optsWithDefaults);
    });
}
