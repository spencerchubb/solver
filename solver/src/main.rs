use std::time::{SystemTime, UNIX_EPOCH};

use crate::{
    algorithm::{different_face_or_same_move, same_face, string_to_alg, Algorithm},
    cube::Cube,
    moves::*,
    solve::run_solve,
};

pub mod algorithm;
pub mod constants;
pub mod cube;
pub mod moves;
pub mod node;
pub mod queue;
pub mod solve;
pub mod visited;

pub fn main() {
    let mut start = Cube::new();
    let mut end = Cube::new();

    // start.set_only_orientation(&[0, 1, 2, 3, 8, 9, 10, 11]);
    // end.set_only_orientation(&[0, 1, 2, 3, 8, 9, 10, 11]);

    let alg = "R U R' F' R U R' U' R' F R2 U' R' U'";
    // let alg = "R U R' F' R U R' U' R' F R U' R' F R F'";
    let alg = string_to_alg(alg);
    let alg = invert_algorithm(alg);
    start.perform_alg(alg);

    // let moves = Moves::from_string("U,U2,U',F,F2,F',R,R2,R'");
    let moves = Moves::from_string("U,U',F,F',R,R'");

    let next_move_valid = |alg: &Algorithm, mooove: u8| -> bool { !same_face(alg, mooove) };
    // let next_move_valid = different_face_or_same_move;

    let max_solutions = 10;

    let start_time = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis();

    run_solve(start, end, &moves, next_move_valid, max_solutions);

    let elapsed = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_millis()
        - start_time;
    println!("Elapsed: {} ms", elapsed);
}
