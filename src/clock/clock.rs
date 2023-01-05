use regex::Regex;
use std::collections::HashMap;
use std::env;
use std::io::Write;
use std::process::exit;
// termion grid
use std::{thread, time};
use termion::{clear, color, cursor, style};

pub fn timer(seconds: u64) {
    let duration = time::Duration::from_secs(seconds);

    // Get the current time
    let start = time::Instant::now();

    // Run a loop until the elapsed time is greater than or equal to the duration
    while start.elapsed() < duration {
        // Calculate the remaining time
        let remaining = duration - start.elapsed();

        // Clear the whole terminal
        print!("{}", clear::All);

        // Move the cursor to the beginning of the line
        print!("{}", cursor::Goto(1, 1));

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
}
