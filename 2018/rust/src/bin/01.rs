use std::fs;
use std::collections::HashSet;

fn main() {
    let mut freq = 0;

    let input = fs::read_to_string("input/01.txt").unwrap();
    let input = input.split('\n');
    let input = input.filter(|line| line.len() > 1);
    let input_2nd = input.clone();

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

    let mut seen: HashSet<isize> = HashSet::new();
    let mut seen_twice: isize = 0;

    while seen_twice == 0 {
        let input_2nd = input_2nd.clone();

        for line in input_2nd {
            let chars = line.chars().map(|c| c.to_string()).collect::<Vec<String>>();
            let new_freq: isize = chars[1..].join("").parse().unwrap();

            if chars[0] == "+" {
                freq += new_freq;
            } else {
                freq -= new_freq;
            }

            if seen.contains(&freq) && seen_twice == 0 {
                seen_twice = freq;
                break
            }

            seen.insert(freq);
        };
    }

    println!("{}", seen_twice);
}
