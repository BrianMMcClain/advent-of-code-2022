use std::fs::File;
use std::io::{BufReader, BufRead, Error};

fn main() -> Result<(), Error> {
    let file = File::open("./input.txt")?;
    let reader = BufReader::new(file);

    let mut calories = Vec::new();
    let mut sum = 0;
    for line in reader.lines() {
        let num = line?.parse::<u32>().unwrap_or(0);
        if num == 0 {
            calories.push(sum);
            sum = 0;
        } else {
            sum += num;
        }
    }

    calories.sort();
    calories.reverse();

    println!("Part 1: {}", calories.first().unwrap());
    println!("Part 2: {}", calories.get(0..3).unwrap().iter().sum::<u32>());

    Ok(())
}