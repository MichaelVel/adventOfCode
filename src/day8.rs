fn parse_matrix(input: &str) -> Vec<Vec<u32>> {
    let mut output = Vec::new();
    for line in input.lines() {
        let row: Vec<u32> = line
            .chars()
            .map(|x| x.to_digit(10).unwrap())
            .collect();
        output.push(row);
    }
    return output;
}

fn transpose_matrix(matrix: &Vec<Vec<u32>>) -> Vec<Vec<u32>> {
    let mut transposed = Vec::new();

    let col_len = matrix.len();
    let row_len = matrix[0].len();
    
    for i in 0..col_len {
        let mut transposed_row = Vec::new();
        for j in 0..row_len {
            let value = matrix[j][i];
            transposed_row.push(value);
        }
        transposed.push(transposed_row);
    }
    return transposed
}

fn is_visible((y,x): (usize,usize), row: &Vec<u32>, col: &Vec<u32>) -> bool {
    let value = row[x];

    let top = *col[..y].iter().max().unwrap();
    if top < value { return true }

    let bottom = *col[(y+1)..].iter().max().unwrap();
    if bottom < value { return true }

    let left = *row[..x].iter().max().unwrap();
    if left < value { return true }

    let right = *row[(x+1)..].iter().max().unwrap();
    if right < value { return true }

    return false
}

fn n_visible_trees(value: u32, trees: &[u32]) -> u32 {
    let mut trees = trees.iter();
    let mut n = 0;
    while let Some(tree) =  trees.next() {
       n += 1; 
       if *tree >= value {
           break;
       }
    }
    return n
}

fn scenic_score((y,x): (usize,usize), row: &Vec<u32>, col: &Vec<u32>) -> u32 {
    let value = row[x];
    let left: Vec<u32> = row[..x].iter().rev().map(|x| *x).collect();
    let top: Vec<u32> = col[..y].iter().rev().map(|x| *x).collect();

    let top = n_visible_trees(value,&top);
    let bottom = n_visible_trees(value,&col[(y+1)..]);
    let left = n_visible_trees(value,&left);
    let right = n_visible_trees(value,&row[(x+1)..]);

    return top * bottom * left * right
}

pub fn solution1(input: &str) -> u32 {
    let matrix = parse_matrix(input);
    let transpose = transpose_matrix(&matrix);

    let col_len = matrix.len() - 1;
    let row_len = matrix[0].len() - 1;

    let mut n_visible_trees = (row_len * 2) + (col_len * 2);

    for i in 1..col_len {
        for j in 1..row_len {
            if is_visible((i,j), &matrix[i], &transpose[j]) {
                n_visible_trees += 1;
            }
        }
    }

    return n_visible_trees as u32
}

pub fn solution2(input: &str) -> u32 {
    let matrix = parse_matrix(input);
    let transpose = transpose_matrix(&matrix);

    let col_len = matrix.len() - 1;
    let row_len = matrix[0].len() - 1;

    let mut scenic_scores = Vec::new();

    for i in 1..col_len {
        for j in 1..row_len {
            let scenic_score = scenic_score((i,j), &matrix[i], &transpose[j]);
            scenic_scores.push(scenic_score);
        }
    }

    return *scenic_scores.iter().max().unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_data_example1() {
        let data = String::from(
            "30373\n\
             25512\n\
             65332\n\
             33549\n\
             35390"
        );
        let result: u32 = 21;
        
        assert_eq!(solution1(&data),result);
    }

    #[test]
    fn test_data_example2() {
        let data = String::from(
            "30373\n\
             25512\n\
             65332\n\
             33549\n\
             35390"
        );
        let result: u32 = 8;

        assert_eq!(solution2(&data),result);
    }

}
