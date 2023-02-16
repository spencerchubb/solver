use solver::{cube::Cube, moves::*, solve::run_solve};

use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    // Use `js_namespace` here to bind `console.log(..)` instead of just `log(..)`
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[wasm_bindgen]
pub fn scramble(alg: &str, moves: &str) -> String {
    log(&format!("scrambling {}", alg));
    let mut start = Cube::new();
    let end = Cube::new();

    let alg = solver::algorithm::string_to_alg(alg);
    let alg = solver::moves::invert_algorithm(alg);
    start.perform_alg(alg);

    let moves = Moves::from_string(moves);

    run_solve(end, start, &moves, 10, false)
        .into_iter()
        .collect::<Vec<String>>()
        .join(",")
}
