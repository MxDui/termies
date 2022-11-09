#![allow(unused)]

use clap::Parser;
use rusqlite::{params, Connection, Result};
extern crate termion;
use std::io::{self, Write};
use termion::cursor::{self, DetectCursorPos};
use termion::event::*;
use termion::input::{MouseTerminal, TermRead};
use termion::raw::IntoRawMode;
#[derive(Parser)]
struct Cli {
    pattern: String,
    color: Option<String>,
    name: Option<String>,
}

fn main() {
    let args = Cli::parse();

    // The first time to run this program, create a file
    if !std::path::Path::new("db.sqlite").exists() {
        let conn = Connection::open("db.sqlite").unwrap();

        conn.execute(
            "CREATE TABLE IF NOT EXISTS history (
                  id              INTEGER PRIMARY KEY,
                  pattern         TEXT NOT NULL
                  )",
            params![],
        )
        .unwrap();

        conn.execute(
            "CREATE TABLE IF NOT EXISTS termies (
                  id              INTEGER PRIMARY KEY,
                  name          TEXT NOT NULL,
                  color           TEXT NOT NULL,
                    history_id      INTEGER NOT NULL,
                    FOREIGN KEY(history_id) REFERENCES history(id)

                  )",
            params![],
        )
        .unwrap();
    } else {
        let conn = Connection::open("db.sqlite").unwrap();
        conn.execute(
            "INSERT INTO history (pattern) VALUES (?1)",
            params![args.pattern],
        )
        .unwrap();
    }

    if args.pattern == "new" {
        // make a new puppet in db
        let conn = Connection::open("db.sqlite").unwrap();

        conn.execute(
            "INSERT INTO termies (name, color, history_id)
                  VALUES (?1, ?2, ?3)",
            params![args.name, args.color, 1],
        )
        .unwrap();
    }

    if args.pattern == "debug" {
        let conn = Connection::open("db.sqlite").unwrap();
        let mut stmt = conn.prepare("SELECT * FROM termies").unwrap();

        let art = std::fs::read_to_string("./assets/sprites/dog.ans").unwrap();

        let art = format!("{}", art);

        println!("\n{art}");
    }
}
