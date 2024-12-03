use regex::Regex;

#[aoc::main(03)]
fn main(input: &str) -> (i32, i32) {
    (part1(input), part2(input))
}

fn part1(input: &str) -> i32 {
    let re = Regex::new(r"mul\((\d+),(\d+)\)").unwrap();

    re.captures_iter(input).fold(0, |sum, val| {
        let x: i32 = val[1].parse().unwrap();
        let y: i32 = val[2].parse().unwrap();
        sum + x * y
    })
}

fn part2(input: &str) -> i32 {
    let re = Regex::new(r"(do\(\))|(don't\(\))|mul\((\d+),(\d+)\)").unwrap();

    let mut active = true;

    re.captures_iter(input).fold(0, |sum, val| {
        if let Some(_) = val.get(1) {
            active = true;
            sum
        } else if let Some(_) = val.get(2) {
            active = false;
            sum
        } else if active {
            let x: i32 = val[3].parse().unwrap();
            let y: i32 = val[4].parse().unwrap();
            sum + x * y
        } else {
            sum
        }
    })
}
