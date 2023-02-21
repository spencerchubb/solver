use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern {
    // Use `js_namespace` here to bind `console.log(..)` instead of just `log(..)`
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[wasm_bindgen]
pub fn scramble(alg: &str, moves: &str, only_orientation: &[usize], disregard: &[usize]) -> String {
    log(&format!("scrambling {}", alg));

    let max_scrambles = 10;

    let scrambles = solver::scramble::scramble(alg, moves, only_orientation, disregard, max_scrambles);

    let scrambles = scrambles.into_iter().collect::<Vec<String>>();
    let rand_index = rand::random::<usize>() % scrambles.len();
    scrambles[rand_index].to_string()
}
