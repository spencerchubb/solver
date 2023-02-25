export function scramble(alg, moves, onlyOrientation, disregard) {
    return new Promise(function (resolve) {
        var worker = new Worker(new URL("./worker", import.meta.url));
        worker.onmessage = function (event) {
            resolve(event.data);
            worker.terminate();
        };
        worker.onerror = function (event) {
            console.log("Caught error");
            console.error(event);
            worker.terminate();
        };
        worker.postMessage({
            alg: alg,
            moves: moves,
            onlyOrientation: onlyOrientation,
            disregard: disregard
        });
    });
}
