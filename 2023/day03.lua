local enum    = require "easy.enum"
local strings = require "easy.strings"

local function part1(lines)
    local nums    = {}
    local symbols = {}

    for row, line in ipairs(lines) do
        for col, token in strings.ifind(line, "%d+") do
            local num = {
                value = tonumber(token),
            }

            for i = 1, #token do
                nums[row]              = nums[row] or {}
                nums[row][col + i - 1] = num
            end
        end

        for col, token in strings.ifind(line, "[^.%d]") do
            symbols[row]      = symbols[row] or {}
            symbols[row][col] = token
        end
    end

    local seen = {}

    for row in pairs(symbols) do
        for col in pairs(symbols[row]) do
            for row = row - 1, row + 1 do
                for col = col - 1, col + 1 do
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

    return enum.sum(seen)
end

local function part2(lines)
    local nums    = {}
    local symbols = {}

    for row, line in ipairs(lines) do
        for col, token in strings.ifind(line, "%d+") do
            local num = {
                value = tonumber(token),
            }

            for i = 1, #token do
                nums[row]              = nums[row] or {}
                nums[row][col + i - 1] = num
            end
        end

        for col, token in strings.ifind(line, "[*]") do
            symbols[row]      = symbols[row] or {}
            symbols[row][col] = token
        end
    end

    local out = {}

    for row in pairs(symbols) do
        for col in pairs(symbols[row]) do
            local seen = {}

            for row = row - 1, row + 1 do
                for col = col - 1, col + 1 do
                    if nums[row] and nums[row][col] then
                        if not seen[nums[row][col]] then
                            seen[#seen + 1] = nums[row][col].value
                        end

                        seen[nums[row][col]] = true
                    end
                end
            end

            if #seen > 1 then
                out[#out + 1] = enum.product(seen)
            end
        end
    end

    return enum.sum(out)
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
