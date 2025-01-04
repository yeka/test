fn main() {
    let values = [1,2,3,4];

    let mut sum = 0;
    for n in 0..4 {
        sum += values[n];
    }
    println!("sum = {}", sum);

    let mut sum2 = 0;
    for v in values {
        sum2 += v;
    }
    println!("sum2 = {}", sum2);

    let mut sum3 = 0;
    for n in &values[0..2] {
        sum3 += n;
    }
    println!("sum3 = {}", sum3);

    let mut sum4 = 0;
    for n in &values[2..4] {
        sum4 = add(sum4, *n);
    }
    println!("sum4 = {}", sum4);
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}