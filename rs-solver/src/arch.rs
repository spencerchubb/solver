pub fn check_32_bit() {
    if std::mem::size_of::<usize>() == 4 {
        panic!("This program is not compatible with 32-bit architectures.");
    }
}