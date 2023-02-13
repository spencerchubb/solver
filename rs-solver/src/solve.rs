use crate::algorithm::Algorithm;
use crate::{arch::check_32_bit, cube::Cube, moves::alg_string, queue::Queue, visited::Visited, node::Node};

use std::collections::HashSet;
use std::time::{SystemTime, UNIX_EPOCH};

const OPPOSITE_FACES: [u8; 6] = [2, 3, 0, 1, 5, 4];

fn same_face(alg: &Algorithm, mooove: u8) -> bool {
    if alg.len() == 0 {
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

pub fn run_solve(start: Cube, end: Cube, moves: &[u8], max_solutions: i32, max_ms: u128, log: bool) -> HashSet<String> {
    check_32_bit();

    let mut depth = 0;
    let mut inverse_depth = 0;

    let mut visited = Visited::new();
    let mut inverse_visited = Visited::new();

    let mut queue = Queue::new();
    queue.push(Node{ cube: start, alg: Algorithm::new() });

    let mut inverse_queue = Queue::new();
    inverse_queue.push(Node{ cube: end, alg: Algorithm::new() });

    let mut solutions = HashSet::new();
    let start_ms = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis();
    let mut cumulative_time = 0;
    while SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - start_ms < max_ms {
        let node = queue.pop();
        let inverse_node = inverse_queue.pop();

        let algs = visited.get(inverse_node.cube);
        for alg in algs {
            let alg_str = alg_string(alg, inverse_node.alg.clone());
            if log {
                println!("{}", alg_str);
            }
            solutions.insert(alg_str);
        }

        let algs = inverse_visited.get(node.cube);
        for alg in algs {
            let alg_str = alg_string(node.alg.clone(), alg);
            if log {
                println!("{}", alg_str);
            }
            solutions.insert(alg_str);
        }

        if solutions.len() >= max_solutions as usize {
            println!("cumulative time: {} ms", cumulative_time);
            println!("elapsed before drop: {} ms", SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - start_ms);
            return solutions;
        }

        if log && node.alg.len() > depth {
            depth = node.alg.len();
            println!("Searching depth: {} ms", depth);
        }

        if log && inverse_node.alg.len() > inverse_depth {
            inverse_depth = inverse_node.alg.len();
            println!("Searching inverse depth: {}", inverse_depth);
        }

        for mooove in moves {
            if !same_face(&node.alg, *mooove) {
                let mut cpy = node.cube;
                cpy.perform_move(*mooove);
        
                let mut new_alg = node.alg.clone();
                new_alg.push(*mooove);
        
                let temp_time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis();
                queue.push(Node{ cube: cpy, alg: new_alg.clone() });
                cumulative_time += SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - temp_time;
        
                visited.add(cpy, new_alg);
            }
            if !same_face(&inverse_node.alg, *mooove) {
                let mut cpy = inverse_node.cube;
                cpy.perform_move(*mooove);
        
                let mut new_alg = inverse_node.alg.clone();
                new_alg.push(*mooove);
        
                let temp_time = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis();
                inverse_queue.push(Node{ cube: cpy, alg: new_alg.clone() });
                cumulative_time += SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - temp_time;
        
                inverse_visited.add(cpy, new_alg);
            }

            // go_to_child(&mut queue, &node, &mut visited, *mooove);
            // go_to_child(&mut inverse_queue, &inverse_node, &mut inverse_visited, *mooove);
        }
    }

    println!("cumulative time: {} ms", cumulative_time);
    println!("elapsed before drop: {} ms", SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_millis() - start_ms);
    solutions
}

// fn go_to_child(queue: &mut VecDeque<Node>, node: &Node, visited: &mut Visited, mooove: u8) {
//     if !same_face(&node.moves, mooove) {
//         let mut cpy = node.cube;
//         cpy.perform_move(mooove);

//         let mut new_moves = node.moves.clone();
//         new_moves.push(mooove);

//         queue.push_back(Node{ cube: cpy, moves: new_moves.clone() });

//         visited.add(cpy, new_moves);
//     }
// }

#[cfg(test)]
mod tests {
    use crate::moves::*;
    use super::*;

    #[test]
    fn test_solve() {
        let mut start = Cube::new();
        let end = Cube::new();
        start.perform_alg_string("R U R' F' R U R' U' R' F R2 U' R' U'");

        let moves = [U1_NUM, U2_NUM, U3_NUM, F1_NUM, F2_NUM, F3_NUM, R1_NUM, R2_NUM, R3_NUM];
        let max_solutions = 10;
        let max_ms = 10_000;
        let log = false;

        let solutions = run_solve(start, end, &moves, max_solutions, max_ms, log);
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
