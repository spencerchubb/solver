export function scramble(
    alg: string,
    moves: string,
    onlyOrientation: number[],
    disregard: number[],
): Promise<string> {
    return new Promise((resolve, reject) => {
        const worker = new Worker("./worker.js");
        worker.onmessage = (event) => {
            resolve(event.data);
        };
        worker.onerror = (event) => {
            reject(event);
        };
        worker.postMessage({
            alg,
            moves,
            onlyOrientation,
            disregard,
        });
    });
}
