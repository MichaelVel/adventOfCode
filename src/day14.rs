
pub fn solution1(input: &str) -> u32 {
    unimplemented!()
}

pub fn solution2(input: &str) -> u32 {
    unimplemented!()
}

#[cfg(test)]
mod tests {
    use super::*;

	#[ignore]
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
           "a\n\
            a"
        );
    }

}
