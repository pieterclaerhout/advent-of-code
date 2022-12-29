use super::Solution;

pub fn run(input: &str) -> (Solution, Solution) {
    let mut sorted_asc = input
        .split("\n\n")
        .map(|line_chunks| line_chunks.split("\n"))
        .map(|chunk| {
            chunk
                .map(|calorie| calorie.parse::<i64>().unwrap())
                .sum::<i64>()
        })
        .collect::<Vec<i64>>();

    sorted_asc.sort_unstable();
    sorted_asc.reverse();
    sorted_asc.truncate(3);

    return (
        Solution::Int64(*sorted_asc.first().unwrap()),
        Solution::Int64(sorted_asc.iter().sum::<i64>()),
    );
}
