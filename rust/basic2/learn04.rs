mod coffee {
    #[derive(Debug)]
    pub enum TempCategory {
        HOT,
        ICED,
    }

    #[derive(Debug, Copy, Clone)]
    pub enum Roast {
        DARK,
        LIGHT,
        MEDIUM,
    }

    #[derive(Debug)]
    pub struct Coffee {
        pub name: String,
        pub temp: TempCategory,
        pub roast: Roast,
        cost: f64,
        count: i32,
        total: f64,
    }

    impl Coffee {
        pub fn new(name: String, temp: TempCategory, roast: Roast, cost: f64, count: i32) -> Coffee {
            Coffee{name, temp, roast, cost, count, total: cost * count as f64}
        }

        pub fn get_total(&self) -> f64 {
            self.total
        }

        pub fn set_cost(&mut self, cost: f64) {
            self.cost = cost;
            self.total = self.cost * self.count as f64;
        }

        pub fn set_count(&mut self, count: i32) {
            self.count = count;
            self.total = self.cost * self.count as f64;
        }
    }

    pub trait Drink {
        fn drink(&self) {
            println!("Drinking");
        }
    }

    pub trait Brew: Drink {
        fn brew(&self) {
            println!("Brewing")
        }
    }

    impl Brew for Coffee{}
    impl Drink for Coffee{
        fn drink(&self) {
            println!("Drinking {} that is {:?} and {:?}", self.name, self.temp, self.roast);
        }
    }

}

use coffee::*;

fn main() {
    let mut x = Coffee::new("Espresso".to_string(), TempCategory::HOT, Roast::DARK, 800.0, 2);
    println!("Coffee {} total cost is {}", x.name, x.get_total());
    x.set_cost(200.0);
    x.set_count(5);
    println!("Coffee {} total cost is {}", x.name, x.get_total());

    fn brew_and_drink(drink: impl Brew) {
        drink.brew();
        drink.drink();
    }

    brew_and_drink(x);
}