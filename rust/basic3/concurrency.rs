use std::thread;
use thread::JoinHandle;
use std::sync::mpsc; // multiple producer single consumer
use std::sync::Mutex;
use std::sync::Arc; // Atomic reference counter
use std::time::Duration;

fn main() {
    simple(); println!();
    simple2(); println!();
    simple3(); println!();
    simple4(); println!();
    mpsc(); println!();
    mpsc2(); println!();
    mutex(); println!();
    mutex2(); println!();
}

fn simple() {
    thread::spawn(|| {
        println!("inside thread!");
    });

    println!("outside thread!");

    // wait for a while to give chance for inside thread to run
    thread::sleep(Duration::from_millis(10));
}

fn simple2() {
    let handle = thread::spawn(|| {
        println!("simple2: inside thread!");
    });

    println!("simple2: outside thread!");

    handle.join().unwrap(); // wait until thread is finish
}

fn simple3() {
    let text = String::from("Hello");
    let handle = thread::spawn(move || {
        println!("variable is: {}", text); // variable text moved into thread and can no longer be accessed outside thread
    });
    handle.join().unwrap();
}

fn simple4() {
    let handle1 = thread::Builder::new().name("My thread".to_string()).spawn(|| {
        println!("Thread 1");
    }).unwrap();

    let handle2 = thread::spawn(|| {
        let th = handle1.thread();
        println!("First thread name & id: {:?}, {:?}", th.name(), th.id());

        handle1.join().unwrap();
        println!("Thread 2");
    });

    // handle1.join().unwrap();
    handle2.join().unwrap();
}

fn mpsc() {
    let (tx, rx) = mpsc::channel();

    let handle = thread::spawn(move || {
        let text = String::from("I'm thread the great!");
        thread::sleep(Duration::from_millis(1000));
        tx.send(text).unwrap(); // text will no longer be available

        thread::sleep(Duration::from_millis(1000));
        println!("thread done");
    });

    println!("Waiting message from thread...");
    let msg = rx.recv().unwrap();
    println!("Got message: {}", msg);

    handle.join().unwrap();
}


fn mpsc2() {
    let (tx, rx) = mpsc::channel();
    let mut handles: Vec<JoinHandle<()>> = vec![];

    for n in 1..6 {
        let mutex = Mutex::new(tx.clone());
        let handle = thread::spawn(move || {
            println!("thread {}", n);
            let text = n.to_string();
            let ntx = mutex.lock().unwrap();
            ntx.send(text).unwrap();
        });
        handles.push(handle);
    }

    let receiver_handle = thread::spawn(move || {
        for s in rx.iter() {
            println!("Got {s}");
        }
    });

    // handle.join().unwrap();
    for h in handles {
        h.join().unwrap();
    }
    drop(tx); // free tx resources so rx can finish the iteration

    println!("sender done");
    receiver_handle.join().unwrap();
    println!("mpsc2 done");
}

fn mutex() {
    let mtx = Mutex::new(42);
    {
        let data = mtx.lock().unwrap();
        println!("Data from mutex: {}", data);
    }
}

fn mutex2() {
    let msg = Arc::new(Mutex::from(String::new()));
    let mut handles: Vec<JoinHandle<()>> = vec![];

    for num in 0..5 {
        let m = Arc::clone(&msg);
        let handle = thread::spawn(move || {
            let mut cur_msg = m.lock().unwrap();

            let mut text = String::from(" Thread: ");
            text.push_str((num + 1).to_string().as_str());
            cur_msg.push_str(&text);
            println!("Message: {}", cur_msg);
        });
        handles.push(handle);
    }

    for h in handles {
        h.join().unwrap();
    }

    println!("Final message: {}", msg.lock().unwrap());
}