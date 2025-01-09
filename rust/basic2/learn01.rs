fn main() {
    example01();
    example02();
    example03();
    example04();
    example05();
    example06();
}

fn hello(a: &mut String) {
    // a.push_str(", hello!");
    *a = "Hello ".to_owned() + a;
}

fn example01() {
    let mut a = String::from("FirstBorn");
    hello(&mut a);
    println!("{}", a);
}

fn example02() {
    let mut a = [1,2,3];
    let b = &mut a[1..];
    b[1] = 5;
    println!("{:?}", a);
}

fn example03() {
    let mut a = 0;
    loop {
        a+=1;
        if a > 10 {
            break;
        }
    }
    println!("loop done");

    // nested loop with label
    a = 0;
    'one: loop {
        'two: loop {
            a += 1;
            if a > 5 {
                println!("exiting loop two");
                break 'two;
            }
        }
        println!("exiting loop one");
        break 'one;
    }
    println!("done");

    //asign value from loop
    a = 0;
    let result = loop {
        a += 1;
        if a>5 {
            break a;
        }
    };
    println!("Result: {}", result);
}

fn example04() {
    for i in 0..5 {
        if i % 2 == 1 {
            continue;
        }
        if i > 2 {
            break;
        }
        println!("{}", i);
    }
}

fn example05() {
    enum MsgState {
        Pending,
        Sending,
        Received
    }

    let msg_state = MsgState::Pending;
    let (_a, _b) = (MsgState::Sending, MsgState::Received);

    match msg_state {
        MsgState::Pending => println!("Pending..."),
        MsgState::Sending => println!("Sending..."),
        MsgState::Received => println!("Received!")
    }

    let status_code = match msg_state {
        MsgState::Pending => 1,
        MsgState::Sending => 2,
        MsgState::Received => 3
    };

    println!("Status code: {}", status_code);
}

fn example06() {
    enum MsgState {
        Pending = 1,
        Sending = 2,
        Received = 3
    }

    let msg_state = MsgState::Received;
    let (_a, _b) = (MsgState::Sending, MsgState::Pending);

    match msg_state {
        MsgState::Pending => println!("Pending..."),
        _ => println!("Not pending -> {}", msg_state as i32),
    }
}