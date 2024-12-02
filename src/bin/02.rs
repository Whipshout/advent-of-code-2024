#[aoc::main(02)]
fn main(input: &str) -> (usize, usize) {
    let lines: Vec<Vec<u32>> = input
        .lines()
        .map(|l| l.split_whitespace().map(|n| n.parse().unwrap()).collect())
        .collect();

    (part1(&lines), part2(&lines))
}

fn part1(data: &Vec<Vec<u32>>) -> usize {
    data.iter().filter(|l| valid_data(&l)).count()
}

fn part2(data: &Vec<Vec<u32>>) -> usize {
    data.iter().filter(|l| valid_data_dampener(&l)).count()
}

// Horrible way to do it, first iteration :')
fn valid_data(data: &[u32]) -> bool {
    let is_increasing = data.windows(2).all(|w| w[0] < w[1]);
    let is_decreasing = data.windows(2).all(|w| w[0] > w[1]);

    if is_increasing || is_decreasing {
        data.windows(2).all(|w| {
            let diff = w[1].abs_diff(w[0]);
            (1..=3).contains(&diff)
        })
    } else {
        false
    }
}

// Even more horrible, but it works XDDDDD
fn valid_data_dampener(data: &[u32]) -> bool {
    if valid_data(data) {
        return true;
    }

    for i in 0..data.len() {
        let mut modified = data.to_vec();
        modified.remove(i);

        if valid_data(&modified) {
            return true;
        }
    }

    false
}
