use crate::cube::Cube;
use crate::moves::*;

use std::collections::HashMap;

pub type Algorithm = Vec<u8>;
pub type Algorithms = Vec<Algorithm>;

pub fn string_to_alg(alg_string: &str) -> Algorithm {
    let mut vec = Vec::new();
    for s in alg_string.split(' ') {
        match s {
            "U" => vec.push(U1_NUM),
            "U2" => vec.push(U2_NUM),
            "U'" => vec.push(U3_NUM),
            "D" => vec.push(D1_NUM),
            "D2" => vec.push(D2_NUM),
            "D'" => vec.push(D3_NUM),
            "F" => vec.push(F1_NUM),
            "F2" => vec.push(F2_NUM),
            "F'" => vec.push(F3_NUM),
            "B" => vec.push(B1_NUM),
            "B2" => vec.push(B2_NUM),
            "B'" => vec.push(B3_NUM),
            "L" => vec.push(L1_NUM),
            "L2" => vec.push(L2_NUM),
            "L'" => vec.push(L3_NUM),
            "R" => vec.push(R1_NUM),
            "R2" => vec.push(R2_NUM),
            "R'" => vec.push(R3_NUM),
            "M" => vec.push(M1_NUM),
            "M2" => vec.push(M2_NUM),
            "M'" => vec.push(M3_NUM),
            "E" => vec.push(E1_NUM),
            "E2" => vec.push(E2_NUM),
            "E'" => vec.push(E3_NUM),
            "S" => vec.push(S1_NUM),
            "S2" => vec.push(S2_NUM),
            "S'" => vec.push(S3_NUM),
            "x" => {
                vec.push(R1_NUM);
                vec.push(M3_NUM);
                vec.push(L3_NUM);
            },
            "x2" => {
                vec.push(R2_NUM);
                vec.push(M2_NUM);
                vec.push(L2_NUM);
            },
            "x'" => {
                vec.push(R3_NUM);
                vec.push(M1_NUM);
                vec.push(L1_NUM);
            },
            "y" => {
                vec.push(U1_NUM);
                vec.push(E3_NUM);
                vec.push(D3_NUM);
            },
            "y2" => {
                vec.push(U2_NUM);
                vec.push(E2_NUM);
                vec.push(D2_NUM);
            },
            "y'" => {
                vec.push(U3_NUM);
                vec.push(E1_NUM);
                vec.push(D1_NUM);
            },
            "z" => {
                vec.push(F1_NUM);
                vec.push(S1_NUM);
                vec.push(B3_NUM);
            },
            "z2" => {
                vec.push(F2_NUM);
                vec.push(S2_NUM);
                vec.push(B2_NUM);
            },
            "z'" => {
                vec.push(F3_NUM);
                vec.push(S3_NUM);
                vec.push(B1_NUM);
            },
            "l" => {
                vec.push(L1_NUM);
                vec.push(M1_NUM);
            },
            "l2" => {
                vec.push(L2_NUM);
                vec.push(M2_NUM);
            },
            "l'" => {
                vec.push(L3_NUM);
                vec.push(M3_NUM);
            },
            "r" => {
                vec.push(R1_NUM);
                vec.push(M3_NUM);
            },
            "r2" => {
                vec.push(R2_NUM);
                vec.push(M2_NUM);
            },
            "r'" => {
                vec.push(R3_NUM);
                vec.push(M1_NUM);
            },
            _ => panic!("Invalid algorithm string: {}", s),
        }
    }
    vec
}

pub struct Visited {
    visited: HashMap<[u8; 26], Algorithms>,
}

impl Visited {
    pub fn new() -> Visited {
        // Visited {
        //     visited: HashMap::new(),
        // }
        Visited {
            visited: HashMap::with_capacity(1_000_000),
        }
    }

    pub fn add(&mut self, cube: Cube, alg: Algorithm) {
        self.visited.entry(cube.state).or_default().push(alg);
    }

    pub fn get(&self, cube: Cube) -> Vec<Vec<u8>> {
        self.visited.get(&cube.state).unwrap_or(&vec![]).to_vec()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn assert_algs_equal(algs1: Algorithms, algs2: Algorithms) {
        assert_eq!(algs1.len(), algs2.len());
        for (alg1, alg2) in algs1.iter().zip(algs2.iter()) {
            assert_eq!(alg1, alg2);
        }
    }

    #[test]
    fn test_add() {
        let c1 = Cube::from_vec([0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
        let c2 = Cube::from_vec([2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);

        let mut visited = Visited::new();
        visited.add(c1, vec![0]);
        visited.add(c1, vec![3]);
        visited.add(c2, vec![6]);

        let algs = visited.get(c1);
        assert_eq!(algs.len(), 2);
        assert_algs_equal(algs, vec![vec![0], vec![3]]);

        let algs = visited.get(c2);
        assert_eq!(algs.len(), 1);
        assert_algs_equal(algs, vec![vec![6]]);
    }
}
