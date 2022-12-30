use super::Solution;

type Input<'a> = Vec<&'a [u8]>;

pub fn run(input: &str) -> (Solution, Solution) {
    let parsed_input = parse(input);

    let sum = part1(&parsed_input);
    let snafu = part2(sum);

    return (Solution::Int64(sum), Solution::String(snafu));
}

fn parse(input: &str) -> Input {
    input.lines().map(str::as_bytes).collect()
}

fn part1(input: &Input) -> i64 {
    return input
        .iter()
        .map(|&line| {
            let mut acc = 0i64;
            for &b in line {
                acc = acc * 5
                    + match b {
                        b'0' => 0,
                        b'1' => 1,
                        b'2' => 2,
                        b'-' => -1,
                        b'=' => -2,
                        _ => panic!("Invalid input"),
                    }
            }
            acc as i64
        })
        .sum::<i64>();
}

fn part2(mut n: i64) -> String {
    let mut out = Vec::new();
    while n != 0 {
        let (carry, character) = match n % 5 {
            0 => (0, '0'),
            1 => (0, '1'),
            2 => (0, '2'),
            3 => (1, '='),
            4 => (1, '-'),
            _ => unreachable!(),
        };
        n = n / 5 + carry;
        out.push(character);
    }
    return out.into_iter().rev().collect();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let parsed_input = parse(test_input());
        let result = part1(&parsed_input);
        assert_eq!(result, 4890);
    }

    #[test]
    fn test_part2() {
        let parsed_input = parse(test_input());
        let sum = part1(&parsed_input);
        let result = part2(sum);
        assert_eq!(result, "2=-1=0");
    }

    fn test_input() -> &'static str {
        return "1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122";
    }
}
