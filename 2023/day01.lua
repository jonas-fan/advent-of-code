local enum    = require "easy.enum"
local strings = require "easy.strings"

local function part1(lines)
    local function dfs(str, out)
        if #str == 0 then
            return out
        end

        local c = string.sub(str, 1, 1)

        if c ~= "0" and tonumber(c) then
            out[#out + 1] = tonumber(c)
        end

        return dfs(string.sub(str, 2), out)
    end

    local nums = enum.map(lines, function(line)
        local digits = dfs(line, {})

        return enum.at(digits, 1) * 10 + enum.at(digits, -1)
    end)

    return enum.sum(nums)
end

local function part2(lines)
    local words = {
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }

    local function dfs(str, out)
        if #str == 0 then
            return out
        end

        local c = string.sub(str, 1, 1)

        if c ~= "0" and tonumber(c) then
            out[#out + 1] = tonumber(c)
        else
            for index, word in ipairs(words) do
                if strings.starts(str, word) then
                    out[#out + 1] = index
                    break
                end
            end
        end

        return dfs(string.sub(str, 2), out)
    end

    local nums = enum.map(lines, function(line)
        local digits = dfs(line, {})

        return enum.at(digits, 1) * 10 + enum.at(digits, -1)
    end)

    return enum.sum(nums)
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
