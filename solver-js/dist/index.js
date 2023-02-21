export function scramble(alg, moves, onlyOrientation, disregard) {
    return new Promise(function (resolve, reject) {
        var worker = new Worker("./worker.js");
        worker.onmessage = function (event) {
            resolve(event.data);
        };
        worker.onerror = function (event) {
            reject(event);
        };
        worker.postMessage({
            alg: alg,
            moves: moves,
            onlyOrientation: onlyOrientation,
            disregard: disregard
        });
    });
}
