use crate::{algorithm::{Algorithm, string_to_alg}, constants::DISREGARD, moves::*};

// I have decided not to store centers because it makes the program slower.
// As a result, this program can only solve states where the centers are solved.
//
//          00 08 01
//          09    10
//          02 11 03
// 00 09 02 02 11 03 03 10 01 01 08 00
// 18    12 12    13 13    19 19    18
// 06 15 04 04 14 05 05 16 07 07 17 06
//          04 14 05
//          15    16
//          06 17 07

pub type CubeState = [u8; 20];

#[derive(Copy, Clone)]
#[derive(PartialEq, Eq, Hash)]
pub struct Cube {
    pub state: CubeState,
}

impl Cube {
    pub fn new() -> Cube {
        Cube {
            state: [
                0, 1, 2, 3, 4, 5, 6, 7, // corners
                0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, //edges
            ],
        }
    }

    pub fn from_vec(vec: CubeState) -> Cube {
        Cube { state: vec }
    }

    pub fn perform_move(&mut self, m: u8) {
        match m {
            U1_NUM => self.u1(),
            U2_NUM => self.u2(),
            U3_NUM => self.u3(),
            F1_NUM => self.f1(),
            F2_NUM => self.f2(),
            F3_NUM => self.f3(),
            D1_NUM => self.d1(),
            D2_NUM => self.d2(),
            D3_NUM => self.d3(),
            B1_NUM => self.b1(),
            B2_NUM => self.b2(),
            B3_NUM => self.b3(),
            L1_NUM => self.l1(),
            L2_NUM => self.l2(),
            L3_NUM => self.l3(),
            R1_NUM => self.r1(),
            R2_NUM => self.r2(),
            R3_NUM => self.r3(),
            M1_NUM => self.m1(),
            M2_NUM => self.m2(),
            M3_NUM => self.m3(),
            E1_NUM => self.e1(),
            E2_NUM => self.e2(),
            E3_NUM => self.e3(),
            S1_NUM => self.s1(),
            S2_NUM => self.s2(),
            S3_NUM => self.s3(),
            _ => panic!("Invalid move number: {}", m),
        }
    }

    pub fn perform_alg(&mut self, alg: Algorithm) {
        for m in alg {
            self.perform_move(m);
        }
    }

    pub fn perform_alg_string(&mut self, alg_string: &str) {
        let alg: Algorithm = string_to_alg(alg_string);
        self.perform_alg(alg);
    }

    // pieces is an array of piece indices.
    // This function makes it so we only consider the orientation of the pieces.
    // The way this works is by settings the 5th-8th bits to 1.
    // This means that the permutation is disregarded.
    pub fn set_only_orientation(&mut self, pieces: &[usize]) {
        for i in pieces {
            self.state[*i] |= 0xF0;
        }
    }

    // pieces is an array of pieces indices.
    // This function makes it so the pieces are disregarded.
    pub fn set_disregard(&mut self, pieces: &[usize]) {
        for i in pieces {
            self.state[*i] |= DISREGARD;
        }
    }

    fn u1(&mut self) {
        let temp = self.state[0];
        self.state[0] = self.state[2];
        self.state[2] = self.state[3];
        self.state[3] = self.state[1];
        self.state[1] = temp;

        let temp = self.state[8];
        self.state[8] = self.state[9];
        self.state[9] = self.state[11];
        self.state[11] = self.state[10];
        self.state[10] = temp;
    }

    fn u2(&mut self) {
        self.state.swap(0, 3);
        self.state.swap(1, 2);
        self.state.swap(8, 11);
        self.state.swap(9, 10);
    }

    fn u3(&mut self) {
        let temp = self.state[0];
        self.state[0] = self.state[1];
        self.state[1] = self.state[3];
        self.state[3] = self.state[2];
        self.state[2] = temp;

        let temp = self.state[8];
        self.state[8] = self.state[10];
        self.state[10] = self.state[11];
        self.state[11] = self.state[9];
        self.state[9] = temp;
    }

    fn f1(&mut self) {
        let temp = self.state[2];
        self.state[2] = twist_ccw(self.state[4]);
        self.state[4] = twist_cw(self.state[5]);
        self.state[5] = twist_ccw(self.state[3]);
        self.state[3] = twist_cw(temp);

        let temp = self.state[11];
        self.state[11] = flip(self.state[12]);
        self.state[12] = flip(self.state[14]);
        self.state[14] = flip(self.state[13]);
        self.state[13] = flip(temp);
    }

