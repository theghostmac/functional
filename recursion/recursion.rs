fn countdown(number: i32) {
    if number == 0 {
       println!("Blast Off!");
      return;
    }
    println!(number);
    countdown(number - 1)
}

fn main() {
    countdown(10);
}
