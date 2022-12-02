use std::fs::File;
use std::io::prelude::*;

pub mod day1;
pub mod day2;

fn main() -> std::io::Result<()> {
    let mut file = File::open("./data/day2.txt")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let result = day2::solution1(&contents);
    let result2 = day2::solution2(&contents);
    println!("The result is: {result}");
    println!("The second part result is: {result2}");
    Ok(())
}
