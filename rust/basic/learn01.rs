fn main() {
    let a = 12;
    let b = 18;
    let sum = a + b;
    println!("{} + {} = {}", a, b, sum);

    let x = (12, 'a');
    println!("Tupple is {} & {}", x.0, x.1);

    let (c, d) = (11, 10);
    let sum2 = c + d;
    println!("{} + {} = {}", c, d, sum2);

    let str_a = String::from("Hello");
    let str_b = str_a;
    // printing str_a will not compile, str_a ownership has been moved to str_b on previous line.
    // so str_a variable is no longer valid/no longer exists
    println!("String is {}", str_b);

    let str_c = str_b.clone(); // by using clone, ownership of str_b isn't moved, so str_b still valid/exists.
    println!("String is {}", str_c);
    println!("String is {}", str_b);
}