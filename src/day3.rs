use std::collections::HashSet;

pub fn solution1(input: &str) -> u32 {
    let mut sum = 0;
    for backpack in input.split('\n') {
        let cmpt_len = backpack.len() / 2;
        let first_cmpt: HashSet<&u8> = backpack[..cmpt_len]
            .as_bytes()
            .into_iter()
            .collect();

        let second_cmpt: HashSet<&u8> = backpack[cmpt_len..]
            .as_bytes()
            .into_iter()
            .collect();
        
        for &&error in first_cmpt.intersection(&second_cmpt) {
            let mut val = 0;
            if (error as char).is_ascii_uppercase() {
                val = error - 38;
            } else if (error as char).is_ascii_lowercase() {
                val = error - 96;
            }
            sum += val as u32; 
        }

    }
    return sum
}

pub fn solution2(input: &str) -> u32 {
    let mut sum = 0;
    let backpacks: Vec<&str> = input.split('\n').collect();
    for group_backpacks in backpacks.chunks(3) {
        let group_backpacks: Vec<HashSet<u8>> = group_backpacks
            .into_iter()
            .map( |x| {
                x.as_bytes().iter().cloned().collect::<HashSet<u8>>()
            })
            .collect();

        let group_backpacks = group_backpacks
            .iter()
            .skip(1)
            .fold(group_backpacks[0].clone(), |acc,x| {
                acc.intersection(x).cloned().collect()
            });

        let badget = group_backpacks.into_iter().next().unwrap_or(0);
        let mut val = 0;
        if (badget as char).is_ascii_uppercase() {
            val = badget - 38;
        } else if (badget as char).is_ascii_lowercase() {
            val = badget- 96;
        }
        sum += val as u32; 
    }
    return sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "vJrwpWtwJgWrhcsFMMfFFhFp\n\
             jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n\
             PmmdzqPrVvPwwTWBwg\n\
             wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n\
             ttgJtRGJQctTZtZT\n\
             CrZsJsPPZsGzwwsLwLmpwMDw"
        );
        let result: u32 = 157;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "vJrwpWtwJgWrhcsFMMfFFhFp\n\
             jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n\
             PmmdzqPrVvPwwTWBwg\n\
             wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n\
             ttgJtRGJQctTZtZT\n\
             CrZsJsPPZsGzwwsLwLmpwMDw"
        );
        let result: u32 = 70;

        assert_eq!(solution2(&data),result);
    }

}
