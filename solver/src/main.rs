use std::time::{SystemTime, UNIX_EPOCH};

use crate::scramble::ScrambleArgs;

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
    // let alg = "R U R' F' R U R' U' R' F R U' R' F R F'";
    let alg = "F R U R' U' R U' R' U' R U R' F'";

    let moves = "U U' F F' R R'";
    let only_orientation: &[usize] = &[];
    let disregard = &[];
    
    // OLL case
    // let alg = "R U2 R2 F R F' U2 R' F R F'";
    // let only_orientation = &[0, 1, 2, 3, 8, 9, 10, 11];

    let max_scrambles = 50;

    let iterations = 10;
    let mut total_time = 0;
    for i in 0..iterations {
        let start_time = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis();

        scramble::scramble(ScrambleArgs {
            alg: alg.to_string(),
            moves: moves.to_string(),
            only_orientation: only_orientation.to_vec(),
            disregard: disregard.to_vec(),
            max_scrambles,
            // solution_found: |s: String| println!("{}", s),
            solution_found: |_: String| {},
        });

        let elapsed = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .unwrap()
            .as_millis()
            - start_time;
        total_time += elapsed;
        println!("Iteration {}: {} ms", i, elapsed)
    }
    println!("Average: {} ms", total_time / iterations);
}
