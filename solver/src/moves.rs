use crate::algorithm::{Algorithm};
use crate::cube::{Cube, DISREGARD};
use crate::algorithm::string_to_alg;

//          00 08 01
//          09 20 10
//          02 11 03
// 00 09 02 02 11 03 03 10 02 02 08 00
// 18 24 12 12 21 13 13 25 19 19 23 18
// 06 15 04 04 14 05 05 16 07 07 17 06
//          04 14 05
//          15 22 16
//          06 17 07

pub const U1_NUM: u8 = 0;
pub const U2_NUM: u8 = 1;
pub const U3_NUM: u8 = 2;
pub const F1_NUM: u8 = 3;
pub const F2_NUM: u8 = 4;
pub const F3_NUM: u8 = 5;
pub const D1_NUM: u8 = 6;
pub const D2_NUM: u8 = 7;
pub const D3_NUM: u8 = 8;
pub const B1_NUM: u8 = 9;
pub const B2_NUM: u8 = 10;
pub const B3_NUM: u8 = 11;
pub const L1_NUM: u8 = 12;
pub const L2_NUM: u8 = 13;
pub const L3_NUM: u8 = 14;
pub const R1_NUM: u8 = 15;
pub const R2_NUM: u8 = 16;
pub const R3_NUM: u8 = 17;
pub const M1_NUM: u8 = 18;
pub const M2_NUM: u8 = 19;
pub const M3_NUM: u8 = 20;
pub const E1_NUM: u8 = 21;
pub const E2_NUM: u8 = 22;
pub const E3_NUM: u8 = 23;
pub const S1_NUM: u8 = 24;
pub const S2_NUM: u8 = 25;
pub const S3_NUM: u8 = 26;
pub const NULL_MOVE: u8 = 0xFF;

pub const MOVE_NAMES: [&str; 27] = [
    "U", "U2", "U'",
    "F", "F2", "F'",
    "D", "D2", "D'",
    "B", "B2", "B'",
    "L", "L2", "L'",
    "R", "R2", "R'",
    "M", "M2", "M'",
    "E", "E2", "E'",
    "S", "S2", "S'",
];

const NO_CANCEL: u8 = 0xFF;
const PERFECT_CANCEL: u8 = 0xFE;

const EQUIVALENCES: [[u8; 3]; 27] = [
    [U2_NUM, U3_NUM, PERFECT_CANCEL],
    [U3_NUM, PERFECT_CANCEL, U1_NUM],
    [PERFECT_CANCEL, U1_NUM, U2_NUM],
    [F2_NUM, F3_NUM, PERFECT_CANCEL],
    [F3_NUM, PERFECT_CANCEL, F1_NUM],
    [PERFECT_CANCEL, F1_NUM, F2_NUM],
    [D2_NUM, D3_NUM, PERFECT_CANCEL],
    [D3_NUM, PERFECT_CANCEL, D1_NUM],
    [PERFECT_CANCEL, D1_NUM, D2_NUM],
    [B2_NUM, B3_NUM, PERFECT_CANCEL],
    [B3_NUM, PERFECT_CANCEL, B1_NUM],
    [PERFECT_CANCEL, B1_NUM, B2_NUM],
    [L2_NUM, L3_NUM, PERFECT_CANCEL],
    [L3_NUM, PERFECT_CANCEL, L1_NUM],
    [PERFECT_CANCEL, L1_NUM, L2_NUM],
    [R2_NUM, R3_NUM, PERFECT_CANCEL],
    [R3_NUM, PERFECT_CANCEL, R1_NUM],
    [PERFECT_CANCEL, R1_NUM, R2_NUM],
    [M2_NUM, M3_NUM, PERFECT_CANCEL],
    [M3_NUM, PERFECT_CANCEL, M1_NUM],
    [PERFECT_CANCEL, M1_NUM, M2_NUM],
    [E2_NUM, E3_NUM, PERFECT_CANCEL],
    [E3_NUM, PERFECT_CANCEL, E1_NUM],
    [PERFECT_CANCEL, E1_NUM, E2_NUM],
    [S2_NUM, S3_NUM, PERFECT_CANCEL],
    [S3_NUM, PERFECT_CANCEL, S1_NUM],
    [PERFECT_CANCEL, S1_NUM, S2_NUM],
];

pub fn invert_move(m: u8) -> u8 {
    match m % 3 {
        0 => m + 2,
        1 => m,
        2 => m - 2,
        _ => panic!("unreachable"),
    }
}

