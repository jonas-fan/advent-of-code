local enum    = require "easy.enum"
local strings = require "easy.strings"

local dfs   = nil
local doDfs = nil

dfs = function(str, nums, count, cache)
    local key = string.format("%s|%s|%s", str, table.concat(nums, ","), count)

    if not cache[key] then
        cache[key] = doDfs(str, nums, count, cache)
    end

    return cache[key]
end

doDfs = function(str, nums, count, cache)
    if str == "" then
        if (#nums == 0 and count == 0) or
           (#nums == 1 and count == nums[1])
        then
            return 1
        end

        return 0
    end

    local first, rest = string.sub(str, 1, 1), string.sub(str, 2)

    if first == "#" then
        if #nums > 0 and count < nums[1] then
            return dfs(rest, nums, count + 1, cache)
        end
    elseif first == "." then
        if count == 0 then
            return dfs(rest, nums, count, cache)
        elseif #nums > 0 and count == nums[1] then
            return dfs(rest, enum.slice(nums, 2), 0, cache)
        end
    elseif first == "?" then
        if count == 0 then
            return dfs(rest, nums, count, cache) + dfs(rest, nums, count + 1, cache)
        elseif #nums > 0 and count == nums[1] then
            return dfs(rest, enum.slice(nums, 2), 0, cache)
        elseif #nums > 0 and count < nums[1] then
            return dfs(rest, nums, count + 1, cache)
        end
    end

    return 0
end

local function part1(lines)
    local strs = {}
    local nums = {}

    for _, line in ipairs(lines) do
        local matches = strings.split(line, " ")

        strs[#strs + 1] = matches[1]
        nums[#nums + 1] = enum.map(strings.split(matches[2], ","), tonumber)
    end

    local out = {}

    for i = 1, #strs do
        out[#out + 1] = dfs(strs[i], nums[i], 0, {})
    end

    return enum.sum(out)
end

local function part2(lines)
    local strs = {}
    local nums = {}

    for _, line in ipairs(lines) do
        local matches = strings.split(line, " ")

        lhs = { matches[1], matches[1], matches[1], matches[1], matches[1], }
        rhs = { matches[2], matches[2], matches[2], matches[2], matches[2], }

        strs[#strs + 1] = table.concat(lhs, "?")
        nums[#nums + 1] = enum.map(strings.split(table.concat(rhs, ","), ","), tonumber)
    end

    local out = {}

    for i = 1, #strs do
        out[#out + 1] = dfs(strs[i], nums[i], 0, {})
    end

    return enum.sum(out)
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
