local operator = require "easy.operator"
local seq      = require "easy.seq"
local strings  = require "easy.strings"

local function part1(lines)
    local out = {}

    for _, line in ipairs(lines) do
        local round, lhs, rhs = string.match(line, "^Card%s+(%d+): (.*) | (.*)$")
        local wins            = {}
        local ours            = {}

        for each in string.gmatch(lhs, "%d+") do
            wins[each] = true
        end

        for each in string.gmatch(rhs, "%d+") do
            ours[each] = true
        end

        local points = 0

        for key in pairs(wins) do
            if ours[key] then
                if points == 0 then
                    points = 1
                else
                    points = points + points
                end
            end
        end

        out[#out + 1] = points
    end

    return seq.reduce(out, operator.add, 0)
end

local function part2(lines)
    local out = {}

    for _, line in ipairs(lines) do
        local round, lhs, rhs = string.match(line, "^Card%s+(%d+): (.*) | (.*)$")
        local wins            = {}
        local ours            = {}

        round = tonumber(round)

        for each in string.gmatch(lhs, "%d+") do
            wins[each] = true
        end

        for each in string.gmatch(rhs, "%d+") do
            ours[each] = true
        end

        local points = 0

        for key in pairs(wins) do
            if ours[key] then
                points = points + 1
            end
        end

        out[round] = out[round] or 1

        for i = 1, points do
            out[round + i] = out[round + i] or 1
            out[round + i] = out[round + i] + out[round]
        end
    end

    return seq.reduce(out, operator.add, 0)
end

local lines = {}

for line in io.lines() do
    lines[#lines + 1] = line
end

print("Part1", part1(lines))
print("Part2", part2(lines))
