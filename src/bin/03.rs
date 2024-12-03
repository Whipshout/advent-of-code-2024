use regex::Regex;

#[aoc::main(03)]
fn main(input: &str) -> (u32, u32) {
    (part1(input), part2(input))
}

fn part1(input: &str) -> u32 {
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();

    re.captures_iter(input).fold(0, |sum, val| {
        let x: u32 = val[1].parse().unwrap();
        let y: u32 = val[2].parse().unwrap();
        sum + x * y
    })
}

fn part2(input: &str) -> u32 {
    let re = Regex::new(r"(do\(\))|(don't\(\))|mul\((\d+),(\d+)\)").unwrap();

    re.captures_iter(input)
        .fold((0, true), |(sum, active), val| {
            if let Some(_) = val.get(1) {
                (sum, true)
            } else if let Some(_) = val.get(2) {
                (sum, false)
            } else if active {
                let x: u32 = val[3].parse().unwrap();
                let y: u32 = val[4].parse().unwrap();
                (sum + x * y, active)
            } else {
                (sum, active)
            }
        })
        .0
}
