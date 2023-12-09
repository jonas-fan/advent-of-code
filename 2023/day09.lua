local enum = require "easy.enum"

local function is(want)
    return function(value)
        return value == want
    end
end

local function differences(nums)
    local out = {}

    for i = 1, #nums - 1 do
        out[i] = nums[i + 1] - nums[i]
    end

    return out
end

local function part1(lines)
    local function dfs(nums)
        if enum.all(nums, is(0)) then
            return 0
        end

        return nums[#nums] + dfs(differences(nums))
    end

    return enum.sum(enum.map(lines, dfs))
end

local function part2(lines)
    local function dfs(nums)
        if enum.all(nums, is(0)) then
            return 0
        end

        return nums[1] - dfs(differences(nums))
    end

    return enum.sum(enum.map(lines, dfs))
end

local lines = {}

for line in io.lines() do
    lines[#lines + 1] = enum.map(string.gmatch(line, "[-]?%d+"), tonumber)
end

print("Part1", part1(lines))
print("Part2", part2(lines))
