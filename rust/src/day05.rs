use super::Solution;

#[path = "helpers.rs"]
mod helpers;

pub use helpers::SliceExt as _;

type Input = (Vec<Vec<u8>>, Vec<(usize, usize, usize)>);

pub fn run(input: &str) -> (Solution, Solution) {
    let parsed_input = parse(input);

    return (
        Solution::String(part1(&parsed_input)),
        Solution::String(part2(&parsed_input)),
    );
}

fn parse(input: &str) -> Input {
    let (start, moves) = input.split_once("\n\n").expect("Invalid input");

    let (start, idxs) = start.rsplit_once('\n').expect("Invalid input");
    let mut start_state = vec![Vec::new(); idxs.split("   ").count()];
    for line in start.lines().rev() {
        for (i, chunk) in line.as_bytes().chunks(4).enumerate() {
            if chunk[1] != b' ' {
                start_state[i].push(chunk[1]);
            }
        }
    }

    let moves = moves
        .lines()
        .map(|line| {
            let (n, rest) = line[5..].split_once(" from ").expect("Invalid input");
            let (from, to) = rest.split_once(" to ").expect("Invalid input");
            let n = n.parse().expect("Invalid input");
            let from = from.parse().expect("Invalid input");
            let to = to.parse().expect("Invalid input");
            (n, from, to)
        })
        .collect();

    return (start_state, moves);
}

fn part1(input: &Input) -> String {
    let (start, moves) = input;
    let mut state = start.clone();

    for &(n, from, to) in moves {
        let (from, to) = state.get2_mut(from - 1, to - 1);
        to.extend(from.drain(from.len() - n..).rev());
    }

    return state
        .into_iter()
        .map(|state| *state.last().expect("Invalid input") as char)
        .collect();
}

fn part2(input: &Input) -> String {
    let (start, moves) = input;
    let mut state = start.clone();

    for &(n, from, to) in moves {
        let (from, to) = state.get2_mut(from - 1, to - 1);
        to.extend(from.drain(from.len() - n..));
    }

    return state
        .into_iter()
        .map(|state| *state.last().expect("Invalid input") as char)
        .collect();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, "CMZ");
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let result = part2(&parsed_input);
        assert_eq!(result, "MCD");
    }

    fn test_input() -> &'static str {
        return "    [D]
[N] [C]
[Z] [M] [P]
  1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
    }
}
