use super::Solution;

type Input = Vec<(u64, u64)>;

pub fn run(input: &str) -> (Solution, Solution) {
    let parsed_input = parse(input);

    return (
        Solution::Int64(part1(&parsed_input)),
        Solution::Int64(part2(&parsed_input)),
    );
}

fn parse(input: &str) -> Input {
    return input
        .lines()
        .map(|line| {
            let (left, right) = line.as_bytes().split_at(line.len() / 2);
            (to_bitset(left), to_bitset(right))
        })
        .collect();
}

fn part1(input: &Input) -> i64 {
    return i64::from(
        input
            .iter()
            .map(|&(left, right)| (left & right).trailing_zeros())
            .sum::<u32>(),
    );
}

fn part2(input: &Input) -> i64 {
    return i64::from(
        input
            .chunks_exact(3)
            .map(|chunk| {
                let [(a1, a2), (b1, b2), (c1, c2)]: [_; 3] = chunk.try_into().unwrap();
                ((a1 | a2) & (b1 | b2) & (c1 | c2)).trailing_zeros()
            })
            .sum::<u32>(),
    );
}

fn to_bitset(s: &[u8]) -> u64 {
    s.iter()
        .map(|b| match b {
            b'a'..=b'z' => b - b'a' + 1,
            b'A'..=b'Z' => b - b'A' + 27,
            _ => panic!("Invalid input"),
        })
        .fold(0, |acc, b| acc | (1 << b))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, 157);
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let result = part2(&parsed_input);
        assert_eq!(result, 70);
    }

    fn test_input() -> &'static str {
        return "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
    }
}
