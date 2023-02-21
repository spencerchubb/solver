use smallvec::SmallVec;

use crate::moves::*;

pub type Algorithm = SmallVec<[u8; 8]>;
pub type Algorithms = SmallVec<[Algorithm; 1]>;

pub fn alg_to_string(alg: &Algorithm) -> String {
    let mut strings = Vec::new();
    for m in alg {
        strings.push(MOVE_NAMES[*m as usize]);
    }
    strings.join(" ")
}

pub fn string_to_alg(alg_string: &str) -> Algorithm {
    let mut alg = Algorithm::new();
    for s in alg_string.split(' ') {
        match s {
            "U" => alg.push(U1_NUM),
            "U2" => alg.push(U2_NUM),
            "U'" => alg.push(U3_NUM),
            "D" => alg.push(D1_NUM),
            "D2" => alg.push(D2_NUM),
            "D'" => alg.push(D3_NUM),
            "F" => alg.push(F1_NUM),
            "F2" => alg.push(F2_NUM),
            "F'" => alg.push(F3_NUM),
            "B" => alg.push(B1_NUM),
            "B2" => alg.push(B2_NUM),
            "B'" => alg.push(B3_NUM),
            "L" => alg.push(L1_NUM),
            "L2" => alg.push(L2_NUM),
            "L'" => alg.push(L3_NUM),
            "R" => alg.push(R1_NUM),
            "R2" => alg.push(R2_NUM),
            "R'" => alg.push(R3_NUM),
            "M" => alg.push(M1_NUM),
            "M2" => alg.push(M2_NUM),
            "M'" => alg.push(M3_NUM),
            "E" => alg.push(E1_NUM),
            "E2" => alg.push(E2_NUM),
            "E'" => alg.push(E3_NUM),
            "S" => alg.push(S1_NUM),
            "S2" => alg.push(S2_NUM),
            "S'" => alg.push(S3_NUM),
            "x" => {
                alg.push(R1_NUM);
                alg.push(M3_NUM);
                alg.push(L3_NUM);
            },
            "x2" => {
                alg.push(R2_NUM);
                alg.push(M2_NUM);
                alg.push(L2_NUM);
            },
            "x'" => {
                alg.push(R3_NUM);
                alg.push(M1_NUM);
                alg.push(L1_NUM);
            },
            "y" => {
                alg.push(U1_NUM);
                alg.push(E3_NUM);
                alg.push(D3_NUM);
            },
            "y2" => {
                alg.push(U2_NUM);
                alg.push(E2_NUM);
                alg.push(D2_NUM);
            },
            "y'" => {
                alg.push(U3_NUM);
                alg.push(E1_NUM);
                alg.push(D1_NUM);
            },
            "z" => {
                alg.push(F1_NUM);
                alg.push(S1_NUM);
                alg.push(B3_NUM);
            },
            "z2" => {
                alg.push(F2_NUM);
                alg.push(S2_NUM);
                alg.push(B2_NUM);
            },
            "z'" => {
                alg.push(F3_NUM);
                alg.push(S3_NUM);
                alg.push(B1_NUM);
            },
            "l" => {
                alg.push(L1_NUM);
                alg.push(M1_NUM);
            },
            "l2" => {
                alg.push(L2_NUM);
                alg.push(M2_NUM);
            },
            "l'" => {
                alg.push(L3_NUM);
                alg.push(M3_NUM);
            },
            "r" => {
                alg.push(R1_NUM);
                alg.push(M3_NUM);
            },
            "r2" => {
                alg.push(R2_NUM);
                alg.push(M2_NUM);
            },
            "r'" => {
                alg.push(R3_NUM);
                alg.push(M1_NUM);
            },
            _ => panic!("Invalid algorithm string: {}", s),
        }
    }
    alg
}

const OPPOSITES: [u8; 6] = [2, 3, 0, 1, 5, 4];

pub fn move_valid_double(alg: &Algorithm, mooove: u8) -> bool {
    if alg.is_empty() {
        return true;
    }

    let last_move = alg[alg.len() - 1];

    if last_move / 3 == mooove / 3 {
        return false;
    }

    if alg.len() == 1 {
        return true;
    }

    let second_last_move = alg[alg.len() - 2];

    mooove / 3 != OPPOSITES[last_move as usize / 3] || mooove / 3 != second_last_move / 3
}

pub fn move_valid_single(alg: &Algorithm, mooove: u8) -> bool {
    if alg.is_empty() {
        return true;
    }

    let last_move = alg[alg.len() - 1];

    last_move / 3 != mooove / 3 || last_move == mooove
}
