use std::fs::{File, remove_file, OpenOptions};
use std::path::Path;
use std::io::{Write, Read, SeekFrom, Seek};

fn main() {
    removing_file();
    writing_file();
    reading_file();
    append_file();
    overwrite_file();
    removing_file();
}

fn writing_file() {
    let file_path = Path::new("file.txt");
    let mut file = File::create(file_path).unwrap(); // todo error handling
    match write!(file, "Hello world!") {
        Err(e) => println!("Writing error: {:?}", e),
        _ => println!("Writing file: ok")
    }
}

fn reading_file() {
    let mut content = String::new();
    let file_path = Path::new("file.txt");
    let mut file = File::open(&file_path).unwrap();
    file.read_to_string(&mut content).unwrap();
    println!("file content: {}", content);
}

fn removing_file() {
    let file_path = Path::new("file.txt");
    let res = remove_file(file_path);
    if res.is_err() {
        println!("Error removing file: {}", res.unwrap_err());
        return;
    }
    println!("file removed");
    // match res {
    //     Err(e) => println!("Error: {:?}", e),
    //     _ => println!("file removed")
    // }
}

fn append_file() {
    let res = OpenOptions::new()
        .read(true)
        .append(true) // also calls write(true)
        .create(true)
        .open("file.txt");

    if res.is_err() {
        println!("Unable to open file: {}", res.unwrap_err());
        return;
    }
    let mut file = res.unwrap();

    let _ = write!(file, " New addition.");

    let _ = file.seek(SeekFrom::Start(0));
    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();
    println!("file content: {}", content);

}

fn overwrite_file() {
    let res = OpenOptions::new()
        .read(true)
        .write(true)
        .truncate(true) // without truncate, it will write from the first byte over the existing content
        .create(true)
        .open("file.txt");

    if res.is_err() {
        println!("Unable to open file: {}", res.unwrap_err());
        return;
    }
    let mut file = res.unwrap();

    let _ = write!(file, "New content.");

    let _ = file.seek(SeekFrom::Start(0));
    let mut content = String::new();
    file.read_to_string(&mut content).unwrap();
    println!("file content: {}", content);

}
