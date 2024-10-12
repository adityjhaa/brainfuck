#include <stdio.h>
#include <stdlib.h>

#define MEMORY_SIZE 30000

void interpret(char *code) {
    unsigned char memory[MEMORY_SIZE] = {0};
    unsigned char *data_pointer = memory;

    for (int i = 0; code[i] != '\0'; i++) {
        switch (code[i]) {
            case '>': 
                data_pointer++; 
                break;
            case '<': 
                data_pointer--; 
                break;
            case '+': 
                (*data_pointer)++; 
                break;
            case '-': 
                (*data_pointer)--; 
                break;
            case '.': 
                putchar(*data_pointer); 
                break;
            case ',': 
                *data_pointer = getchar(); 
                break;
            case '[':
                if (*data_pointer == 0) {
                    int loop = 1;
                    while (loop > 0) {
                        i++;
                        if (code[i] == '[') loop++;
                        else if (code[i] == ']') loop--;
                    }
                }
                break;
            case ']':
                if (*data_pointer != 0) {
                    int loop = 1;
                    while (loop > 0) {
                        i--;
                        if (code[i] == '[') loop--;
                        else if (code[i] == ']') loop++;
                    }
                }
                break;
        }
    }
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <brainfuck_code>\n", argv[0]);
        return 1;
    }

    FILE *file = fopen(argv[1], "r");
    if (!file) {
        perror("Error opening file");
        return 1;
    }

    fseek(file, 0, SEEK_END);
    long file_size = ftell(file);
    fseek(file, 0, SEEK_SET);
    char *code = (char *)malloc(file_size + 1);
    fread(code, 1, file_size, file);
    code[file_size] = '\0';
    fclose(file);

    interpret(code);

    free(code);

    return 0;
}
