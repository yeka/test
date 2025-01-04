fn main() {
    let mut values = vec![1,2];

    values.push(3);
    values.push(4);

    let mut sum = 0;
    for n in values {
        sum = add(sum, n);
    }
    println!("sum = {}", sum);
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}