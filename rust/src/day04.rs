use super::Solution;

type Input = Vec<(usize, usize, usize, usize)>;

pub fn run(input: &str) -> (Solution, Solution) {
    let parsed_input = parse(input);

    return (
        Solution::Usize(part1(&parsed_input)),
        Solution::Usize(part2(&parsed_input)),
    );
}

fn parse(input: &str) -> Input {
    return input
        .lines()
        .map(|line| {
            let (a, b) = line.split_once(',').expect("Invalid input");
            let (a1, a2) = a.split_once('-').expect("Invalid input");
            let (b1, b2) = b.split_once('-').expect("Invalid input");
            let a1 = a1.parse().expect("Invalid input");
            let a2 = a2.parse().expect("Invalid input");
            let b1 = b1.parse().expect("Invalid input");
            let b2 = b2.parse().expect("Invalid input");
            return (a1, a2, b1, b2);
        })
        .collect();
}

fn part1(input: &Input) -> usize {
    return input
        .iter()
        .filter(|(a1, a2, b1, b2)| (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2))
        .count();
}

fn part2(input: &Input) -> usize {
    return input
        .iter()
        .filter(|(a1, a2, b1, b2)| a1 <= b2 && a2 >= b1)
        .count();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, 2);
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let result = part2(&parsed_input);
        assert_eq!(result, 4);
    }

    fn test_input() -> &'static str {
        return "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";
    }
}
