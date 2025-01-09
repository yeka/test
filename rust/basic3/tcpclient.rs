use std::io::{Read, Write};
use std::net::TcpStream;

fn main() -> std::io::Result<()> {
    // Connect to the server on localhost:8080
    let mut stream = TcpStream::connect("127.0.0.1:8080")?;
    println!("Connected to the server!");

    // Send a message to the server
    let message = "Hello from client!";
    stream.write_all(message.as_bytes())?;

    // Read the response from the server
    let mut buffer = [0; 1024];
    let size = stream.read(&mut buffer)?;
    println!("Server response: {}", String::from_utf8_lossy(&buffer[..size]));

    Ok(())
}