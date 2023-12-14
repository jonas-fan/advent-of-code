local enum = require "easy.enum"

local function tilt(matrix, row, col, deltaRow, deltaCol)
    local nextRow = row + deltaRow
    local nextCol = col + deltaCol

    if nextRow < 1 or nextRow > #matrix then
        return
    elseif nextCol < 1 or nextCol > #matrix[row] then
        return
    elseif matrix[row][col] ~= "O" then
        return
    elseif matrix[nextRow][nextCol] ~= "." then
        return
    end

    matrix[row][col]         = "."
    matrix[nextRow][nextCol] = "O"

    return tilt(matrix, nextRow, nextCol, deltaRow, deltaCol)
end

local function sol(matrix, deltas, cycles)
    local out   = 0
    local step  = 1
    local cache = {}

    while step <= cycles do
        for _, delta in ipairs(deltas) do
            for r = (delta[1] < 0 and 1 or #matrix), (delta[1] < 0 and #matrix or 1), (delta[1] < 0 and 1 or -1) do
                for c = (delta[2] < 0 and 1 or #matrix[r]), (delta[2] < 0 and #matrix[r] or 1), (delta[2] < 0 and 1 or -1) do
                    tilt(matrix, r, c, delta[1], delta[2])
                end
            end
        end

        if step >= cycles then
            break
        end

        local key = table.concat(enum.map(matrix, function(cols)
                return table.concat(cols, "")
            end), "")

        if cache[key] then
            local diff = step - cache[key]

            while step <= cycles do
                step = step + diff
            end

            if step > cycles then
                step = step - diff
            end
        end

        cache[key] = step
        step       = step + 1
    end

    for r = 1, #matrix do
        for c = 1, #matrix[r] do
            if matrix[r][c] == "O" then
                out = out + #matrix - r + 1
            end
        end
    end

    return out
end

local matrix = {}

for line in io.lines() do
    matrix[#matrix + 1] = enum.slice(string.gmatch(line, "."))
end

print("Part1", sol(matrix, { { -1, 0 } }, 1))
print("Part2", sol(matrix, { { -1, 0 }, { 0, -1 }, { 1, 0 }, { 0, 1 } }, 1000000000))
