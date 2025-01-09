trait Brew {
    fn brew(&self) -> String {
        "Brewing...".to_string()
    }
}

struct Coffee {}
impl Brew for Coffee {}

struct Beer{}
impl Brew for Beer {
    fn brew(&self) -> String {
        "Brewing beer...".to_string()
    }
}

fn brew_drink(drink: impl Brew) {
    println!("{}", drink.brew());
}

fn main() {
    let a = Coffee{};
    brew_drink(a);

    let b = Beer{};
    brew_drink(b);
}