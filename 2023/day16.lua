local enum = require "easy.enum"

local delta = {
    ["^"] = { -1,  0 },
    ["v"] = {  1,  0 },
    ["<"] = {  0, -1 },
    [">"] = {  0,  1 },
}

local function move(grid, beam, direct)
    if beam.row < 1 or beam.row > #grid then
        return {}
    elseif beam.col < 1 or beam.col > #grid[beam.row] then
        return {}
    end

    local out = {
        row     = beam.row,
        col     = beam.col,
        direct  = direct or beam.direct,
        history = beam.history
    }

    local key = out.row .. "," .. out.col

    out.history[key] = out.history[key] or {}

    if out.history[key][out.direct] then
        return {}
    end

    out.history[key][out.direct] = true

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
            local left  = move(grid, out, "<")
            local right = move(grid, out, ">")

            return table.move(left, 1, #left, #right + 1, right)
        end
    elseif grid[out.row][out.col] == [[|]] then
        if not string.find("^v", out.direct, 1, true) then
            local up   = move(grid, out, "^")
            local down = move(grid, out, "v")

            return table.move(up, 1, #up, #down + 1, down)
        end
    end

    out.row = out.row + delta[out.direct][1]
    out.col = out.col + delta[out.direct][2]

    return { out }
end

local function sol(grid, beam)
    local nodes = { beam }

    while #nodes > 0 do
        local queue = {}

        for _, beam in ipairs(nodes) do
            for _, next in pairs(move(grid, beam)) do
                queue[#queue + 1] = next
            end
        end

        nodes = queue
    end

    local out = 0

    for _ in pairs(beam.history) do
        out = out + 1
    end

    return out
end

local function part1(grid)
    local beam = {
        row     = 1,
        col     = 1,
        direct  = ">",
        history = {},
    }

    return sol(grid, beam)
end

local function part2(grid)
    local out = 0

    for row = 1, #grid do
        for col = 1, #grid[row] do
            if row == 1 or row == #grid or col == 1 or col == #grid[row] then
                for _, direct in ipairs({ "^", "v", "<", ">" }) do
                    local beam = {
                        row     = row,
                        col     = col,
                        direct  = direct,
                        history = {},
                    }

                    out = math.max(out, sol(grid, beam))
                end
            end
        end
    end

    return out
end

local grid = {}

for line in io.lines() do
    grid[#grid + 1] = enum.slice(string.gmatch(line, "."))
end

print("Part1", part1(grid))
print("Part2", part2(grid))
