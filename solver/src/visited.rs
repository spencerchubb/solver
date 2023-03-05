use crate::cube::Cube;

use smallvec::SmallVec;
use std::collections::HashMap;

// We use FNV because it's fast and we don't need it to be cryptographically secure.
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
const FNV_OFFSET: u64 = 14695981039346656037;
const FNV_PRIME: u64 = 1099511628211;

pub struct MyHasher {
    state: u64,
}

impl std::hash::Hasher for MyHasher {
    fn write(&mut self, bytes: &[u8]) {
        for &byte in bytes {
            self.state ^= u64::from(byte);
            self.state = self.state.wrapping_mul(FNV_PRIME);
        }
    }

    fn finish(&self) -> u64 {
        self.state
    }
}

pub struct BuildMyHasher;

impl std::hash::BuildHasher for BuildMyHasher {
    type Hasher = MyHasher;
    fn build_hasher(&self) -> MyHasher {
        MyHasher { state: FNV_OFFSET }
    }
}

pub struct Visited {
    data: HashMap<[u8; 26], SmallVec<[u8; 1]>, BuildMyHasher>,
}

impl Visited {
    pub fn new() -> Visited {
        Visited {
            // data: HashMap::with_capacity(1_000_000),
            data: HashMap::with_capacity_and_hasher(1_000_000, BuildMyHasher),
        }
    }

    pub fn add(&mut self, cube: Cube, mooove: u8) {
        self.data
            .entry(cube.state)
            .or_insert_with(SmallVec::new)
            .push(mooove);
    }

    pub fn get(&self, cube: Cube) -> SmallVec<[u8; 1]> {
        self.data
            .get(&cube.state)
            .cloned()
            .unwrap_or(SmallVec::new())
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
        let c1 = Cube::from_vec([
            0, 1, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        ]);
        let c2 = Cube::from_vec([
            2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        ]);

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
