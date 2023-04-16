type SolveOptions = {
    alg: string;
    moves: string;
    onlyOrientation?: number[];
    disregard?: number[];
    maxSolutions?: number;
};
export declare function scramble(opts: SolveOptions): Promise<string[]>;
export declare function solve(opts: SolveOptions): Promise<string[]>;
export {};
//# sourceMappingURL=index.d.ts.map