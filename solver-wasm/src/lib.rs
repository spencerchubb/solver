use solver::{scramble::ScrambleArgs, solve::SolveArgs, cube::Cube, moves::Moves, algorithm::string_to_alg};
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    // Use `js_namespace` here to bind `console.log(..)` instead of just `log(..)`
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[wasm_bindgen]
pub fn scramble(alg: &str, moves: &str, only_orientation: &[usize], disregard: &[usize], max_scrambles: i32) -> String {
    let scrambles = solver::scramble::scramble(ScrambleArgs {
        alg: alg.to_string(),
        moves: moves.to_string(),
        only_orientation: only_orientation.to_vec(),
        disregard: disregard.to_vec(),
        max_scrambles,
        solution_found: |_: String| {},
    });

    let scrambles = scrambles.into_iter().collect::<Vec<String>>();
    scrambles.join(",")
}

#[wasm_bindgen]
pub fn solve(alg: &str, moves: &str, only_orientation: &[usize], disregard: &[usize], max_solutions: i32) -> String {
    let mut start = Cube::new();
    let mut end = Cube::new();
    
    start.set_only_orientation(&only_orientation);
    end.set_only_orientation(&only_orientation);

    start.set_disregard(&disregard);
    end.set_disregard(&disregard);

    let alg = string_to_alg(&alg);
    start.perform_alg(alg);

    let moves = Moves::from_string(moves);

    let next_move_valid = if moves.has_double() {
        solver::algorithm::move_valid_double
    } else {
        solver::algorithm::move_valid_single
    };

    let scrambles = solver::solve::solve(SolveArgs {
        start: end,
        end: start,
        moves: moves,
        move_valid: next_move_valid,
        max_solutions,
        solution_found: |_: String| {},
        log: |_: String| {},
    });

    let scrambles = scrambles.into_iter().collect::<Vec<String>>();
    scrambles.join(",")
}
