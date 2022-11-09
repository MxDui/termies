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

