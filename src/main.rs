use std::fs::File;
use std::io::prelude::*;

pub mod day1;
pub mod day2;
pub mod day3;
pub mod day4;
pub mod day5;
pub mod day6;
pub mod day7;

fn main() -> std::io::Result<()> {
    let mut file = File::open("./data/day7.txt")?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    let result = day7::solution1(&contents);
    println!("The result is: {result}");

    let result2 = day7::solution2(&contents);
    println!("The second part result is: {result2}");
    Ok(())
}
