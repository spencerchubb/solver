use crate::{moves::*, cube::Cube, solve::run_solve};

mod arch;
mod cube;
mod moves;
mod node;
mod solve;
mod visited;

fn main() {
    let start_time = std::time::Instant::now();

    let mut start = Cube::new();
    let end = Cube::new();
    start.perform_alg_string("R U R' F' R U R' U' R' F R2 U' R' U'");

    let moves = [U1_NUM, U2_NUM, U3_NUM, F1_NUM, F2_NUM, F3_NUM, R1_NUM, R2_NUM, R3_NUM];
    let max_solutions = 100;
    let max_ms = 10_000;
    let log = true;

    run_solve(start, end, &moves, max_solutions, max_ms, log);

    let elapsed = start_time.elapsed();
    println!("Elapsed: {}.{:09} seconds", elapsed.as_secs(), elapsed.subsec_nanos());
}
