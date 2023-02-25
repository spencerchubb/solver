export function scramble(alg, moves, onlyOrientation, disregard, maxSolutions) {
    return new Promise(function (resolve) {
        var worker = new Worker(new URL("./worker", import.meta.url));
        worker.onmessage = function (event) {
            // data should be a string of comma-separated scrambles
            var data = event.data;
            var scrambles = data.split(",");
            resolve(scrambles);
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
            disregard: disregard,
            maxSolutions: maxSolutions
        });
    });
}
