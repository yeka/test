fn main() {
    println!("iterator example 1");
    interator1();
    println!();

    println!("iterator example 2");
    interator2();
    println!();

    println!("iterator example 3");
    iterator3();
    println!();

    println!("iterator example 4");
    iterator4();
    println!();
}

fn interator1() {
    let _non_used_variable = 0; // without underscore, the compiler will complain

    let numbers = (1..5).inspect(|n| println!("n = {}", n)); // code inside inspect is called when numbers is iterated.

    for n in numbers {
        println!("for n = {}", n);
    }
}

fn interator2() {
    // the filter & inspect order is matter. try inspect first, then filter, or vice versa.
    let numbers = (1..5)
    .inspect(|n| println!("first inspect n = {}", n)) // code inside inspect is called when numbers is iterated.
    .filter(|n| { println!("filter n = {}", n); n % 2 == 0 }) // for loop will only result even numbers
    .inspect(|n| println!("second inspect n = {}", n))
    ;

    for n in numbers {
        println!("for n = {}", n);
    }
}

fn iterator3() {
    let sum = (1..5)
    .fold(0, |tally, n| {
        println!("n = {}, tally = {}", n, tally);
        return tally + n;
    });

    println!("sum = {}", sum);
}


fn iterator4() {
    let sum = (1..5).sum::<i32>();
    println!("sum = {}", sum);

    let sum2: i32 = (1..5).sum();
    println!("sum2 = {}", sum2);
}