pub fn invert_algorithm(alg: Algorithm) -> Algorithm {
    let mut out = Algorithm::new();
    for m in alg {
        out.push(invert_move(m));
    }
    out.reverse();
    out
}

fn cancel_pair_of_moves(m1: u8, m2: u8) -> u8 {
    let face1 = m1 / 3;
    let face2 = m2 / 3;
    if face1 != face2 {
        return NO_CANCEL;
    }
    EQUIVALENCES[m1 as usize][m2 as usize % 3]
}

pub fn build_alg_string(forward: Algorithm, inverse: Algorithm) -> String {
    // println!("build_alg_string: {} - {}", alg_to_string(&forward), alg_to_string(&inverse));
    let mut reversed = inverse;
    reversed.reverse();
    let mut combined = forward;
    combined.extend(reversed);
    let mut cleaned: Vec<u8> = Vec::new();
    for m in combined {
        let len = cleaned.len();
        if len > 0 {
            let cancel = cancel_pair_of_moves(cleaned[cleaned.len() - 1], m);
            if cancel == NO_CANCEL {
                cleaned.push(m);
            } else if cancel == PERFECT_CANCEL {
                cleaned.pop();
            } else {
                cleaned[len - 1] = cancel;
            }
        } else {
            cleaned.push(m);
        }
    }
    cleaned
        .iter()
        .map(|m| MOVE_NAMES[*m as usize])
        .collect::<Vec<&str>>()
        .join(" ")
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

impl Cube {
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

        let temp = self.state[20];
        self.state[20] = self.state[23];
        self.state[23] = self.state[22];
        self.state[22] = self.state[21];
        self.state[21] = temp;
    }

    fn m2(&mut self) {
        self.state.swap(8, 14);
        self.state.swap(11, 17);
        self.state.swap(20, 22);
        self.state.swap(21, 23);
    }

    fn m3(&mut self) {
        let temp = self.state[8];
        self.state[8] = flip(self.state[11]);
        self.state[11] = flip(self.state[14]);
        self.state[14] = flip(self.state[17]);
        self.state[17] = flip(temp);

        let temp = self.state[20];
        self.state[20] = self.state[21];
        self.state[21] = self.state[22];
        self.state[22] = self.state[23];
        self.state[23] = temp;
    }

    fn e1(&mut self) {
        let temp = self.state[12];
        self.state[12] = flip(self.state[18]);
        self.state[18] = flip(self.state[19]);
        self.state[19] = flip(self.state[13]);
        self.state[13] = flip(temp);

        let temp = self.state[21];
        self.state[21] = self.state[24];
        self.state[24] = self.state[23];
        self.state[23] = self.state[25];
        self.state[25] = temp;
    }

    fn e2(&mut self) {
        self.state.swap(12, 19);
        self.state.swap(13, 18);
        self.state.swap(21, 23);
        self.state.swap(24, 25);
    }

    fn e3(&mut self) {
        let temp = self.state[12];
        self.state[12] = flip(self.state[13]);
        self.state[13] = flip(self.state[19]);
        self.state[19] = flip(self.state[18]);
        self.state[18] = flip(temp);

        let temp = self.state[21];
        self.state[21] = self.state[25];
        self.state[25] = self.state[23];
        self.state[23] = self.state[24];
        self.state[24] = temp;
    }

    fn s1(&mut self) {
        let temp = self.state[9];
        self.state[9] = flip(self.state[15]);
        self.state[15] = flip(self.state[16]);
        self.state[16] = flip(self.state[10]);
        self.state[10] = flip(temp);

        let temp = self.state[20];
        self.state[20] = self.state[24];
        self.state[24] = self.state[22];
        self.state[22] = self.state[25];
        self.state[25] = temp;
    }

    fn s2(&mut self) {
        self.state.swap(9, 16);
        self.state.swap(10, 15);
        self.state.swap(20, 22);
        self.state.swap(24, 25);
    }

    fn s3(&mut self) {
        let temp = self.state[9];
        self.state[9] = flip(self.state[10]);
        self.state[10] = flip(self.state[16]);
        self.state[16] = flip(self.state[15]);
        self.state[15] = flip(temp);

        let temp = self.state[20];
        self.state[20] = self.state[25];
        self.state[25] = self.state[22];
        self.state[22] = self.state[24];
        self.state[24] = temp;
    }
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
