#[derive(Debug)] // to be able to print struct using {:?} format
struct Accumulator {
    sum: i32
}

impl Accumulator {
    fn new(init: i32) -> Accumulator { // the Accumulator can also be referred as Self
        Accumulator{sum: init} // inferred return
    }

    fn new2(sum: i32) -> Self {
        Self{sum} // value is inferred because the variable have the same name
    }

    fn get(self) -> i32 {
        self.sum
    }

    fn add(self, num: i32) -> Self {
        Self{sum: self.sum + num} // returning new Accumulator
    }
}


fn main() {
    let acc1 = Accumulator::new(0);
    let mut acc2 = Accumulator::new2(2);
    println!("acc1.sum = {}", acc1.sum);

    println!("acc2 = {:?}", acc2);
    println!("acc2.get = {}", acc2.get()); // <-- not working
    acc2 = acc2.add(10);
    println!("acc2.get after = {}", acc2.get());
}