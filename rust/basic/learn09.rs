#[derive(Debug)]
struct Accumulator {
    sum: i32
}

impl Accumulator {
    fn new(init: i32) -> Self {
        Self{sum: init}
    }

    fn get(&self) -> i32 { // immutable borrow of self
        self.sum
    }

    fn add(&mut self, num: i32) { // mutable borrow of self
        self.sum += num
    }

    fn combine(acc1: Self, acc2: Self) -> Self { // Move of both acc1 and acc2
        Self::new(acc1.sum + acc2.sum)
    }
}


fn main() {
    let mut acc = Accumulator::new(0);
    println!("acc = {:?}", acc);
    println!("acc.get = {}", acc.get());
    acc.add(10);
    println!("acc.get after = {}", acc.get());

    let mut acc2 = Accumulator::new(0);
    acc2.add(12);

    let acc3 = Accumulator::combine(acc, acc2); // "Move" consume both acc and acc2, so it will no longer works after
    println!("combined value: {}", acc3.get());
    // println!("acc value: {}", acc.get()); // <-- no longer work, because ownership have been "moved" to combine
}