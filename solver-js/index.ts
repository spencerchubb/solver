let wasm;

type WasmObject = {
    scramble: (alg: string) => string,
};

export async function init(path: string): Promise<WasmObject> {
    wasm = await import(path);
    return {
        scramble,
    }
}

function scramble (alg: string): string {
    let scramblesAsStr: string = wasm.scramble(alg);
    let scrambles: string[] = scramblesAsStr.split(",");
    
    return randElement(scrambles);
}

function randElement<T>(arr: T[]): T {
    return arr[Math.floor(Math.random() * arr.length)];
}
