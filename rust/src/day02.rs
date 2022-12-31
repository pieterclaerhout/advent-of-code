use super::Solution;

type Input = Vec<(usize, usize)>;

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
        .map(str::as_bytes)
        .map(|line| ((line[0] - b'A') as usize, (line[2] - b'X') as usize))
        .collect();
}

fn part1(input: &Input) -> usize {
    return input
        .iter()
        .copied()
        .map(|(a, x)| {
            let move_score = x + 1;
            let game_result = (3 + x - a + 1) % 3;
            move_score + 3 * game_result
        })
        .sum();
}

fn part2(input: &Input) -> usize {
    return input
        .iter()
        .copied()
        .map(|(a, x)| {
            let move_to_play = (a + x + 3 - 1) % 3;
            let move_score = move_to_play + 1;
            let game_result = x;
            game_result * 3 + move_score
        })
        .sum();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, 15);
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let result = part2(&parsed_input);
        assert_eq!(result, 12);
    }

    fn test_input() -> &'static str {
        return "A Y
B X
C Z";
    }
}
