fn main() {
    let a = 18;
    let b = 12;
    let sum = add(a, b);

    println!("{} + {} = {}", a, b, sum);
    println!("{} x {} = {}", a, b, mul(a, b));

    let (res, err) = div(a, b);
    println!("{} / {} = {} {}", a, b, res, err);
}

fn add(a: i32, b: i32) -> i32 {
    return a + b;
}

fn mul(a: i32, b: i32) -> i32 {
    a * b // the return statement is optional
}

fn div(a: i32, b: i32) -> (i32, bool) {
    if b == 0 {
        (0, false)
    } else {
        (a / b, true)
    }
}