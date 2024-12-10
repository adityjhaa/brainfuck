use std::io::{self, Read, Write};
use std::env;
use std::fs;

const MAX_SIZE: usize = 30000;

struct Interpreter {
    tape: Vec<u8>,
    ptr: usize,
}

impl Interpreter {
    fn new() -> Self {
        Interpreter {
            tape: vec![0; MAX_SIZE],
            ptr: 0,
        }
    }

    fn interpret(&mut self, code: &str) -> io::Result<()> {
        let code_chars: Vec<char> = code.chars().collect();
        let mut i = 0;

        while i < code_chars.len() {
            match code_chars[i] {
                '>' => {
                    self.ptr = (self.ptr + 1) % MAX_SIZE;
                }
                '<' => {
                    self.ptr = (self.ptr - 1) % MAX_SIZE;
                }
                '+' => {
                    self.tape[self.ptr] = self.tape[self.ptr].wrapping_add(1);
                }
                '-' => {
                    self.tape[self.ptr] = self.tape[self.ptr].wrapping_sub(1);
                }
                '.' => {
                    print!("{}", self.tape[self.ptr] as char);
                    io::stdout().flush()?;
                }
                ',' => {
                    let mut input = [0];
                    io::stdin().read_exact(&mut input)?;
                    self.tape[self.ptr] = input[0];
                }
                '[' => {
                    if self.tape[self.ptr] == 0 {
                        let mut depth = 1;
                        while depth > 0 {
                            i += 1;
                            match code_chars[i] {
                                '[' => depth += 1,
                                ']' => depth -= 1,
                                _ => {}
                            }
                        }
                    }
                }
                ']' => {
                    if self.tape[self.ptr] != 0 {
                        let mut depth = 1;
                        while depth > 0 {
                            i -= 1;
                            match code_chars[i] {
                                ']' => depth += 1,
                                '[' => depth -= 1,
                                _ => {}
                            }
                        }
                    }
                }
                _ => {}
            }
            i += 1;
        }
        Ok(())
    }
}

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    if args.len() != 2 {
        eprintln!("Usage: {} <file>", args[0]);
        std::process::exit(1);
    }
    let filename = &args[1];
    if !filename.ends_with(".bf") {
        eprintln!("Invalid file extension");
        std::process::exit(1);
    }

    let code = fs::read_to_string(filename)?;

    let mut interpreter = Interpreter::new();
    interpreter.interpret(&code)?;

    Ok(())
}

