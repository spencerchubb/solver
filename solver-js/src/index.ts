export function scramble(alg: string, moves: string, onlyOrientation: number[], disregard: number[]): Promise<string> {
    return new Promise(resolve => {
        const worker = new Worker(new URL("./worker", import.meta.url));
        worker.onmessage = (event) => {
            resolve(event.data);
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
        });
    });
}
