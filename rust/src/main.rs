use std::env;
use std::fmt;
use std::fs;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;
mod day25;

#[derive(Debug)]
pub enum Solution {
    String(String),
    Int64(i64),
    Usize(usize),
}

impl fmt::Display for Solution {
    fn fmt(&self, fmt: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Solution::String(s) => write!(fmt, "{}", s),
            Solution::Int64(s) => write!(fmt, "{}", s),
            Solution::Usize(s) => write!(fmt, "{}", s),
        }
    }
}

fn main() {
    let day_arg = env::args().nth(1);

    if day_arg.is_none() {
        for day in 1..26 {
            run_day(day)
        }
        return;
    }

    let day: u32 = day_arg.unwrap().parse().unwrap();
    run_day(day);
}

fn run_day(day: u32) {
    println!(">>> Day: {day} <<<");

    let input: String = read_file_for_day(day)
        .replace("\r\n", "\n")
        .replace("\r", "\n")
        .trim_end()
        .to_string();

    let (result1, result2) = match day {
        1 => day01::run(input.as_str()),
        2 => day02::run(input.as_str()),
        3 => day03::run(input.as_str()),
        4 => day04::run(input.as_str()),
        5 => day05::run(input.as_str()),
        25 => day25::run(input.as_str()),
        _ => unreachable!(),
    };

    println!("Day 1: {result1}");
    println!("Day 2: {result2}");
    println!();
}

fn read_file_for_day(day: u32) -> String {
    let input_path: String = format!("../inputs/day{day:0>2}/input.txt");

    let contents = fs::read_to_string(input_path);

    match contents {
        Ok(contents) => contents,
        Err(e) => {
            println!("{}", e);
            panic!("{day} not implemented yet");
        }
    }
}
