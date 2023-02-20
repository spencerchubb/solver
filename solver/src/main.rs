use std::time::{SystemTime, UNIX_EPOCH};

pub mod algorithm;
pub mod constants;
pub mod cube;
pub mod moves;
pub mod node;
pub mod queue;
pub mod scramble;
pub mod solve;
pub mod visited;

pub fn main() {
    // let alg = "R U R' F' R U R' U' R' F R2 U' R' U'";
    let alg = "R U R' F' R U R' U' R' F R U' R' F R F'";

    let moves = "U,U',F,F',R,R'";

    // let only_orientation = &[];
    let only_orientation = &[0, 1, 2, 3, 4, 5, 6, 7, 8];

    let max_solutions = 10;

    let start_time = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis();

    scramble::scramble(alg, &moves, only_orientation, max_solutions);

    let elapsed = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis()
        - start_time;
    println!("Elapsed: {} ms", elapsed);
}
