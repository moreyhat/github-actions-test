fn quick_sort(arr: &mut Vec<i32>, low: usize, high: usize) {
    if low < high {
        let mut i = low;
        let pivot = arr[high];
        for j in low..high {
            if arr[j] < pivot {
                arr.swap(i, j);
                i += 1;
            }
        }
        arr.swap(i, high);

        if i > 0 {
            quick_sort(arr, low, i - 1);
        }
        quick_sort(arr, i+1, high);
    }
}

fn main() {
    let mut vector = vec![4,2,8,10,-2];
    println!("Before: {:?}", &vector);
    let length = vector.len();
    quick_sort(&mut vector, 0, length - 1);
    println!("After:  {:?}", &vector);
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_quick_sort() {
        let mut test_cases = vec![
            (vec![], vec![]),
            (vec![1], vec![1]),
            (vec![3,1,2], vec![1,2,3]),
            (vec![1,2,3,4,5], vec![1,2,3,4,5]),
            (vec![5,4,3,2,1], vec![1,2,3,4,5]),
            (vec![1,1,1,1,1], vec![1,1,1,1,1]),
            (vec![4,1,3,2], vec![1,2,3,4]),
            (vec![4,1,3], vec![1,3,4]),
            (vec![1000,2000,3000,4000], vec![1000,2000,3000,4000]),
            (vec![-1,-3,-2,-4,-5], vec![-5,-4,-3,-2,-1]),
            (vec![-1,2,-3,4,-5], vec![-5,-3,-1,2,4]),
            (vec![3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5], vec![1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]),
            (vec![1, 2, 3, 2, 1, 3, 2], vec![1, 1, 2, 2, 2, 3, 3]),
            (vec![5, 3, 9, 1, 6, 7, 2, 4, 8], vec![1, 2, 3, 4, 5, 6, 7, 8, 9]),
            (vec![-100, 0, 100, -50, 50], vec![-100, -50, 0, 50, 100]),
        ];

        for test_case in &mut test_cases {
            let length = test_case.0.len();
            if length > 0 {
                quick_sort(&mut test_case.0, 0, length - 1);
            }
            assert_eq!(test_case.0, test_case.1);
        }
    }
}