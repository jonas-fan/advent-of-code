local enum    = require "easy.enum"
local strings = require "easy.strings"

local function hash(str)
    local sum = 0

    for each in string.gmatch(str, ".") do
        sum = ((sum + string.byte(each)) * 17) & (2 ^ 8 - 1)
    end

    return sum
end

local function part1(line)
    return enum.sum(enum.map(strings.split(line, ","), hash))
end

local function part2(line)
    local out   = {}
    local boxes = {}

    for _, str in ipairs(strings.split(line, ",")) do
        local index = string.find(str, "=", 1) or string.find(str, "-", 1)
        local label = string.sub(str, 1, index - 1)
        local value = tonumber(string.sub(str, index + 1))
        local where = tostring(hash(label))

        if string.sub(str, index, index) == "-" then
            local box = {}

            for _, lens in ipairs(boxes[where] or {}) do
                if lens.label ~= label then
                    box[#box + 1] = lens
                end
            end

            boxes[where] = (#box > 0) and box or nil
        else
            local found = false 

            for _, lens in ipairs(boxes[where] or {}) do
                if lens.label == label then
                    lens.value = value
                    found      = true
                end
            end

            if not found then
                boxes[where] = boxes[where] or {}

                table.insert(boxes[where], {
                    label = label,
                    value = value,
                })
            end
        end
    end

    local out = 0

    for where, box in pairs(boxes) do
        where = tonumber(where)

        for index, lens in ipairs(box) do
            out = out + (where + 1) * index * lens.value
        end
    end

    return out
end

local line = enum.slice(io.lines())[1]

print("Part1", part1(line))
print("Part1", part2(line))
