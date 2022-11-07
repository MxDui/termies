use std::io::Write;
use std::{fmt, io};
struct Termie {
    ascii: String,
}

impl fmt::Display for Termie {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.ascii)
    }
}

const PUPPER: &str = r#"
         /^ ^\
        / 0 0 \
        V\ Y /V
        / - \
        /    |
        V__) ||
"#;

const KITTEN: &str = r#"
      |\      _,,,---,,_
ZZZzz /,`.-'`'    -.  ;-;;,_
     |,4-  ) )-,_. ,\ (  `'-'
    '---''(_/--'  `-'\_)
"#;

pub fn ascii_pupper() {
    let pupper = Termie {
        ascii: PUPPER.to_string(),
    };
    // print pupper line by line
    for line in pupper.ascii.lines() {
        writeln!(io::stdout(), "{}", line).unwrap();
    }
}

pub fn ascii_kitten() {
    println!("{}", KITTEN);
}

// implement display
