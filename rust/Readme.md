# Rust

Learn rust using docker container. Enter container using:
`docker run -it --rm -v $(PWD):/app -w /app rust:alpine sh`

Compile using:
`rustc helloworld.rs`

Then run the compiled version:
`./helloworld`

Alternatively, use `run.sh` for compiling, run the compiled version, and then removed the compiled binary.

`./run.sh helloworld.sh`

# Package Manager

Rust use `cargo` as its package manager. It works similarly like `npm` on NodeJS.
You can create new project by running:
`cargo new <path>`

enter the directory, and run the project by running:
`cargo run`