    fn f2(&mut self) {
        self.state.swap(2, 5);
        self.state.swap(3, 4);
        self.state.swap(11, 14);
        self.state.swap(12, 13);
    }

    fn f3(&mut self) {
        let temp = self.state[2];
        self.state[2] = twist_ccw(self.state[3]);
        self.state[3] = twist_cw(self.state[5]);
        self.state[5] = twist_ccw(self.state[4]);
        self.state[4] = twist_cw(temp);

        let temp = self.state[11];
        self.state[11] = flip(self.state[13]);
        self.state[13] = flip(self.state[14]);
        self.state[14] = flip(self.state[12]);
        self.state[12] = flip(temp);
    }

    fn d1(&mut self) {
        let temp = self.state[4];
        self.state[4] = self.state[6];
        self.state[6] = self.state[7];
        self.state[7] = self.state[5];
        self.state[5] = temp;

        let temp = self.state[14];
        self.state[14] = self.state[15];
        self.state[15] = self.state[17];
        self.state[17] = self.state[16];
        self.state[16] = temp;
    }

    fn d2(&mut self) {
        self.state.swap(4, 7);
        self.state.swap(5, 6);
        self.state.swap(14, 17);
        self.state.swap(15, 16);
    }

    fn d3(&mut self) {
        let temp = self.state[4];
        self.state[4] = self.state[5];
        self.state[5] = self.state[7];
        self.state[7] = self.state[6];
        self.state[6] = temp;

        let temp = self.state[14];
        self.state[14] = self.state[16];
        self.state[16] = self.state[17];
        self.state[17] = self.state[15];
        self.state[15] = temp;
    }

    fn b1(&mut self) {
        let temp = self.state[0];
        self.state[0] = twist_cw(self.state[1]);
        self.state[1] = twist_ccw(self.state[7]);
        self.state[7] = twist_cw(self.state[6]);
        self.state[6] = twist_ccw(temp);

        let temp = self.state[8];
        self.state[8] = flip(self.state[19]);
        self.state[19] = flip(self.state[17]);
        self.state[17] = flip(self.state[18]);
        self.state[18] = flip(temp);
    }

    fn b2(&mut self) {
        self.state.swap(0, 7);
        self.state.swap(1, 6);
        self.state.swap(8, 17);
        self.state.swap(18, 19);
    }

    fn b3(&mut self) {
        let temp = self.state[0];
        self.state[0] = twist_cw(self.state[6]);
        self.state[6] = twist_ccw(self.state[7]);
        self.state[7] = twist_cw(self.state[1]);
        self.state[1] = twist_ccw(temp);

        let temp = self.state[8];
        self.state[8] = flip(self.state[18]);
        self.state[18] = flip(self.state[17]);
        self.state[17] = flip(self.state[19]);
        self.state[19] = flip(temp);
    }

    fn l1(&mut self) {
        let temp = self.state[0];
        self.state[0] = twist_ccw(self.state[6]);
        self.state[6] = twist_cw(self.state[4]);
        self.state[4] = twist_ccw(self.state[2]);
        self.state[2] = twist_cw(temp);

        let temp = self.state[9];
        self.state[9] = self.state[18];
        self.state[18] = self.state[15];
        self.state[15] = self.state[12];
        self.state[12] = temp;
    }

    fn l2(&mut self) {
        self.state.swap(0, 4);
        self.state.swap(2, 6);
        self.state.swap(9, 15);
        self.state.swap(12, 18);
    }

    fn l3(&mut self) {
        let temp = self.state[0];
        self.state[0] = twist_ccw(self.state[2]);
        self.state[2] = twist_cw(self.state[4]);
        self.state[4] = twist_ccw(self.state[6]);
        self.state[6] = twist_cw(temp);

        let temp = self.state[9];
        self.state[9] = self.state[12];
        self.state[12] = self.state[15];
        self.state[15] = self.state[18];
        self.state[18] = temp;
    }

    fn r1(&mut self) {
        let temp = self.state[1];
        self.state[1] = twist_cw(self.state[3]);
        self.state[3] = twist_ccw(self.state[5]);
        self.state[5] = twist_cw(self.state[7]);
        self.state[7] = twist_ccw(temp);

        let temp = self.state[10];
        self.state[10] = self.state[13];
        self.state[13] = self.state[16];
        self.state[16] = self.state[19];
        self.state[19] = temp;
    }

