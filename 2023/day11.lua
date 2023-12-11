local enum = require "easy.enum"

local function findGalaxy(image, row, col)
    local out = {}

    for i in ipairs(image) do
        for j, val in ipairs(image[i]) do
            if val == "#" then
                out[#out + 1] = {
                    row      = i,
                    row_orig = i,
                    col      = j,
                    col_orig = j,
                }
            end
        end
    end

    return out
end

local function sol(image, galaxies, expandTimes)
    local horizon  = {}
    local vertical = {}

    for _, g in ipairs(galaxies) do
        horizon[g.row]  = true
        vertical[g.col] = true
    end

    for row = 1, #image do
        if not horizon[row] then
            for _, g in ipairs(galaxies) do
                if g.row_orig > row then
                    g.row = g.row + expandTimes - 1
                end
            end
        end
    end

    for col = 1, #image[1] do
        if not vertical[col] then
            for _, g in ipairs(galaxies) do
                if g.col_orig > col then
                    g.col = g.col + expandTimes - 1
                end
            end
        end
    end

    local out = {}

    for i = 1, #galaxies do
        for j = i + 1, #galaxies do
            out[#out + 1] = math.abs(galaxies[j].row - galaxies[i].row) +
                            math.abs(galaxies[j].col - galaxies[i].col)
        end
    end

    return enum.sum(out)
end

local image = {}

for line in io.lines() do
    image[#image + 1] = enum.slice(string.gmatch(line, "."))
end

print("Part1", sol(image, findGalaxy(image), 2))
print("Part2", sol(image, findGalaxy(image), 1000000))
