fn main() {
    // closures

    // full definition
    let is_even1 = | val: i32 | -> bool {
        if val % 2 == 0 {
            return true;
        } else {
            return false;
        }
    };

    // closure with inferred return type
    let is_even2 = | val: i32 | {
        val % 2 == 0
    };

    // closure as a single statement
    let is_even3 = | val: i32 | val % 2 == 0;

    let num = 42;
    println!("1. {} is even: {}", num, is_even1(num));
    println!("2. {} is even: {}", num, is_even2(num));
    println!("3. {} is even: {}", num, is_even3(num));
}