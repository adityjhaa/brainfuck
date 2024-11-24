local MEMORY_SIZE = 30000

local function interpret(code)
    local memory = {}
    for i = 1, MEMORY_SIZE do memory[i] = 0 end
    local ptr = 1

    local function find_matching_bracket(start, direction)
        local depth = 0
        local index = start
        while true do
            if code:sub(index, index) == "[" then
                depth = depth + direction
            elseif code:sub(index, index) == "]" then
                depth = depth - direction
            end
            if depth == 0 then return index end
            index = index + direction
        end
    end

    local i = 1
    while i <= #code do
        local char = code:sub(i, i)
        if char == ">" then
            ptr = (ptr % MEMORY_SIZE) + 1
        elseif char == "<" then
            ptr = (ptr - 2) % MEMORY_SIZE + 1
        elseif char == "+" then
            memory[ptr] = (memory[ptr] + 1) % 256
        elseif char == "-" then
            memory[ptr] = (memory[ptr] - 1) % 256
        elseif char == "." then
            io.write(string.char(memory[ptr]))
        elseif char == "," then
            memory[ptr] = string.byte(io.read(1) or "\0")
        elseif char == "[" then
            if memory[ptr] == 0 then
                i = find_matching_bracket(i, 1)
            end
        elseif char == "]" then
            if memory[ptr] ~= 0 then
                i = find_matching_bracket(i, -1)
            end
        end
        i = i + 1
    end
end

local function main()
    if #arg == 0 or not arg[1]:match("%.bf$") then
        print("Usage: lua interpreter.lua <file.bf>")
        return
    end

    local file_name = arg[1]
    local file, err = io.open(file_name, "r")
    if not file then
        print("Error opening file:", err)
        return
    end

    local code = file:read("*a")
    file:close()
    interpret(code)
end

main()

