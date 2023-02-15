use std::time::{SystemTime, UNIX_EPOCH};

use crate::{moves::*, cube::Cube, solve::run_solve};

pub mod algorithm;
pub mod cube;
pub mod moves;
pub mod node;
pub mod solve;
pub mod queue;
pub mod visited;

pub fn main() {
    
    let mut start = Cube::new();
    let end = Cube::new();
    start.perform_alg_string("R U R' F' R U R' U' R' F R2 U' R' U'");
    
    let moves = [U1_NUM, U2_NUM, U3_NUM, F1_NUM, F2_NUM, F3_NUM, R1_NUM, R2_NUM, R3_NUM];
    let max_solutions = 100;
    let log = true;
    
    let start_time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis();
    run_solve(start, end, &moves, max_solutions, log);

    let elapsed = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - start_time;
    println!("Elapsed: {} ms", elapsed);
}
