mod coffee; // import file coffee.rs
mod beer; // import file beer.rs, or beer/mod.rs
mod food; // import food.rs or food/mod.rs

use coffee::Coffee;
use beer::beer::Beer;
use food::*;

fn main() {
    let x = Coffee{name: "Espresso".to_string()};
    println!("{}", x.name);

    let y = Beer{name: "Biru".to_string()};
    println!("{}", y.name);

    let r = rice::Rice{};
    println!("{:?}", r);

    let s = soup::Soup{};
    println!("{:?}", s);
}