use std::fs;
use std::collections::BTreeMap;

use chrono;

fn pull_input() -> BTreeMap<chrono::NaiveDateTime, String> {
    let input = fs::read_to_string("input/04.txt").unwrap();
    let input = input.split("\n");

    let mut mapped = BTreeMap::new();
    input.for_each(|line| {
        if line.len() < 2 {
            return
        }
        let clean = line.clone();
        
        let right_split = line.split("]").collect::<Vec<&str>>();
        let lhs = right_split[0];
        let lhs = &lhs[1..]; // the first bracket is ignored
        
        let clean = clean.split("] ").collect::<Vec<&str>>();
        let rhs = clean[1];

        let time = if let Ok(val) = chrono::NaiveDateTime::parse_from_str(lhs, "%Y-%m-%d %H:%M") {
            val
        } else {
            return
        };
        mapped.insert(time, rhs.to_string());
    });
    mapped
}

/*
fn guard_sleep_minutes(map: BTreeMap::<chrono::NaiveDateTime, String>) -> Vec<String> {
    let mut last_guard = String::new();
    let mut mins_asleep = 
    map.iter().map(|(k, v)| {

    }).collect::<Vec<String>>()
}
*/

fn main() {
    let input = pull_input();
    println!("{:#?}", input);
}
