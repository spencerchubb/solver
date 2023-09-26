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
    // OLL case
    // let alg = "R U2 R2 F R F' U2 R' F R F'";
    // let moves = "U U' D D' L L' R R' F F' B B'";
    // let only_orientation: &[usize] = &[0, 1, 2, 3, 8, 9, 10, 11];
    // let disregard: &[usize] = &[];

    // Find cross solution
    // let alg = "L2 D2 F2 L' F2 R' F2 L2 B2 R B2 R2 B D F2 R' F' U B R2 B'";
    // let moves = "U U' D D' R R' L L' F F' B B'";
    // let only_orientation: &[usize] = &[];
    // let disregard: &[usize] = &[0,1,2,3,4,5,6,7,12,13,14,15,16,17,18,19];

    // EOLine
    let alg = "L2 F2 L2 B R2 F' R2 U2 F' R2 U2 L2 R' U L2 F2 R' D' B2 R' U";
    let moves = "U U' D D' R R' L L' F F' B B'";
    let only_orientation: &[usize] = &[8, 9, 10, 11, 12, 13, 15, 16, 18, 19];
    let disregard: &[usize] = &[0, 1, 2, 3, 4, 5, 6, 7];

    let mut start = crate::cube::Cube::new();
    let mut end = crate::cube::Cube::new();

    start.set_only_orientation(&only_orientation);
    end.set_only_orientation(&only_orientation);

    start.set_disregard(&disregard);
    end.set_disregard(&disregard);

    let alg = crate::algorithm::string_to_alg(&alg);
    start.perform_alg(alg);

    let moves = crate::moves::Moves::from_string(moves);

    let next_move_valid = if moves.has_double() {
        crate::algorithm::move_valid_double
    } else {
        crate::algorithm::move_valid_single
    };

    crate::solve::solve(crate::solve::SolveArgs {
        start: end,
        end: start,
        moves: moves,
        move_valid: next_move_valid,
        max_solutions: 5,
        solution_found: |s: String| println!("{}", s),
        log: |_: String| {},
    });
}
