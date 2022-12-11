
struct Monkey {
    id: i32,
    items: Vec<i32>,
    operation: fn(i32) -> i32,
    test: fn(i32) -> bool,
    partner1_id: i32,
    partner2_id: i32,
}

impl Monkey {
    pub fn from(monkey_data: &str) -> Self {
        let lines: Vec<&str> = monkey_data.lines().collect();

        let (_,monkey_id) = lines[0].split_once(' ').unwrap();
        let monkey_id = monkey_id.trim_matches(':').parse::<i32>().unwrap();
        
        let (_,items) = lines[1].split_once(':').unwrap();
        let items: Vec<i32> = items
            .split(',')
            .map(|x| x.trim().parse::<i32>().unwrap())
            .collect();
        
        let (_,operation) = lines[2].split_once("= ").unwrap();
        let operation_tokens: Vec<&str> = operation.split(' ').collect();
        let right_token = operation_tokens[2].parse::<i32>();
        let operation = match operation_tokens[1] {
            "*" if right_token.is_err() =>  |n: i32| n * n,
            // have to use trait objects, type to learn some dynamic concepts of rust

            //"*" if right_token.is_ok() =>  |n: i32| n * right_token.unwrap(),
            _ => |n|  n
        };

        unimplemented!()
    }
}

pub fn solution1(input: &str) -> u32 {
    let mut monkeys = vec![];

    for monkey_data in input.split("\n\n"){
        let monkey = Monkey::from(monkey_data);
        monkeys.push(monkey);
    }
    unimplemented!()
}

pub fn solution2(input: &str) -> u32 {
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = get_data();
        let result: u32 = 24000;
        
        assert_eq!(solution1(&data),result);
    }

    #[ignore] 
    #[test]
    fn test_data_example2() {
        let data = get_data();
        let result: u32 = 45000;

        assert_eq!(solution2(&data),result);
    }

    fn get_data() -> String {
        return String::from(
           "
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1"
        );
    }

}
