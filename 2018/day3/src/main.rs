use std::io;
use std::io::Read;
use std::process;

fn main() {
    let mut input = String::new();

    io::stdin()
        .read_to_string(&mut input)
        .expect("Failed to read stdin.");

    process::exit(0)
}