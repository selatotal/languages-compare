use std::thread;

fn main() {

    thread::spawn(move || {
        println!("Hello Thread!");
    });

    println!("Main function");
}
