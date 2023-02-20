use std::collections::HashSet;

use crate::{
    algorithm::{string_to_alg, Algorithm},
    cube::Cube,
    moves::{invert_algorithm, Moves},
    solve::solve,
};

pub fn scramble(
    alg: &str,
    moves: &str,
    only_orientation: &[usize],
    max_scrambles: i32,
) -> HashSet<String> {
    let mut start = Cube::new();
    let mut end = Cube::new();

    start.set_only_orientation(only_orientation);
    end.set_only_orientation(only_orientation);

    let alg = string_to_alg(alg);
    let alg = invert_algorithm(alg);
    start.perform_alg(alg);

    let moves = Moves::from_string(moves);

    let next_move_valid = if moves.has_double() {
        |alg: &Algorithm, mooove: u8| -> bool { !crate::algorithm::same_face(alg, mooove) }
    } else {
        crate::algorithm::different_face_or_same_move
    };

    solve(start, end, &moves, next_move_valid, max_scrambles)
}
