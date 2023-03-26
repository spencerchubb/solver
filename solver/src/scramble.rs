use std::collections::HashSet;

use crate::{
    algorithm::string_to_alg,
    cube::Cube,
    moves::{invert_algorithm, Moves},
    solve::{solve, SolveArgs},
};

pub struct ScrambleArgs {
    pub alg: String,
    pub moves: String,
    pub only_orientation: Vec<usize>,
    pub disregard: Vec<usize>,
    pub max_scrambles: i32,
    pub solution_found: fn(String) -> (),
}

pub fn scramble(args: ScrambleArgs) -> HashSet<String> {
    let mut start = Cube::new();
    let mut end = Cube::new();

    start.set_only_orientation(&args.only_orientation);
    end.set_only_orientation(&args.only_orientation);

    start.set_disregard(&args.disregard);
    end.set_disregard(&args.disregard);

    let alg = string_to_alg(&args.alg);
    let alg = invert_algorithm(alg);
    start.perform_alg(alg);

    let moves = Moves::from_string(&args.moves);

    let next_move_valid = if moves.has_double() {
        crate::algorithm::move_valid_double
    } else {
        crate::algorithm::move_valid_single
    };

    solve(SolveArgs {
        start,
        end,
        moves: moves,
        move_valid: next_move_valid,
        max_solutions: args.max_scrambles,
        solution_found: args.solution_found,
        log: |_: String| {},
    })
}
