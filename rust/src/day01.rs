use super::Solution;

type Input = Vec<Vec<usize>>;

pub fn run(input: &str) -> (Solution, Solution) {
    let parsed_input = parse(input);

    return (
        Solution::Usize(part1(&parsed_input)),
        Solution::Usize(part2(&parsed_input)),
    );
}

fn parse(input: &str) -> Input {
    return input
        .split("\n\n")
        .map(|lines| {
            lines
                .lines()
                .map(|line| line.parse().expect("Invalid input"))
                .collect()
        })
        .collect();
}

fn part1(input: &Input) -> usize {
    return input
        .iter()
        .map(|elf| elf.iter().sum::<usize>())
        .max()
        .expect("Invalid input");
}

fn part2(input: &Input) -> usize {
    return input
        .iter()
        .rev()
        .take(3)
        .map(|elf| elf.iter().sum::<usize>())
        .sum::<usize>();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, 24000);
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let result = part2(&parsed_input);
        assert_eq!(result, 45000);
    }

    fn test_input() -> &'static str {
        return "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";
    }
}
