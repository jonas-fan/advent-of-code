local operator = require "easy.operator"
local seq      = require "easy.seq"
local strings  = require "easy.strings"

local function part1(lines)
    local nums    = {}
    local symbols = {}

    for row, line in ipairs(lines) do
        for col, token in strings.enum(line, "%d+") do
            local num = {
                value = tonumber(token),
            }

            for i = 1, #token do
                nums[row]              = nums[row] or {}
                nums[row][col + i - 1] = num
            end
        end

        for col, token in strings.enum(line, "[^.%d]") do
            symbols[row]      = symbols[row] or {}
            symbols[row][col] = token
        end
    end

    local seen = {}

    for row in pairs(symbols) do
        for col in pairs(symbols[row]) do
            for i = -1, 1 do
                for j = -1, 1 do
                    local row = row + i
                    local col = col + j

                    if nums[row] and nums[row][col] then
                        if not seen[nums[row][col]] then
                            seen[#seen + 1] = nums[row][col].value
                        end

                        seen[nums[row][col]] = true
                    end
                end
            end
        end
    end

    return seq.reduce(seen, operator.add, 0)
end

local function part2(lines)
    local nums    = {}
    local symbols = {}

    for row, line in ipairs(lines) do
        for col, token in strings.enum(line, "%d+") do
            local num = {
                value = tonumber(token),
            }

            for i = 1, #token do
                nums[row]              = nums[row] or {}
                nums[row][col + i - 1] = num
            end
        end

        for col, token in strings.enum(line, "[*]") do
            symbols[row]      = symbols[row] or {}
            symbols[row][col] = token
        end
    end

    local out = 0

    for row in pairs(symbols) do
        for col in pairs(symbols[row]) do
            local seen = {}

            for i = -1, 1 do
                for j = -1, 1 do
                    local row = row + i
                    local col = col + j

                    if nums[row] and nums[row][col] then
                        if not seen[nums[row][col]] then
                            seen[#seen + 1] = nums[row][col].value
                        end

                        seen[nums[row][col]] = true
                    end
                end
            end

            if #seen > 1 then
                out = out + seq.reduce(seen, operator.mul, 1)
            end
        end
    end

    return out
end

local lines = {}

for line in io.lines() do
    lines[#lines + 1] = line
end

print("Part1", part1(lines))
print("Part2", part2(lines))
