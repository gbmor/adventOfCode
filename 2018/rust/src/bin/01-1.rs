use std::fs;

fn main() {
    let mut freq = 0;

    let input = fs::read_to_string("input/01.txt").unwrap();
    let input = input.split('\n');
    let input = input.filter(|line| line.len() > 1);

    input
        .for_each(|line| {
            let chars = line.chars().map(|c| c.to_string()).collect::<Vec<String>>();
            let new_freq: isize = chars[1..].join("").parse().unwrap();

            if chars[0] == "+" {
                freq += new_freq;
            } else {
                freq -= new_freq;
            }
        });

    println!("{}", freq);
}
