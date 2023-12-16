local enum = require "easy.enum"

local delta = {
    ["^"] = { -1,  0 },
    ["v"] = {  1,  0 },
    ["<"] = {  0, -1 },
    [">"] = {  0,  1 },
}

local function move(grid, history, beam, direct)
    if beam.row < 1 or beam.row > #grid then
        return {}
    elseif beam.col < 1 or beam.col > #grid[beam.row] then
        return {}
    end

    local out = {
        row     = beam.row,
        col     = beam.col,
        direct  = direct or beam.direct,
    }

    local key = out.row .. "," .. out.col

    history[key] = history[key] or {}

    if history[key][out.direct] then
        return {}
    end

    history[key][out.direct] = true

    if grid[out.row][out.col] == [[/]] then
        if out.direct == "^" then
            out.direct = ">"
        elseif out.direct == "v" then
            out.direct = "<"
        elseif out.direct == "<" then
            out.direct = "v"
        else
            out.direct = "^"
        end
    elseif grid[out.row][out.col] == [[\]] then
        if out.direct == "^" then
            out.direct = "<"
        elseif out.direct == "v" then
            out.direct = ">"
        elseif out.direct == "<" then
            out.direct = "^"
        else
            out.direct = "v"
        end
    elseif grid[out.row][out.col] == [[-]] then
        if not string.find("<>", out.direct, 1, true) then
            local left  = move(grid, history, out, "<")
            local right = move(grid, history, out, ">")

            return table.move(left, 1, #left, #right + 1, right)
        end
    elseif grid[out.row][out.col] == [[|]] then
        if not string.find("^v", out.direct, 1, true) then
            local up   = move(grid, history, out, "^")
            local down = move(grid, history, out, "v")

            return table.move(up, 1, #up, #down + 1, down)
        end
    end

    out.row = out.row + delta[out.direct][1]
    out.col = out.col + delta[out.direct][2]

    return { out }
end

local function sol(grid, row, col, direct)
    local beam    = { row = row, col = col, direct = direct }
    local nodes   = { beam }
    local history = {}

    while #nodes > 0 do
        local queue = {}

        for _, beam in ipairs(nodes) do
            for _, next in pairs(move(grid, history, beam)) do
                queue[#queue + 1] = next
            end
        end

        nodes = queue
    end

    local out = 0

    for _ in pairs(history) do
        out = out + 1
    end

    return out
end

local function part1(grid)
    return sol(grid, 1, 1, ">")
end

local function part2(grid)
    local out = 0

    for row = 1, #grid do
        out = math.max(out, sol(grid, row, 1, ">"))
        out = math.max(out, sol(grid, row, #grid[row], "<"))
    end

    for col = 1, #grid[1] do
        out = math.max(out, sol(grid, 1, col, "v"))
        out = math.max(out, sol(grid, #grid, col, "^"))
    end

    return out
end

local grid = {}

for line in io.lines() do
    grid[#grid + 1] = enum.slice(string.gmatch(line, "."))
end

print("Part1", part1(grid))
print("Part2", part2(grid))
