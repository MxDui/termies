pub struct Termie {
    pub id: i32,
    pub name: String,
    pub color: String,
}

impl Termie {
    pub fn new(name: String, color: String) -> Self {
        Termie {
            id: 0,
            name,
            color,
        }
    }
}

pub fn print_termie(termie: &Termie) {
    let art = std::fs::read_to_string(format!("./assets/sprites/{}.ans", termie.name)).unwrap();
    println!("{}", art);
}

