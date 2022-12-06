use std::collections::HashMap;
use std::str;

fn parse_header(header: &str) -> HashMap<u32,Vec<char>> {
    let mut stacks: HashMap<u32, Vec<char>>= HashMap::new();
    let mut header: Vec<&str> = header.split('\n').collect();

    let stacks_headers = header.pop().unwrap();
    stacks_headers
        .split_whitespace()
        .map(|x| x.parse::<u32>().unwrap_or(999))
        .for_each(|x| { stacks.insert(x, Vec::new()); } );


    for line in header.into_iter().rev() {
        line.as_bytes()
            .chunks(4)
            .map(|x| {
                str::from_utf8(x).unwrap().chars().nth(1).unwrap_or(' ')
            })
            .enumerate()
            .for_each(|(i,x)| { 
                if x != ' ' {
                    let key = (i as u32) + 1;
                    stacks.get_mut(&key).unwrap().push(x);
                }
            });

    }
    return stacks;
}

fn parse_instruction(instruction: &str) -> (u32,u32,u32) {
    let instructions: Vec<u32>= instruction
        .split_whitespace()
        .filter( |x| !"move from to".contains(x))
        .map(|x| x.parse::<u32>().unwrap_or(0))
        .collect();
    return (instructions[0],instructions[1],instructions[2])
}

pub fn solution1(input: &str) -> String {
    let mut solution = String::new();
    let (header,body) = input.split_once("\n\n").unwrap();
    let mut stacks = parse_header(header);
    
    for line in body.split_terminator('\n') {
        let (moves, from, to) = parse_instruction(line);
        for _ in 0..moves {
            let item;
            {
                let from_stack = stacks.get_mut(&from).unwrap();
                item = from_stack.pop().unwrap();
            }
            {
                let to_stack = stacks.get_mut(&to).unwrap();
                to_stack.push(item);
            }
        }
    }

    let mut keys: Vec<&u32> = stacks.keys().collect();
    keys.sort();
    
    for key in keys {
        let last_elem_stack = stacks.get(key).unwrap().last().unwrap();
        solution.push(*last_elem_stack);
    }
    return solution
}

pub fn solution2(input: &str) -> String {
    let mut solution = String::new();
    let (header,body) = input.split_once("\n\n").unwrap();
    let mut stacks = parse_header(header);
    
    for line in body.split_terminator('\n') {
        let (moves, from, to) = parse_instruction(line);
        let mut tmp_stack = vec![];

        for _ in 0..moves {
            let from_stack = stacks.get_mut(&from).unwrap();
            let item = from_stack.pop().unwrap();
            tmp_stack.push(item);
        }
        for _ in 0..moves {
            let item = tmp_stack.pop().unwrap();
            let to_stack = stacks.get_mut(&to).unwrap();
            to_stack.push(item);
        }
    }

    let mut keys: Vec<&u32> = stacks.keys().collect();
    keys.sort();
    
    for key in keys {
        let last_elem_stack = stacks.get(key).unwrap().last().unwrap();
        solution.push(*last_elem_stack);
    }
    return solution
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "    [D]    \n\
             [N] [C]    \n\
             [Z] [M] [P]\n\
             1   2   3 \n\n\
             move 1 from 2 to 1\n\
             move 3 from 1 to 3\n\
             move 2 from 2 to 1\n\
             move 1 from 1 to 2\n\
            "
        );
        let result = "CMZ";
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "    [D]    \n\
             [N] [C]    \n\
             [Z] [M] [P]\n\
             1   2   3 \n\n\
             move 1 from 2 to 1\n\
             move 3 from 1 to 3\n\
             move 2 from 2 to 1\n\
             move 1 from 1 to 2\n\
            "
        );
        let result = "MCD";

        assert_eq!(solution2(&data),result);
    }

}
