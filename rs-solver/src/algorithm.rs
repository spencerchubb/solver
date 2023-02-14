use smallvec::SmallVec;

use crate::moves::*;

pub type Algorithm = SmallVec<[u8; 8]>;
pub type Algorithms = SmallVec<[Algorithm; 1]>;

pub fn alg_to_string(alg: &Algorithm) -> String {
    // let mut s = String::new();
    let mut strings = Vec::new();
    for m in alg {
        // match m {
        //     U1_NUM => s.push_str("U "),
        //     U2_NUM => s.push_str("U2 "),
        //     U3_NUM => s.push_str("U' "),
        //     D1_NUM => s.push_str("D "),
        //     D2_NUM => s.push_str("D2 "),
        //     D3_NUM => s.push_str("D' "),
        //     F1_NUM => s.push_str("F "),
        //     F2_NUM => s.push_str("F2 "),
        //     F3_NUM => s.push_str("F' "),
        //     B1_NUM => s.push_str("B "),
        //     B2_NUM => s.push_str("B2 "),
        //     B3_NUM => s.push_str("B' "),
        //     L1_NUM => s.push_str("L "),
        //     L2_NUM => s.push_str("L2 "),
        //     L3_NUM => s.push_str("L' "),
        //     R1_NUM => s.push_str("R "),
        //     R2_NUM => s.push_str("R2 "),
        //     R3_NUM => s.push_str("R' "),
        //     M1_NUM => s.push_str("M "),
        //     M2_NUM => s.push_str("M2 "),
        //     M3_NUM => s.push_str("M' "),
        //     E1_NUM => s.push_str("E "),
        //     E2_NUM => s.push_str("E2 "),
        //     E3_NUM => s.push_str("E' "),
        //     S1_NUM => s.push_str("S "),
        //     S2_NUM => s.push_str("S2 "),
        //     S3_NUM => s.push_str("S' "),
        //     _ => panic!("Invalid move number"),
        // }
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
