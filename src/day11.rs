
struct Monkey {
    id: i64,
    items: Vec<i64>,
    operation: Box<dyn Fn(i64) -> i64>,
    test: Box<dyn Fn(i64) -> bool>,
    partner1_id: i64,
    partner2_id: i64,
    count: i64,
}

impl Monkey {
    pub fn from(monkey_data: &str) -> Self {
        let lines: Vec<&str> = monkey_data.lines().collect();
        
        let (_,monkey_id) = lines[0].split_once(' ').unwrap();
        let id = monkey_id.trim_matches(':').parse::<i64>().unwrap();
        
        let (_,items) = lines[1].split_once(':').unwrap();
        let items: Vec<i64> = items
            .split(',')
            .map(|x| x.trim().parse::<i64>().unwrap())
            .collect();
        
        let (_,operation) = lines[2].split_once("= ").unwrap();
        let operation_tokens: Vec<&str> = operation.split(' ').collect();
        let right_token = Box::new(operation_tokens[2].parse::<i64>());
        let right_token = Box::leak(right_token);
        let operation: Box<dyn Fn(i64) -> i64> = match operation_tokens[1] {
            "*" if right_token.is_err() => {
                Box::new(|n: i64| n * n)
            },
            "*" if right_token.is_ok() =>  {
                Box::new(|n: i64| n * right_token.as_ref().unwrap())
            },
            "+" if right_token.is_err() => {
                Box::new(|n: i64| n + n)
            },
            "+" if right_token.is_ok() => {
                Box::new(|n: i64| n + right_token.as_ref().unwrap())
            },
            _ => Box::new(|n: i64| n)
        };

        let (_,test) = lines[3].split_once("by ").unwrap();
        let number = Box::new(test.parse::<i64>().unwrap());
        let number = Box::leak(number);
        let test: Box<dyn Fn(i64) -> bool > = Box::new(|x| x % *number== 0);

        let (_,partner1_id) = lines[4].split_once("monkey ").unwrap();
        let partner1_id = partner1_id.parse::<i64>().unwrap();

        let (_,partner2_id) = lines[5].split_once("monkey ").unwrap();
        let partner2_id = partner2_id.parse::<i64>().unwrap();

        Self {
            id,
            items,
            operation,
            test,
            partner1_id,
            partner2_id,
            count: 0,
        }
    }

    pub fn inspect_items(&mut self) -> ((i64,Vec<i64>),(i64,Vec<i64>)) {
        let mut parter1_bag = (self.partner1_id, Vec::<i64>::new());
        let mut parter2_bag = (self.partner2_id, Vec::<i64>::new());
        
        for elem in self.items.iter() {
            let worry_level = self.operation.as_ref()(*elem) / 30;
            if self.test.as_ref()(worry_level) {
                parter1_bag.1.push(worry_level);
            } else {
                parter2_bag.1.push(worry_level);
            }
        }

        self.count += self.items.len() as i64;
        self.items = vec![];
        
        (parter1_bag,parter2_bag)
    }

    pub fn add_to_bag(&mut self, new_items: Vec<i64>) {
        self.items.extend(new_items.iter())
    }
}

pub fn solution1(input: &str) -> i64 {
    let mut monkeys = vec![];

    for monkey_data in input.split("\n\n"){
        let monkey = Monkey::from(monkey_data);
        monkeys.push(monkey);
    }

    for _ in 0..20 {
        for monkey in 0..monkeys.len() {
            let monkey = &mut monkeys[monkey];

            let (partner1, partner2) = monkey.inspect_items();

            let monkey1 = &mut monkeys[partner1.0 as usize];
            monkey1.add_to_bag(partner1.1);

            let monkey2 = &mut monkeys[partner2.0 as usize];
            monkey2.add_to_bag(partner2.1);
        }
    }
    
    let mut counts: Vec<i64> = monkeys.iter().map(|monkey| monkey.count).collect();
    counts.sort_by(|a,b| b.cmp(a));
    
    println!("{:?}", counts);
    return counts[0] * counts[1]
}

pub fn solution2(input: &str) -> i64 {
    let mut monkeys = vec![];

    for monkey_data in input.split("\n\n"){
        let monkey = Monkey::from(monkey_data);
        monkeys.push(monkey);
    }

    for _ in 0..1000 {
        for monkey in 0..monkeys.len() {
            let monkey = &mut monkeys[monkey];

            let (partner1, partner2) = monkey.inspect_items();

            let monkey1 = &mut monkeys[partner1.0 as usize];
            monkey1.add_to_bag(partner1.1);

            let monkey2 = &mut monkeys[partner2.0 as usize];
            monkey2.add_to_bag(partner2.1);
        }
    }
    
    let mut counts: Vec<i64> = monkeys.iter().map(|monkey| monkey.count).collect();
    counts.sort_by(|a,b| b.cmp(a));
    
    println!("{:?}", counts);
    return counts[0] * counts[1]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = get_data();
        let result: i64 = 10605;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = get_data();
        let result: i64 = 2713310158;

        assert_eq!(solution2(&data),result);
    }

    fn get_data() -> String {
        return String::from(
"Monkey 0:
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
