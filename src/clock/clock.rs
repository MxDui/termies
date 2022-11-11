use regex::Regex;
use std::collections::HashMap;
use std::env;
use std::process::exit;
use std::time;
// termion grid
use termion::raw::IntoRawMode;
use termion::screen::AlternateScreen;
use termion::terminal_size;


// make a draw clock function
pub fn draw (remain: u64){
    // get the time
    let now = time::SystemTime::now();

    // set til out draw clock will be done
    let until = now + time::Duration::from_secs(remain);

    // get the time left
    let mut remain = remain_to_fmt(remain);

    // print the clock in loop
    loop {
        // get the time left

        // print the clock
        termion::clear::All;

        // print the clock in a grid like this
        // 00:00:00
        // if the time is up
        if time::SystemTime::now() >= until {
            // break the loop
            break;
        }
    }

}


pub(crate) fn remain_to_fmt(remain: u64) -> String {
    let (hours, minutes, seconds) = (remain / 3600, (remain % 3600) / 60, remain % 60);
    if hours == 0 {
        format!("{:02}:{:02}", minutes, seconds)
    } else {
        format!("{:02}:{:02}:{:02}", hours, minutes, seconds)
    }
}
