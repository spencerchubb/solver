type ScrambleOptions = {
    alg: string;
    moves: string;
    onlyOrientation?: number[];
    disregard?: number[];
    maxSolutions?: number;
};
export declare function scramble(opts: ScrambleOptions): Promise<string[]>;
export {};
//# sourceMappingURL=index.d.ts.map