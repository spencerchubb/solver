type ScrambleOptions = {
    alg: string;
    moves: string;
    onlyOrientation?: number[];
    disregard?: number[];
    maxSolutions?: number;
};

function scrambleOptionsDefaults(opts: ScrambleOptions): ScrambleOptions {
    if (opts.alg === null || opts.alg === undefined) {
        throw new Error("alg must be defined");
    }
    if (opts.moves === null || opts.moves === undefined) {
        throw new Error("moves must be defined");
    }
    return {
        alg: opts.alg,
        moves: opts.moves,
        onlyOrientation: opts.onlyOrientation ?? [],
        disregard: opts.disregard ?? [],
        maxSolutions: opts.maxSolutions ?? 10,
    };
}

export function scramble(opts: ScrambleOptions): Promise<string[]> {
    const optsWithDefaults = scrambleOptionsDefaults(opts);
    
    return new Promise(resolve => {
        const worker = new Worker(new URL("./worker", import.meta.url));
        worker.onmessage = (event) => {
            // data should be a string of comma-separated scrambles
            const data: string = event.data;
            const scrambles = data.split(",");
            resolve(scrambles);
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
