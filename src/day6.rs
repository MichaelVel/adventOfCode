use std::collections::HashSet;

pub fn solution1(input: &str) -> u32 {
    let mut index = 4;
    let input: Vec<char> = input.chars().collect();
    for chrs in input.windows(4){
        let chrs_set: HashSet<char> = chrs.iter().map(|x| *x).collect();
        if chrs_set.len() == 4 {
            break;
        }
        index = index + 1;
    }
    return index
}

pub fn solution2(input: &str) -> u32 {
    let mut index = 14;
    let input: Vec<char> = input.chars().collect();
    for chrs in input.windows(14){
        let chrs_set: HashSet<char> = chrs.iter().map(|x| *x).collect();
        if chrs_set.len() == 14 {
            println!("{:?}",chrs);
            break;
        }
        index = index + 1;
    }
    return index
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let datastreams = vec![
            "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
            "bvwbjplbgvbhsrlpgdmjqwftvncz",
            "nppdvjthqldpwncqszvftbrmjlhg",
            "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
            "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
        ];

        let results = vec![7, 5, 6, 10, 11];
        
        for i in 0..5 {
            assert_eq!(solution1(&datastreams[i]),results[i]);
        }
    }

    #[test]
    fn test_data_example2() {
        let datastreams = vec![
            "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
            "bvwbjplbgvbhsrlpgdmjqwftvncz",
            "nppdvjthqldpwncqszvftbrmjlhg",
            "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
            "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
        ];

        let results = vec![19, 23, 23, 29, 26];
        
        for i in 0..5 {
            assert_eq!(solution2(&datastreams[i]),results[i]);
        }
    }

}
