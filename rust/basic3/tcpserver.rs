use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use std::thread;

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 1024];

    // Read data from the stream
    match stream.read(&mut buffer) {
        Ok(size) => {
            if size > 0 {
                println!("Received: {}", String::from_utf8_lossy(&buffer[..size]));
                // Respond with a simple message
                if let Err(e) = stream.write(b"Hello from server!") {
                    eprintln!("Failed to send response: {}", e);
                }
            }
        },
        Err(e) => {
            eprintln!("Failed to read from stream: {}", e);
        }
    }
}

fn main() -> std::io::Result<()> {
    // Bind the server to address "localhost:8080"
    let listener = TcpListener::bind("127.0.0.1:8080")?;
    println!("Server listening on 127.0.0.1:8080");

    // Loop to accept incoming connections
    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                // Handle each client in a new thread
                thread::spawn(move || {
                    handle_client(stream);
                });
            },
            Err(e) => {
                eprintln!("Failed to accept connection: {}", e);
            }
        }
    }

    Ok(())
}