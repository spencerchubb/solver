pub const DISREGARD: u8 = 0xFF;

#[derive(Copy, Clone)]
pub struct Cube {
    pub state: [u8; 26],
}

impl Cube {
    pub fn new() -> Cube {
        Cube {
            state: [
                0, 1, 2, 3, 4, 5, 6, 7, // corners
                0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, //edges
                0, 1, 2, 3, 4, 5, // centers
            ],
        }
    }

    pub fn from_vec(vec: [u8; 26]) -> Cube {
        Cube { state: vec }
    }
}