use crate::cube::Cube;

// According to my tests, it seems that ahash is the fastest out of these hash functions:
// - ahash
// - cityhash-sys
// - twox-hash
// - Rust's default hash function

// TODO: Figure out how to use ahash with wasm. (And if it's even faster on wasm.)
// use ahash::random_state::RandomState;

use std::collections::HashMap;
use smallvec::SmallVec;

pub struct Visited {
    data: HashMap<[u8; 26], SmallVec<[u8; 1]>>,
}

impl Visited {
    pub fn new() -> Visited {
        Visited {
            data: HashMap::with_capacity(1_000_000),
        }
    }

    pub fn add(&mut self, cube: Cube, mooove: u8) {
        self.data.entry(cube.state).or_insert_with(SmallVec::new).push(mooove);
    }

    pub fn get(&self, cube: Cube) -> SmallVec<[u8; 1]> {
        self.data.get(&cube.state).cloned().unwrap_or(SmallVec::new())
    }

    pub fn contains(&self, cube: &Cube) -> bool {
        self.data.contains_key(&cube.state)
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
        visited.add(c1, 0);
        visited.add(c1, 1);
        visited.add(c2, 2);

        let algs = visited.get(c1);
        assert_eq!(algs.to_vec(), vec![0, 1]);

        let algs = visited.get(c2);
        assert_eq!(algs.to_vec(), vec![2]);
    }
}
