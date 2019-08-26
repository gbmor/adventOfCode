use std::fs;
use std::collections::HashMap;

fn pull_input() -> Vec<String> {
    let input = fs::read_to_string("input/02.txt").unwrap();
    let input = input.split("\n");
    input.map(|line| line.to_string()).collect::<Vec<String>>()
}

fn main() {
    let mut count_threes: usize = 0;
    let mut count_twos: usize = 0;
    let input = pull_input();

    input.into_iter().for_each(|line| {
        let mut map: HashMap<String, usize> = HashMap::new();
        line.chars().for_each(
            |c| {
                let c = c.to_string();
                if map.contains_key(&c) {
                    let num = map.get(&c).unwrap();
                    let num = num + 1;
                    map.insert(c.to_string(), num);
                } else {
                    map.insert(c.to_string(), 1);
                }
            });

        let mut threes = false;
        let mut twos = false;
        for (c, v) in map.iter() {
            if !threes && *v as usize == 3 {
                threes = true;
                count_threes += 1;
            }
            if !twos && *v as usize == 2 {
                twos = true;
                count_twos += 1;
            }
            if twos && threes {
                break
            }
        }
    });

    println!("{}", count_twos * count_threes);
}
