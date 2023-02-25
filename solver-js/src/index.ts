export function scramble(alg: string, moves: string, onlyOrientation: number[], disregard: number[], maxSolutions: number): Promise<string[]> {
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

        worker.postMessage({
            alg,
            moves,
            onlyOrientation,
            disregard,
            maxSolutions,
        });
    });
}
