
pub fn solution1(input: &str) -> u32 {
    let mut count = 0;
    let sections: Vec<u32>= input
        .split([',','\n','-'])
        .map( |x| x.parse::<u32>().unwrap_or(0))
        .collect();

    for section in sections.chunks_exact(4){
        if (section[0] >= section[2] && section[1] <= section[3] ) ||
           (section[2] >= section[0] && section[3] <= section[1] ) {
            count += 1;
        }
    }
    return count
}

pub fn solution2(input: &str) -> u32 {
    let mut count = 0;
    let sections: Vec<u32>= input
        .split([',','\n','-'])
        .map( |x| x.parse::<u32>().unwrap_or(0))
        .collect();

    for section in sections.chunks_exact(4) {
        if (section[0]..=section[1]).contains(&section[2]) || 
           (section[0]..=section[1]).contains(&section[3]) ||
           (section[2]..=section[3]).contains(&section[0]) ||
           (section[2]..=section[3]).contains(&section[1]) {
           count += 1 
        }
    }
    return count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "2-4,6-8\n\
             2-3,4-5\n\
             5-7,7-9\n\
             2-8,3-7\n\
             6-6,4-6\n\
             2-6,4-8\n\
            "
        );
        let result: u32 = 2;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "2-4,6-8\n\
             2-3,4-5\n\
             5-7,7-9\n\
             2-8,3-7\n\
             6-6,4-6\n\
             2-6,4-8\n\
            "
        );
        let result: u32 = 4;

        assert_eq!(solution2(&data),result);
    }

}
