use arrayvec::ArrayVec;

use crate::algorithm::{Algorithm, AlgorithmSegment};

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

pub fn invert_alg_segment(alg: AlgorithmSegment) -> AlgorithmSegment {
    let mut out = ArrayVec::new();
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

pub fn combine_algs(forward: AlgorithmSegment, inverse: AlgorithmSegment) -> String {
    let mut combined = Algorithm::new();
    combined.extend(forward);
    combined.extend(inverse);
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

// The moves that can be performed while searching for a solution.
// For example, the user may want to restrict solutions to the RUF moves.
pub struct Moves {
    moves: Vec<u8>,
}

impl Moves {

    // This function accepts a string of comma-separated move names.
    // Example: "U,U2,U',F,F2,F',R,R2,R'"
    pub fn from_string(s: &str) -> Moves {
        let mut moves = Vec::new();
        for m in s.split(' ') {
            match m {
                "U" => moves.push(U1_NUM),
                "U2" => moves.push(U2_NUM),
                "U'" => moves.push(U3_NUM),
                "F" => moves.push(F1_NUM),
                "F2" => moves.push(F2_NUM),
                "F'" => moves.push(F3_NUM),
                "D" => moves.push(D1_NUM),
                "D2" => moves.push(D2_NUM),
                "D'" => moves.push(D3_NUM),
                "B" => moves.push(B1_NUM),
                "B2" => moves.push(B2_NUM),
                "B'" => moves.push(B3_NUM),
                "L" => moves.push(L1_NUM),
                "L2" => moves.push(L2_NUM),
                "L'" => moves.push(L3_NUM),
                "R" => moves.push(R1_NUM),
                "R2" => moves.push(R2_NUM),
                "R'" => moves.push(R3_NUM),
                _ => panic!("invalid move: {}", m),
            }
        }
        Moves { moves }
    }

    pub fn get_moves(&self) -> &Vec<u8> {
        &self.moves
    }

    // Returns true if the moves contain one or more double moves.
    pub fn has_double(&self) -> bool {
        for m in &self.moves {
            if m % 3 == 1 {
                return true;
            }
        }
        false
    }
}

#[cfg(test)]
mod tests {
    use crate::moves::*;

    #[test]
    fn test_moves() {
        let moves = Moves::from_string("U U2 U' F F2 F' R R2 R'");
        assert_eq!(*moves.get_moves(), vec![U1_NUM, U2_NUM, U3_NUM, F1_NUM, F2_NUM, F3_NUM, R1_NUM, R2_NUM, R3_NUM]);
    }
}
