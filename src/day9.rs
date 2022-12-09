use std::collections::HashSet;

fn distance((x1,y1): (i32, i32), (x2,y2): (i32,i32)) -> f32 {
    (((x1-x2).pow(2) + (y1-y2).pow(2)) as f32).sqrt()
}

pub fn solution1(input: &str) -> u32 {
    let mut tail_positions: HashSet<(i32,i32)>= HashSet::new();

    let mut head_pos = (0,0);
    let mut tail_pos = (0,0);

    tail_positions.insert(tail_pos);

    for movement in input.lines() {
        let (direction,steps) = movement.split_once(' ').unwrap();
        for _ in 0..steps.parse().unwrap() {
            let (x, y) = head_pos;
            let (x1, y1) = tail_pos;

            head_pos = match direction {
                "U" => (x,y+1),
                "D" => (x,y-1),
                "L" => (x-1,y),
                "R" => (x+1,y),
                 _  => (x,y)
            };


            tail_pos = match (x-x1,y-y1) {
                (-1, 2) | (-2, 1) => (x1-1, y1+1),          // top-left 
                ( 0, 2)           => (x1  , y1+1),          // top
                ( 1, 2) | ( 2, 1) => (x1+1, y1+1),          // top-right 
                ( 2, 0)           => (x1+1, y1  ),          // right
                ( 2,-1) | ( 1,-2) => (x1+1, y1-1),          // bottom-right
                ( 0,-2)           => (x1  , y1-1),          // bottom
                (-1,-2) | (-2,-1) => (x1-1, y1-1),          // bottom-left
                (-2, 0)           => (x1-1, y1  ),          // left
                _ => (x1,y1)
            };

            tail_positions.insert(tail_pos);
        }
    } 

    println!("{:?}",tail_positions);
    return tail_positions.len() as u32
}

pub fn solution2(input: &str) -> u32 {
    let mut elf_totals = vec![];
    for elf_items in input.split("\n\n") {
        let mut val = 0;
        for item in elf_items.split('\n') {
            val += item.parse::<u32>().unwrap_or(0);
        }
        elf_totals.push(val)
    }
    elf_totals.sort_by(|a,b| b.cmp(a));
    return elf_totals[..3].iter().sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "R 4\n\
             U 4\n\
             L 3\n\
             D 1\n\
             R 4\n\
             D 1\n\
             L 5\n\
             R 2"
        );
        let result: u32 = 13;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "R 4\n\
             U 4\n\
             L 3\n\
             D 1\n\
             R 4\n\
             D 1\n\
             L 5\n\
             R 2"
        );
        let result: u32 = 45000;

        assert_eq!(solution2(&data),result);
    }

}
