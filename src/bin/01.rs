use hashbrown::HashMap;

#[aoc::main(01)]
fn main(input: &str) -> (u32, u32) {
    let mut left_vector: Vec<u32> = Vec::with_capacity(1000);
    let mut right_vector: Vec<u32> = Vec::with_capacity(1000);

    input.lines().for_each(|l| {
        let (l, r) = l.split_once("   ").unwrap();
        left_vector.push(l.parse().unwrap());
        right_vector.push(r.parse().unwrap());
    });

    left_vector.sort();
    right_vector.sort();

    (
        part1(&left_vector, &right_vector),
        part2(&left_vector, &right_vector),
    )
}

fn part1(left: &[u32], right: &[u32]) -> u32 {
    left.iter().zip(right).map(|(l, r)| l.abs_diff(*r)).sum()
}

fn part2(left: &[u32], right: &[u32]) -> u32 {
    let mut frequency = HashMap::with_capacity(1000);

    right
        .iter()
        .for_each(|n| *frequency.entry(n).or_insert(0) += 1);

    left.iter()
        .filter_map(|n| frequency.get(n).map(|f| n * f))
        .sum()
}
