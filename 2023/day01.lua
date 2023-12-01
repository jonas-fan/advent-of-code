local operator = require "easy.operator"
local seq      = require "easy.seq"
local strings  = require "easy.strings"

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

    local nums = seq.map(lines, function(line)
        local digits = dfs(line, {})

        return seq.first(digits) * 10 + seq.last(digits)
    end)

    return seq.reduce(nums, operator.add, 0)
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

    local nums = seq.map(lines, function(line)
        local digits = dfs(line, {})

        return seq.first(digits) * 10 + seq.last(digits)
    end)

    return seq.reduce(nums, operator.add, 0)
end

local lines = {}

for line in io.lines() do
    lines[#lines + 1] = line
end

print("Part1", part1(lines))
print("Part2", part2(lines))
