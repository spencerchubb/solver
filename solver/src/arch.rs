pub fn check_32_bit() -> bool {
    if std::mem::size_of::<usize>() == 4 {
        println!("This program is not compatible with 32-bit architectures.");
        return true;
    }
    false
}