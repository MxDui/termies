use regex::Regex;
use std::collections::HashMap;
use std::env;
use std::io::Write;
use std::process::exit;
// termion grid
use image::{self, DynamicImage};
use std::{thread, time};
use termion::{clear, color, cursor, style};
use viuer::{print_from_file, Config};

pub fn timer(seconds: u64) {
    let duration = time::Duration::from_secs(seconds);

    // Get the current time
    let start = time::Instant::now();

    // Get the size of the terminal
    let (term_width, term_height) = termion::terminal_size().unwrap();

    let conf = Config {
        width: Some(40),
        height: Some(30),
        x: 2,
        y: 2,
        ..Default::default()
    };

    // Calculate the position of the image
    print!("{}", clear::All);

    print_from_file("image.jpg", &conf).expect("Image printing failed.");

    // Run a loop until the elapsed time is greater than or equal to the duration
    while start.elapsed() < duration {
        // Calculate the remaining time
        let remaining = duration - start.elapsed();

        // Clear the whole terminal

        // Move the cursor to the beginning of the line
        print!("{}", cursor::Goto(3, 1));

        // Print the remaining time using the termion crate with a bigger font size and a different color
        write!(
            std::io::stdout(),
            "{}{}{} seconds remaining",
            style::Bold,
            color::Fg(color::Red),
            remaining.as_secs()
        )
        .unwrap();

        // Flush the output so that it is displayed immediately
        std::io::stdout().flush().unwrap();

        // Sleep the current thread for 1 second
        thread::sleep(time::Duration::from_secs(1));
    }
    print!("{}", clear::All);
}
