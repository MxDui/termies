pub struct Pets {
    pub name: String,
    pub age: u8,
    pub species: String,
}

 impl PetsC {
    pub fn new(name: String, age: u8, species: String) -> Pets {
        Pets { name, age, species }
    }
    pub fn show(&self) {
        println!("{} is a {} year old {}", self.name, self.age, self.species);
    }
}
