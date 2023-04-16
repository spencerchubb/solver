/* tslint:disable */
/* eslint-disable */
/**
* @param {string} alg
* @param {string} moves
* @param {Uint32Array} only_orientation
* @param {Uint32Array} disregard
* @param {number} max_scrambles
* @returns {string}
*/
export function scramble(alg: string, moves: string, only_orientation: Uint32Array, disregard: Uint32Array, max_scrambles: number): string;
/**
* @param {string} alg
* @param {string} moves
* @param {Uint32Array} only_orientation
* @param {Uint32Array} disregard
* @param {number} max_solutions
* @returns {string}
*/
export function solve(alg: string, moves: string, only_orientation: Uint32Array, disregard: Uint32Array, max_solutions: number): string;

export type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

export interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly scramble: (a: number, b: number, c: number, d: number, e: number, f: number, g: number, h: number, i: number, j: number) => void;
  readonly solve: (a: number, b: number, c: number, d: number, e: number, f: number, g: number, h: number, i: number, j: number) => void;
  readonly __wbindgen_add_to_stack_pointer: (a: number) => number;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_free: (a: number, b: number) => void;
}

export type SyncInitInput = BufferSource | WebAssembly.Module;
/**
* Instantiates the given `module`, which can either be bytes or
* a precompiled `WebAssembly.Module`.
*
* @param {SyncInitInput} module
*
* @returns {InitOutput}
*/
export function initSync(module: SyncInitInput): InitOutput;

/**
* If `module_or_path` is {RequestInfo} or {URL}, makes a request and
* for everything else, calls `WebAssembly.instantiate` directly.
*
* @param {InitInput | Promise<InitInput>} module_or_path
*
* @returns {Promise<InitOutput>}
*/
export default function init (module_or_path?: InitInput | Promise<InitInput>): Promise<InitOutput>;
