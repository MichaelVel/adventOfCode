
pub fn solution1(input: &str) -> u32 {
    let mut max = 0;
    for elf_items in input.split("\n\n") {
        let mut val = 0;
        for item in elf_items.split('\n') {
            val += item.parse::<u32>().unwrap_or(0);
        }
        if val > max {
            max = val;
        }
    }
    return max
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
            "1000\n\
             2000\n\
             3000\n\n\
             4000\n\n\
             5000\n\
             6000\n\n\
             7000\n\
             8000\n\
             9000\n\n\
             10000"
        );
        let result: u32 = 24000;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "1000\n\
             2000\n\
             3000\n\n\
             4000\n\n\
             5000\n\
             6000\n\n\
             7000\n\
             8000\n\
             9000\n\n\
             10000"
        );
        let result: u32 = 45000;

        assert_eq!(solution2(&data),result);
    }

}
