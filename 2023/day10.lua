local enum = require "easy.enum"

local function move(maze, visited, pos)
    local direct = ({
        ["-"] = { {  0, -1 }, { 0,  1 } },
        ["|"] = { { -1,  0 }, { 1,  0 } },
        ["F"] = { {  0,  1 }, { 1,  0 } },
        ["7"] = { {  0, -1 }, { 1,  0 } },
        ["L"] = { { -1,  0 }, { 0,  1 } },
        ["J"] = { { -1,  0 }, { 0, -1 } },
    })[maze[pos.row][pos.col]]

    local index = 0
    local iter  = nil

    iter = function()
        index = index + 1

        if not direct[index] then
            return nil
        end

        local row = pos.row + direct[index][1]
        local col = pos.col + direct[index][2]

        if row < 1 or row > #maze then
            return iter()
        elseif col < 1 or col > #maze[1] then
            return iter()
        elseif visited and visited[row] and visited[row][col] then
            return iter()
        end

        return { row = row, col = col }
    end

    return iter
end

local function sol(maze, pos)
    local step    = 0
    local nodes   = { pos }
    local visited = {}

    while #nodes > 0 do
        local queue = {}

        for _, pos in ipairs(nodes) do
            visited[pos.row]          = visited[pos.row] or {}
            visited[pos.row][pos.col] = true

            for pos in move(maze, visited, pos) do
                queue[#queue+1] = pos
            end
        end

        nodes = queue
        step  = step + (#nodes > 0 and 1 or 0)
    end

    local tiles = 0

    for i in ipairs(maze) do
        local line = {}

        for j in ipairs(maze[i]) do
            if visited[i] and visited[i][j] then
                line[#line + 1] = maze[i][j]
            else
                line[#line + 1] = "."
            end
        end

        line = table.concat(line, "")
        line = string.gsub(line, "F[-]*7", "||")
        line = string.gsub(line, "F[-]*J", "|")
        line = string.gsub(line, "L[-]*7", "|")
        line = string.gsub(line, "L[-]*J", "||")

        local counting = false

        for each in string.gmatch(line, ".") do
            if each == "|" then
                counting = not counting
            elseif each == "." and counting then
                tiles = tiles + 1
            end
        end
    end

    return step, tiles
end

local maze = {}
local pos  = nil

for i, line in ipairs(enum.slice(io.lines())) do
    maze[i] = enum.slice(string.gmatch(line, "."))

    for j, each in ipairs(maze[i]) do
        if each == "S" then
            pos = { row = i, col = j }
        end
    end
end

local minRow = 1
local maxRow = #maze
local minCol = 1
local maxCol = #maze[1]

if false then
    -- never go here
elseif string.find("|F7", maze[math.max(pos.row-1, minRow)][pos.col], 1, true) and
       string.find("-FL", maze[pos.row][math.max(pos.col-1, minCol)], 1, true) then
    maze[pos.row][pos.col] = "J"
elseif string.find("-FL", maze[pos.row][math.max(pos.col-1, minCol)], 1, true) and
       string.find("|LJ", maze[math.min(pos.row+1, maxRow)][pos.col], 1, true) then
    maze[pos.row][pos.col] = "7"
elseif string.find("|F7", maze[math.max(pos.row-1, minRow)][pos.col], 1, true) and
       string.find("-J7", maze[pos.row][math.min(pos.col+1, maxCol)], 1, true) then
    maze[pos.row][pos.col] = "L"
else
    maze[pos.row][pos.col] = "F"
end

local step, tiles = sol(maze, pos)

print("Part1", step)
print("Part2", tiles)
