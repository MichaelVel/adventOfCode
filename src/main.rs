use std::fs::File;
use std::io::prelude::*;

pub mod day1;

fn main() -> std::io::Result<()> {
    let mut file = File::open("./data/day1_1.txt")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let result = day1::solution1(&contents);
    let result2 = day1::solution2(&contents);
    println!("The result is: {result}");
    println!("The second part result is: {result2}");
    Ok(())
}
