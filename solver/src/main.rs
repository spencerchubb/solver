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
    // Example of finding cross solution
    // let alg = "L2 U2 L2 R2 B2 D' R2 U L2 D L2 U' L U B F2 D2 L F' L B'";
    // let moves = "U U' F F' D D' R R' L L' B B'";

    // let mut start = crate::cube::Cube::new();
    // let mut end = crate::cube::Cube::new();

    // let only_orientation: &[usize] = &[];
    // let disregard: &[usize] = &[0,1,2,3,4,5,6,7,12,13,14,15,16,17,18,19];
    
    // start.set_only_orientation(&only_orientation);
    // end.set_only_orientation(&only_orientation);

    // start.set_disregard(&disregard);
    // end.set_disregard(&disregard);

    // let alg = crate::algorithm::string_to_alg(&alg);
    // start.perform_alg(alg);

    // let moves = crate::moves::Moves::from_string(moves);

    // let next_move_valid = if moves.has_double() {
    //     crate::algorithm::move_valid_double
    // } else {
    //     crate::algorithm::move_valid_single
    // };

    // crate::solve::solve(crate::solve::SolveArgs {
    //     start: end,
    //     end: start,
    //     moves: moves,
    //     move_valid: next_move_valid,
    //     max_solutions: 5,
    //     solution_found: |s: String| println!("{}", s),
    //     log: |_: String| {},
    // });

    // Performance testing
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
