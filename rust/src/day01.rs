pub fn run(input: &str) -> (i32, i32) {
    // let parsed = parse(input);

    let mut sorted_asc = input
        .split("\n\n")
        .map(|line_chunks| line_chunks.split("\n"))
        .map(|chunk| {
            chunk
                .map(|calorie| calorie.parse::<i32>().unwrap())
                .sum::<i32>()
        })
        .collect::<Vec<i32>>();

    sorted_asc.sort_unstable();
    sorted_asc.reverse();
    sorted_asc.truncate(3);

    return (*sorted_asc.first().unwrap(), sorted_asc.iter().sum());
}
