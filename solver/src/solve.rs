use smallvec::SmallVec;

use crate::algorithm::{Algorithm};
use crate::moves::{NULL_MOVE, invert_move, build_alg_string, Moves};
use crate::{cube::Cube, queue::Queue, visited::Visited, node::Node};

use std::collections::HashSet;

const OPPOSITE_FACES: [u8; 6] = [2, 3, 0, 1, 5, 4];

fn same_face(alg: &Algorithm, mooove: u8) -> bool {
    if alg.is_empty() {
        return false;
    }

    let last_move = alg[alg.len() - 1];

    if last_move / 3 == mooove / 3 {
        return true;
    }

    if alg.len() == 1 {
        return false;
    }

    let second_last_move = alg[alg.len() - 2];

    mooove / 3 == OPPOSITE_FACES[last_move as usize / 3] && mooove / 3 == second_last_move / 3
}

// TODO make a logger interface and remove the 'log' boolean
// TODO make a Solver struct with a 'solve' method and a 'with_logger' method
pub fn run_solve(start: Cube, end: Cube, moves: &Moves, max_solutions: i32, log: bool) -> HashSet<String> {
    let mut depth = 0;
    let mut inverse_depth = 0;

    let mut visited = Visited::new();
    let mut inverse_visited = Visited::new();

    let mut queue = Queue::new();
    queue.push(Node{ cube: start, alg: Algorithm::new() });

    let mut inverse_queue = Queue::new();
    inverse_queue.push(Node{ cube: end, alg: Algorithm::new() });

    let mut solutions = HashSet::new();
    loop {
        let node = queue.pop();
        let inverse_node = inverse_queue.pop();

        let mut seen: HashSet<Cube> = HashSet::new();
        let algs = reconstruct_algs(&mut seen, &visited, &inverse_node.cube);
        for alg in algs {
            let mut inverse_node_alg = inverse_node.alg.clone();
            inverse_node_alg.reverse();
            let alg_str = build_alg_string(inverse_node.alg.clone(), alg);
            if log && !solutions.contains(&alg_str) {
                println!("{}", alg_str);
            }
            solutions.insert(alg_str);
        }

        let mut seen: HashSet<Cube> = HashSet::new();
        let algs = reconstruct_algs(&mut seen, &inverse_visited, &node.cube);
        for alg in algs {
            let alg_str = build_alg_string(node.alg.clone(), alg);
            if log && !solutions.contains(&alg_str) {
                println!("{}", alg_str);
            }
            solutions.insert(alg_str);
        }

        if solutions.len() >= max_solutions as usize {
            return solutions;
        }

        if log && node.alg.len() > depth {
            depth = node.alg.len();
            println!("depth: {}", depth);
        }

        if log && inverse_node.alg.len() > inverse_depth {
            inverse_depth = inverse_node.alg.len();
            println!("inverse depth: {}", inverse_depth);
        }

        for mooove in moves.get_moves() {
            go_to_child(&mut queue, &node, &mut visited, *mooove);
            go_to_child(&mut inverse_queue, &inverse_node, &mut inverse_visited, *mooove);
        }
    }
}

fn reconstruct_algs(seen: &mut HashSet<Cube>, visited: &Visited, cube: &Cube) -> Vec<Algorithm> {
    let mut algs: Vec<Algorithm> = Vec::new();

    let moves = visited.get(*cube);
    for mooove in moves {
        let mut cpy = *cube;
        if mooove == NULL_MOVE {
            return algs;
        }
        let inverted_move = invert_move(mooove);
        cpy.perform_move(inverted_move);

        if seen.contains(&cpy) {
            continue;
        } else {
            seen.insert(cpy);
        }

        let algs_subset = reconstruct_algs(seen, visited, &cpy);
        if algs_subset.is_empty() {
            let mut small_vec = SmallVec::new();
            small_vec.push(inverted_move);
            algs.push(small_vec);
        } else {
            for alg in algs_subset {
                let mut alg = alg;
                alg.push(inverted_move);
                algs.push(alg);
            }
        }
    }

    // If 'moves' is empty, then this will just return an empty vec.
    // This is the base case of the recursion.
    algs
}

fn go_to_child(queue: &mut Queue<Node>, node: &Node, visited: &mut Visited, mooove: u8) {
    if same_face(&node.alg, mooove) {
        return;
    }
    let mut cpy = node.cube;
    cpy.perform_move(mooove);

    let mut new_alg = node.alg.clone();
    new_alg.push(mooove);

    queue.push(Node{ cube: cpy, alg: new_alg });

    visited.add(cpy, mooove);
}

#[cfg(test)]
mod tests {
    use crate::moves::*;
    use super::*;

    #[test]
    fn test_solve() {
        let mut start = Cube::new();
        let end = Cube::new();
        start.perform_alg_string("R U R' F' R U R' U' R' F R2 U' R' U'");

        let moves = Moves::from_string("U,U2,U',F,F2,F',R,R2,R'");
        let max_solutions = 10;
        let log = false;

        let solutions = run_solve(start, end, &moves, max_solutions, log);
        let solutions: Vec<String> = solutions.into_iter().collect();
        let expected = [
            "R U2 F' R' F U' F' R F U' R' U'",
            "U R U F' R' F U F' R F U2 R'",
            "R U' F U' R' U' R U F' U2 R' U'",
            "U R U2 F U' R' U R U F' U R'",
            "U R U R' F U F' R F U' F' U' R'",
            "R U F U F' R' F U' F' R U' R' U'",
            "U R2 U R2 F U F' R2 F U' F' U' R2",
            "R2 U F U F' R2 F U' F' R2 U' R2 U'",
            "R U F2 U F' R' F U' F' R F' U' R' U'",
            "U R U F R' F U F' R F U' F2 U' R'",
        ];
        assert_eq!(solutions.len(), expected.len());
        for i in 0..solutions.len() {
            assert_eq!(solutions[i], expected[i]);
        }
    }
}
