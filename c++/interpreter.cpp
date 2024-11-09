#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include <algorithm>

constexpr int MAX_SIZE = 30000;

class Interpreter {
    std::vector<unsigned char> tape;
    std::vector<unsigned char>::iterator ptr;

public:
    Interpreter() : tape(MAX_SIZE, 0), ptr(tape.begin()) {};

    void interpret(const std::string& code) {
        for (auto it = code.begin(); it != code.end(); ++it) {
            switch (*it) {
                case '>': ++ptr; break;
                case '<': --ptr; break;
                case '+': ++*ptr; break;
                case '-': --*ptr; break;
                case '.': std::cout.put(*ptr); break;
                case ',': *ptr = std::cin.get(); break;
                case '[':
                    if (!*ptr) it = std::find(it, code.end(), ']');
                    break;
                case ']':
                    if (*ptr) it = std::find(code.rbegin() + (code.end() - it - 1), code.rend(), '[').base() - 1;
                    break;
                default : break;
            }
        }
    }
};

int main(int argc, char* argv[]) {
    if (argc != 2) return std::cerr << "Usage: " << argv[0] << " <file>\n", 1;
    std::string filename(argv[1]);
    if (filename.substr(filename.find_last_of(".") + 1) != "bf") return std::cerr << "Invalid file extension\n", 1;
    std::ifstream file(filename);
    if (!file) return std::cerr << "Error opening file\n", 1;\
    std::string code((std::istreambuf_iterator<char>(file)), {});
    file.close();

    Interpreter brainfuck;
    brainfuck.interpret(code);
    return 0;
}
