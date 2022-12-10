
pub fn solution1(input: &str) -> i32 {
    let mut signal_strength = 0;

    let mut register = 1;
    let mut cycle = 1;

    for line in input.lines() {
        let (instruction, value) = line
            .split_once(' ')
            .unwrap_or((line,"x"));

        if cycle % 40 == 20 {
            signal_strength += cycle*register;
        }

        match instruction{
            "noop" => cycle += 1,
            "addx" => {
                if (cycle+2) % 40 == 21 {
                    signal_strength += (cycle+1)*register;
                }
                register += value.parse::<i32>().unwrap();
                cycle +=2;
            },
            _ => println!("This should not reach")
        }

    }

    return signal_strength
}

fn get_pixel(position: i32, sprite: i32) -> char {
    let sprite = [sprite-1, sprite, sprite+1];
    return if sprite.contains(&position) {'#'} else {'.'}
}

pub fn solution2(input: &str) -> u32 {
    let mut register = 1;
    let mut cycle = 0;

    for line in input.lines() {
        let (instruction, value) = line
            .split_once(' ')
            .unwrap_or((line,"x"));

        if cycle % 40 == 0 {
            print!("{}", '\n');
        }

        match instruction{
            "noop" => {
                print!("{}", get_pixel(cycle % 40, register));
                //println!("{} {}", cycle, register);
                cycle += 1;
            },
            "addx" => {
                print!("{}", get_pixel(cycle % 40, register));
                //println!("{} {}", cycle, register);

                cycle += 1;
                if cycle % 40 == 0 {
                    print!("{}", '\n');
                }
                //println!("{} {}", cycle, register);
                print!("{}", get_pixel(cycle % 40, register));
                cycle +=1;
                register += value.parse::<i32>().unwrap();
            },
            _ => println!("This should not reach")
        }


    }
    print!("\n");

    return 0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "addx 15\n\
             addx -11\n\
             addx 6\n\
             addx -3\n\
             addx 5\n\
             addx -1\n\
             addx -8\n\
             addx 13\n\
             addx 4\n\
             noop\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx -35\n\
             addx 1\n\
             addx 24\n\
             addx -19\n\
             addx 1\n\
             addx 16\n\
             addx -11\n\
             noop\n\
             noop\n\
             addx 21\n\
             addx -15\n\
             noop\n\
             noop\n\
             addx -3\n\
             addx 9\n\
             addx 1\n\
             addx -3\n\
             addx 8\n\
             addx 1\n\
             addx 5\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             addx -36\n\
             noop\n\
             addx 1\n\
             addx 7\n\
             noop\n\
             noop\n\
             noop\n\
             addx 2\n\
             addx 6\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             addx 1\n\
             noop\n\
             noop\n\
             addx 7\n\
             addx 1\n\
             noop\n\
             addx -13\n\
             addx 13\n\
             addx 7\n\
             noop\n\
             addx 1\n\
             addx -33\n\
             noop\n\
             noop\n\
             noop\n\
             addx 2\n\
             noop\n\
             noop\n\
             noop\n\
             addx 8\n\
             noop\n\
             addx -1\n\
             addx 2\n\
             addx 1\n\
             noop\n\
             addx 17\n\
             addx -9\n\
             addx 1\n\
             addx 1\n\
             addx -3\n\
             addx 11\n\
             noop\n\
             noop\n\
             addx 1\n\
             noop\n\
             addx 1\n\
             noop\n\
             noop\n\
             addx -13\n\
             addx -19\n\
             addx 1\n\
             addx 3\n\
             addx 26\n\
             addx -30\n\
             addx 12\n\
             addx -1\n\
             addx 3\n\
             addx 1\n\
             noop\n\
             noop\n\
             noop\n\
             addx -9\n\
             addx 18\n\
             addx 1\n\
             addx 2\n\
             noop\n\
             noop\n\
             addx 9\n\
             noop\n\
             noop\n\
             noop\n\
             addx -1\n\
             addx 2\n\
             addx -37\n\
             addx 1\n\
             addx 3\n\
             noop\n\
             addx 15\n\
             addx -21\n\
             addx 22\n\
             addx -6\n\
             addx 1\n\
             noop\n\
             addx 2\n\
             addx 1\n\
             noop\n\
             addx -10\n\
             noop\n\
             noop\n\
             addx 20\n\
             addx 1\n\
             addx 2\n\
             addx 2\n\
             addx -6\n\
             addx -11\n\
             noop\n\
             noop\n\
             noop"
        );
        let result: i32 = 13140;
        
        assert_eq!(solution1(&data),result);
    }
    
    #[test]
    fn test_data_example2() {
        let data = String::from(
            "addx 15\n\
             addx -11\n\
             addx 6\n\
             addx -3\n\
             addx 5\n\
             addx -1\n\
             addx -8\n\
             addx 13\n\
             addx 4\n\
             noop\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx 5\n\
             addx -1\n\
             addx -35\n\
             addx 1\n\
             addx 24\n\
             addx -19\n\
             addx 1\n\
             addx 16\n\
             addx -11\n\
             noop\n\
             noop\n\
             addx 21\n\
             addx -15\n\
             noop\n\
             noop\n\
             addx -3\n\
             addx 9\n\
             addx 1\n\
             addx -3\n\
             addx 8\n\
             addx 1\n\
             addx 5\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             addx -36\n\
             noop\n\
             addx 1\n\
             addx 7\n\
             noop\n\
             noop\n\
             noop\n\
             addx 2\n\
             addx 6\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             noop\n\
             addx 1\n\
             noop\n\
             noop\n\
             addx 7\n\
             addx 1\n\
             noop\n\
             addx -13\n\
             addx 13\n\
             addx 7\n\
             noop\n\
             addx 1\n\
             addx -33\n\
             noop\n\
             noop\n\
             noop\n\
             addx 2\n\
             noop\n\
             noop\n\
             noop\n\
             addx 8\n\
             noop\n\
             addx -1\n\
             addx 2\n\
             addx 1\n\
             noop\n\
             addx 17\n\
             addx -9\n\
             addx 1\n\
             addx 1\n\
             addx -3\n\
             addx 11\n\
             noop\n\
             noop\n\
             addx 1\n\
             noop\n\
             addx 1\n\
             noop\n\
             noop\n\
             addx -13\n\
             addx -19\n\
             addx 1\n\
             addx 3\n\
             addx 26\n\
             addx -30\n\
             addx 12\n\
             addx -1\n\
             addx 3\n\
             addx 1\n\
             noop\n\
             noop\n\
             noop\n\
             addx -9\n\
             addx 18\n\
             addx 1\n\
             addx 2\n\
             noop\n\
             noop\n\
             addx 9\n\
             noop\n\
             noop\n\
             noop\n\
             addx -1\n\
             addx 2\n\
             addx -37\n\
             addx 1\n\
             addx 3\n\
             noop\n\
             addx 15\n\
             addx -21\n\
             addx 22\n\
             addx -6\n\
             addx 1\n\
             noop\n\
             addx 2\n\
             addx 1\n\
             noop\n\
             addx -10\n\
             noop\n\
             noop\n\
             addx 20\n\
             addx 1\n\
             addx 2\n\
             addx 2\n\
             addx -6\n\
             addx -11\n\
             noop\n\
             noop\n\
             noop"
        );
        let result: u32 = 13140;

        assert_eq!(solution2(&data),result);
    }

}
