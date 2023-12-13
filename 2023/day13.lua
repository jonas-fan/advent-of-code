local enum = require "easy.enum"

local function sameRow(matrix, a, b)
    if a < 1 or a > #matrix then
        return false
    elseif b < 1 or b > #matrix then
        return false
    end

    for col = 1, #matrix[a] do
        if matrix[a][col] ~= matrix[b][col] then
            return false
        end
    end

    return true
end

local function sameCol(matrix, a, b)
    if a < 1 or a > #matrix[1] then
        return false
    elseif b < 1 or b > #matrix[1] then
        return false
    end

    for row = 1, #matrix do
        if matrix[row][a] ~= matrix[row][b] then
            return false
        end
    end

    return true
end

local function findVertical(matrix, left, right, except)
    if left > right then
        return 0
    end

    local mid  = math.floor((left + right) / 2)
    local l    = mid
    local r    = mid + 1
    local same = false

    while l >= 1 and r <= #matrix[1] do
        same = sameCol(matrix, l, r)

        if not same then
            break
        end

        l = l - 1
        r = r + 1
    end

    if (not same) or (mid == except) then
        return math.max(findVertical(matrix, left, mid-1, except),
                        findVertical(matrix, mid+1, right, except))
    end

    return mid
end

local function findHorizontal(matrix, left, right, except)
    if left > right then
        return 0
    end

    local mid  = math.floor((left + right) / 2)
    local l    = mid
    local r    = mid + 1
    local same = false

    while l >= 1 and r <= #matrix do
        same = sameRow(matrix, l, r)

        if not same then
            break
        end

        l = l - 1
        r = r + 1
    end

    if (not same) or (mid == except) then
        return math.max(findHorizontal(matrix, left, mid-1, except),
                        findHorizontal(matrix, mid+1, right, except))
    end

    return mid
end

local function part1(matrixes)
    local out = {}

    for i, matrix in ipairs(matrixes) do
        local cols = findVertical(matrix, 1, #matrix[1])
        local rows = findHorizontal(matrix, 1, #matrix)

        out[#out + 1] = rows * 100 + cols
    end

    return enum.sum(out)
end

local function part2(matrixes)
    local out = {}

    for i, matrix in ipairs(matrixes) do
        local cols = findVertical(matrix, 1, #matrix[1])
        local rows = findHorizontal(matrix, 1, #matrix)

        for row = 1, #matrix do
            local ok = false

            for col = 1, #matrix[row] do
                local old = matrix[row][col]

                matrix[row][col] = old == "." and "#" or "."
                local c = findVertical(matrix, 1, #matrix[1], cols)
                local r = findHorizontal(matrix, 1, #matrix, rows)
                matrix[row][col] = old

                if c > 0 or r > 0 then
                    ok = true
                    cols = c
                    rows = r
                    break
                end
            end

            if ok then
                break
            end
        end

        out[#out + 1] = rows * 100 + cols
    end

    return enum.sum(out)
end

local matrixes = {}
local matrix   = nil

for line in io.lines() do
    if line == "" then
        matrixes[#matrixes + 1] = matrix
        matrix                  = nil
    else
        matrix              = matrix or {}
        matrix[#matrix + 1] = enum.slice(string.gmatch(line, "."))
    end
end

if matrix then
    matrixes[#matrixes + 1] = matrix
end

print("Part1", part1(matrixes))
print("Part2", part2(matrixes))
