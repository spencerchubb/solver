function scrambleOptionsDefaults(opts) {
    var _a, _b, _c;
    if (opts.alg === null || opts.alg === undefined) {
        throw new Error("alg must be defined");
    }
    if (opts.moves === null || opts.moves === undefined) {
        throw new Error("moves must be defined");
    }
    return {
        alg: opts.alg,
        moves: opts.moves,
        onlyOrientation: (_a = opts.onlyOrientation) !== null && _a !== void 0 ? _a : [],
        disregard: (_b = opts.disregard) !== null && _b !== void 0 ? _b : [],
        maxSolutions: (_c = opts.maxSolutions) !== null && _c !== void 0 ? _c : 10
    };
}
export function scramble(opts) {
    var optsWithDefaults = scrambleOptionsDefaults(opts);
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
        worker.postMessage(optsWithDefaults);
    });
}
