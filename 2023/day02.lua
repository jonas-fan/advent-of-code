local enum    = require "easy.enum"
local strings = require "easy.strings"

local function part1(lines)
    local out = {}

    for _, line in ipairs(lines) do
        local id, rest = string.match(line, "^Game (%d+): (.*)$")

        local cube = {
            blue  = 0,
            green = 0,
            red   = 0,
        }

        for _, set in ipairs(strings.split(rest, ";")) do
            for num, color in string.gmatch(set, "(%d+)%s(%w+)") do
                cube[color] = math.max(cube[color], tonumber(num))
            end
        end

        if cube.red <= 12 and cube.green <= 13 and cube.blue <=14 then
            out[#out + 1] = id
        end
    end

    return enum.sum(out)
end

local function part2(lines)
    local out = {}

    for _, line in ipairs(lines) do
        local id, rest = string.match(line, "^Game (%d+): (.*)$")

        local cube = {
            blue  = 0,
            green = 0,
            red   = 0,
        }

        for _, set in ipairs(strings.split(rest, ";")) do
            for num, color in string.gmatch(set, "(%d+)%s(%w+)") do
                cube[color] = math.max(cube[color], tonumber(num))
            end
        end

        out[#out + 1] = cube.red * cube.green * cube.blue
    end

    return enum.sum(out)
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
