#[derive(Debug)]
struct Coffee<'a, T, U> {
    name: &'a str,
    cost: T,
    size: U,
}

fn main() {
    let c1 = Coffee{name: "Drip", cost: 4.99, size: 3};
    let c2 = Coffee{name: "Drip", cost: 3.50, size: "Grande"};

    println!("{:?}", c1);
    println!("{:?}", c2);
}