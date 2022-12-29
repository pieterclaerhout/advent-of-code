use super::Solution;

pub fn run(input: &str) -> (Solution, Solution) {
    let result1 = to_dec(input);
    let result2 = to_snafu(result1);

    return (Solution::Int64(result1), Solution::String(result2));
}

fn to_dec(input: &str) -> i64 {
    return input
        .lines()
        .map(|l| {
            l.as_bytes().iter().fold(0, |value, &digit| {
                5 * value
                    + match digit {
                        b'2' => 2,
                        b'1' => 1,
                        b'0' => 0,
                        b'-' => -1,
                        b'=' => -2,
                        _ => unreachable!(),
                    }
            })
        })
        .sum::<i64>();
}

fn to_snafu(sum: i64) -> String {
    let mut digits = Vec::new();
    let mut sum = sum.clone();
    while sum != 0 {
        let v = sum % 5;
        sum /= 5;
        if v > 2 {
            sum += 1;
        }
        digits.push(match v {
            3 => '=',
            4 => '-',
            0 => '0',
            1 => '1',
            2 => '2',
            _ => unreachable!(),
        });
    }
    return digits.iter().rev().collect();
}
