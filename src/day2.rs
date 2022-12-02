
struct Round {
    my_move: char,
    adversary: char,
}

impl Round {
    fn from(round: &str) -> Self {
        Round {
           my_move: round.chars().nth(2).unwrap_or('N'),
           adversary: round.chars().nth(0).unwrap_or('N'),
        }
    }

    fn from_strategy(round: &str) -> Self {
        let adversary = round.chars().nth(0).unwrap_or('N');
        let my_move = 
            match (adversary, round.chars().nth(2).unwrap_or('N')) {
              ('A', 'Y') | ('B', 'X') | ('C', 'Z')  => 'X',
              ('A', 'Z') | ('B', 'Y') | ('C', 'X')  => 'Y',
              ('A', 'X') | ('B', 'Z') | ('C', 'Y')  => 'Z',
               _ => 'N' 
        };

        Round {
            my_move,
            adversary,
        }
    }

    fn score(&self) -> u32 {
        match  (self.adversary, self.my_move) {
            ('A', 'X') => 1 + 3,
            ('A', 'Y') => 2 + 6,
            ('A', 'Z') => 3 + 0,
            ('B', 'X') => 1 + 0,
            ('B', 'Y') => 2 + 3,
            ('B', 'Z') => 3 + 6,
            ('C', 'X') => 1 + 6,
            ('C', 'Y') => 2 + 0,
            ('C', 'Z') => 3 + 3,
            _          => 0,
        }
    }
}

pub fn solution1(input: &str) -> u32 {
    let mut score = 0;
    for round_input in input.split("\n") {
        let round = Round::from(round_input);
        score += round.score();
    }
    return score 
}

pub fn solution2(input: &str) -> u32 {
    let mut score = 0;
    for round_input in input.split("\n") {
        let round = Round::from_strategy(round_input);
        score += round.score();
    }
    return score 
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "A Y\n\
             B X\n\
             C Z"
        );
        let result: u32 = 15;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "A Y\n\
             B X\n\
             C Z"
        );
        let result: u32 = 12;

        assert_eq!(solution2(&data),result);
    }

}