    fn r2(&mut self) {
        self.state.swap(1, 5);
        self.state.swap(3, 7);
        self.state.swap(10, 16);
        self.state.swap(13, 19);
    }

    fn r3(&mut self) {
        let temp = self.state[1];
        self.state[1] = twist_cw(self.state[7]);
        self.state[7] = twist_ccw(self.state[5]);
        self.state[5] = twist_cw(self.state[3]);
        self.state[3] = twist_ccw(temp);

        let temp = self.state[10];
        self.state[10] = self.state[19];
        self.state[19] = self.state[16];
        self.state[16] = self.state[13];
        self.state[13] = temp;
    }

    fn m1(&mut self) {
        let temp = self.state[8];
        self.state[8] = flip(self.state[17]);
        self.state[17] = flip(self.state[14]);
        self.state[14] = flip(self.state[11]);
        self.state[11] = flip(temp);
    }

    fn m2(&mut self) {
        self.state.swap(8, 14);
        self.state.swap(11, 17);
    }

    fn m3(&mut self) {
        let temp = self.state[8];
        self.state[8] = flip(self.state[11]);
        self.state[11] = flip(self.state[14]);
        self.state[14] = flip(self.state[17]);
        self.state[17] = flip(temp);
    }

    fn e1(&mut self) {
        let temp = self.state[12];
        self.state[12] = flip(self.state[18]);
        self.state[18] = flip(self.state[19]);
        self.state[19] = flip(self.state[13]);
        self.state[13] = flip(temp);
    }

    fn e2(&mut self) {
        self.state.swap(12, 19);
        self.state.swap(13, 18);
    }

    fn e3(&mut self) {
        let temp = self.state[12];
        self.state[12] = flip(self.state[13]);
        self.state[13] = flip(self.state[19]);
        self.state[19] = flip(self.state[18]);
        self.state[18] = flip(temp);
    }

    fn s1(&mut self) {
        let temp = self.state[9];
        self.state[9] = flip(self.state[15]);
        self.state[15] = flip(self.state[16]);
        self.state[16] = flip(self.state[10]);
        self.state[10] = flip(temp);
    }

    fn s2(&mut self) {
        self.state.swap(9, 16);
        self.state.swap(10, 15);
    }

    fn s3(&mut self) {
        let temp = self.state[9];
        self.state[9] = flip(self.state[10]);
        self.state[10] = flip(self.state[16]);
        self.state[16] = flip(self.state[15]);
        self.state[15] = flip(temp);
    }
}

fn twist_cw(b: u8) -> u8 {
    if b == DISREGARD {
        return b;
    }
    let upper = b & 0xF0;
    let lower = b & 0x0F;
    ((upper + 1) % 3) << 4 | lower
}

fn twist_ccw(b: u8) -> u8 {
    if b == DISREGARD {
        return b;
    }
    let upper = b & 0xF0;
    let lower = b & 0x0F;
    ((upper + 2) % 3) << 4 | lower
}

fn flip(b: u8) -> u8 {
    if b == DISREGARD {
        return b;
    }
    b ^ 0b00010000
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_random() {
        let mut c1 = Cube::new();
        let c2 = Cube::new();

        // Scramble
        c1.perform_alg_string("R F B2 U2 B' L' D F' B L2 D2 R2 U R' D' F2 U' L");

        // Cross
        c1.perform_alg_string("R D2 L2");

        // Pair 1
        c1.perform_alg_string("D F D' F' D F' D2 F");

        // Pair 2
        c1.perform_alg_string("D2 R' D R");

        // Pair 3+4
        c1.perform_alg_string("L' D' L D' B D B2 D' B D' R D R'");

        // OLL
        c1.perform_alg_string("B' D' L' D L B");

        // PLL
        c1.perform_alg_string("R D R' D R D R' B' R D R' D' R' B R2 D' R' D2 R D' R'");

        assert_eq!(c1.state, c2.state);
    }

    #[test]
    fn test_slice() {
        let mut c1 = Cube::new();
        let c2 = Cube::new();

        c1.perform_alg_string("M E S M' S2 E' M2 S' E2 S M2 S E2");

        assert_eq!(c1.state, c2.state);
    }
}
