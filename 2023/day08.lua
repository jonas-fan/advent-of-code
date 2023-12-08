local enum    = require "easy.enum"
local strings = require "easy.strings"

local instructions = ""
local route        = {}

for i, line in ipairs(enum.slice(io.lines())) do
    if i == 1 then
        instructions = line
    elseif i > 2 then
        local matches = { string.match(line, "(%w+) = %((%w+), (%w+)%)") }

        route[matches[1]] = {
            L = matches[2],
            R = matches[3],
        }
    end
end

local function part1(route)
    local node  = "AAA"
    local steps = 0

    while true do
        for next in string.gmatch(instructions, ".") do
            steps = steps + 1
            node  = route[node][next]

            if node == "ZZZ" then
                return steps
            end
        end
    end
end

local function part2(route)
    local nodes = {}
    local steps = {}

    for node in pairs(route) do
        if strings.ends(node, "A") then
            nodes[#nodes + 1] = node
            steps[#steps + 1] = 0
        end
    end

    local out = {}

    for i, node in ipairs(nodes) do
        local done = false

        repeat
            for next in string.gmatch(instructions, ".") do
                steps[i] = steps[i] + 1
                node     = route[node][next]

                if strings.ends(node, "Z") then
                    done = true
                    break
                end
            end
        until done
    end

    local function gcd(lhs, rhs)
        if rhs == 0 then
            return lhs
        end

        return gcd(rhs, lhs % rhs)
    end

    local function lcm(lhs, rhs)
        return math.tointeger(lhs * rhs / gcd(lhs, rhs))
    end

    return enum.reduce(steps, lcm)
end

print("Part1", part1(route))
print("Part2", part2(route))
