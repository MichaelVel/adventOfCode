
pub fn solution1(input: &str) -> u32 {
    let mut stack = vec![
        ("/", 0), 
    ];

    let mut total = 0;

    for line in input.lines() {
        if line == "$ cd /" || line == "$ ls" {
            continue;
        }
        
        if line.starts_with("$ cd ") {
            let dir = &line[5..];
            if dir == ".." {
                let (_, amount) = stack.pop().unwrap();
                if amount <= 100_000 {
                    total += amount;
                }
                stack.last_mut().unwrap().1 += amount;
            } else {
                stack.push((dir,0));
            }
            continue;
        }

        let (size, _) = line.split_once(' ').unwrap();
        
        if let Ok(size) = size.parse::<u32>() {
            stack.last_mut().unwrap().1 += size;
        }
    
    }
    return total
}

pub fn solution2(input: &str) -> u32 {
    let mut stack = vec![("/", 0),];
    let mut dirs = vec![];
    
    let total_space = 70_000_000;
    let space_to_delete = 30_000_000;

    for line in input.lines() {
        if line == "$ cd /" || line == "$ ls" {
            continue;
        }
        
        if line.starts_with("$ cd ") {
            let dir = &line[5..];
            if dir == ".." {
                let (name, size) = stack.pop().unwrap();
                stack.last_mut().unwrap().1 += size;
                dirs.push((name,size)); 
            } else {
                stack.push((dir,0));
            }
            continue;
        }

        let (size, _) = line.split_once(' ').unwrap();
        
        if let Ok(size) = size.parse::<u32>() {
            stack.last_mut().unwrap().1 += size;
        }

    }

    while let Some((name,size)) = stack.pop() {
       dirs.push((name,size)); 
       if let Some((_,last_size)) = stack.last_mut() {
           *last_size +=  size;
       }
    }

    let free_space = total_space - dirs.last().unwrap().1;
    let space_required = space_to_delete - free_space;
    
    return dirs
        .into_iter()
        .filter(|(_,amount)| *amount >= space_required)
        .map(|(_,amount)| amount)
        .min()
        .unwrap();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "$ cd /\n\
             $ ls\n\
             dir a\n\
             14848514 b.txt\n\
             8504156 c.dat\n\
             dir d\n\
             $ cd a\n\
             $ ls\n\
             dir e\n\
             29116 f\n\
             2557 g\n\
             62596 h.lst\n\
             $ cd e\n\
             $ ls\n\
             584 i\n\
             $ cd ..\n\
             $ cd ..\n\
             $ cd d\n\
             $ ls\n\
             4060174 j\n\
             8033020 d.log\n\
             5626152 d.ext\n\
             7214296 k"
        );
        
        let result: u32 = 95437;

        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "$ cd /\n\
             $ ls\n\
             dir a\n\
             14848514 b.txt\n\
             8504156 c.dat\n\
             dir d\n\
             $ cd a\n\
             $ ls\n\
             dir e\n\
             29116 f\n\
             2557 g\n\
             62596 h.lst\n\
             $ cd e\n\
             $ ls\n\
             584 i\n\
             $ cd ..\n\
             $ cd ..\n\
             $ cd d\n\
             $ ls\n\
             4060174 j\n\
             8033020 d.log\n\
             5626152 d.ext\n\
             7214296 k"
        );

        let result: u32 = 24933642;

        assert_eq!(solution2(&data),result);
    }

}
