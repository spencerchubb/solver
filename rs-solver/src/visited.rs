use crate::algorithm::{Algorithm, Algorithms};
use crate::cube::Cube;

use std::collections::HashMap;

pub struct Visited {
    data: HashMap<[u8; 26], Algorithms>,
}

impl Visited {
    pub fn new() -> Visited {
        Visited {
            data: HashMap::with_capacity(1_000_000),
        }
    }

    pub fn add(&mut self, cube: Cube, alg: Algorithm) {
        self.data.entry(cube.state).or_insert(Algorithms::new()).push(alg);
    }

    pub fn get(&self, cube: Cube) -> Algorithms {
        self.data.get(&cube.state).cloned().unwrap_or(Algorithms::new())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        let c1 = Cube::from_vec([0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
        let c2 = Cube::from_vec([2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);

        let mut visited = Visited::new();
        visited.add(c1, Algorithm::from_vec(vec![0]));
        visited.add(c1, Algorithm::from_vec(vec![3]));
        visited.add(c2, Algorithm::from_vec(vec![6]));

        let algs = visited.get(c1);
        let algs = algs.to_vec().iter().map(|x| x.to_vec()).collect::<Vec<_>>();
        assert_eq!(algs, vec![vec![0], vec![3]]);

        let algs = visited.get(c2);
        let algs = algs.to_vec().iter().map(|x| x.to_vec()).collect::<Vec<_>>();
        assert_eq!(algs, vec![vec![6]]);
    }
}
