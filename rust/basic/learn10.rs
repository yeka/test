fn main() {
    let sum = add(10, -11);
    println!("sum = {}", sum);
}

// complicated add function
fn add(n1: i32, n2: i32) -> i32 {
    let mut sum = n1;
    let (count, increment) = if n2 > 0 { (n2, 1) } else { (-n2, -1)};
    for _ in 0..count {
        sum += increment;
    }
    sum
}