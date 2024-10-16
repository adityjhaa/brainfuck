import sys
import argparse
from typing import List

MEMORY_SIZE = 30000

def interpret(code: str) -> None:
    memory: List[int] = [0] * MEMORY_SIZE
    ptr = 0

    def find_matching_bracket(start: int, direction: int) -> int:
        depth = 0
        while True:
            if code[start] == '[':
                depth += direction
            elif code[start] == ']':
                depth -= direction
            if depth == 0:
                return start
            start += direction

    i = 0
    while i < len(code):
        match code[i]:
            case '>':
                ptr = (ptr + 1) % MEMORY_SIZE
            case '<':
                ptr = (ptr - 1) % MEMORY_SIZE
            case '+':
                memory[ptr] = (memory[ptr] + 1) % 256
            case '-':
                memory[ptr] = (memory[ptr] - 1) % 256
            case '.':
                print(chr(memory[ptr]), end='', flush=True)
            case ',':
                memory[ptr] = ord(sys.stdin.read(1))
            case '[':
                if memory[ptr] == 0:
                    i = find_matching_bracket(i, 1)
            case ']':
                if memory[ptr] != 0:
                    i = find_matching_bracket(i, -1)
            case _:
                pass

        i += 1


def main() -> None:
    parser = argparse.ArgumentParser(description="Brainfuck Interpreter")
    parser.add_argument('file', nargs='?', help="Brainfuck source file")
    args = parser.parse_args()

    if not args.file:
        print("Usage: python3 compiler.py <file>")
        sys.exit(1)

    if not args.file.endswith('.bf'):
        print("Error: Input file must have a .bf extension", file=sys.stderr)
        sys.exit(1)

    try:
        with open(args.file, 'r') as file:
            code = file.read()
    except IOError as e:
        print(f"Error opening file: {e}", file=sys.stderr)
        sys.exit(1)
    
    interpret(code)

if __name__ == "__main__":
    main()

