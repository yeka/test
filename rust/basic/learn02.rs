fn main() {
    let a = 12;
    let b = 18;
    let mut sum = a; // `mut` indicates the variable is mutable. variable by default is immutable.
    sum += b;

    println!("{} + {} = {}", a, b, sum);
}