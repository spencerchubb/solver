use std::time::{SystemTime, UNIX_EPOCH};

use crate::{
    algorithm::{Algorithm, different_face_or_same_move, same_face},
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
    let end = Cube::new();
    start.perform_alg_string("R U R' F' R U R' U' R' F R2 U' R' U'");

    // let moves = Moves::from_string("U,U2,U',F,F2,F',R,R2,R'");
    let moves = Moves::from_string("U,U',F,F',R,R'");

    // let next_move_valid = |alg: &Algorithm, mooove: u8| -> bool {
    //     !same_face(alg, mooove)
    // };
    let next_move_valid = |alg: &Algorithm, mooove: u8| -> bool {
        different_face_or_same_move(alg, mooove)
    };

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
