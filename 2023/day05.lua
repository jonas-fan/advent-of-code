local enum    = require "easy.enum"
local strings = require "easy.strings"

local function overlapped(x, y, w, z)
    if math.max(x, y) < math.min(w, z) then
        return false
    elseif math.min(x, y) > math.max(w, z) then
        return false
    end

    return true
end

local function part1(lines)
    local seeds = {}
    local maps  = {}
    local rules = nil

    for _, line in ipairs(lines) do
        if line == "" then
            -- nothing
        elseif strings.starts(line, "seeds:") then
            for each in string.gmatch(strings.split(line, ":")[2], "%d+") do
                seeds[#seeds + 1] = tonumber(each)
            end
        elseif string.find(line, "%w+%-to%-%w+ map:") then
            rules           = {}
            maps[#maps + 1] = rules
        else
            local matches = { string.match(line, "(%d+) (%d+) (%d+)") }

            rules[#rules + 1] = {
                base = tonumber(matches[1]),
                min  = tonumber(matches[2]),
                max  = tonumber(matches[2]) + tonumber(matches[3]),
            }
        end
    end

    local function dfs(maps, seeds)
        if #maps == 0 then
            return seeds
        end

        for i, seed in ipairs(seeds) do
            for _, rule in ipairs(maps[1]) do
                if overlapped(seed, seed, rule.min, rule.max) then
                    seed = seed - rule.min + rule.base
                    break
                end
            end

            seeds[i] = seed
        end

        return dfs(enum.slice(maps, 2), seeds)
    end

    return enum.min(dfs(maps, seeds))
end

local function part2(lines)
    local seeds = {}
    local maps  = {}
    local rules = nil

    for _, line in ipairs(lines) do
        if line == "" then
            -- nothing
        elseif strings.starts(line, "seeds:") then
            for lhs, rhs in string.gmatch(strings.split(line, ":")[2], "(%d+) (%d+)") do
                seeds[#seeds + 1] = {
                    min = tonumber(lhs),
                    max = tonumber(lhs) + tonumber(rhs),
                }
            end
        elseif string.find(line, "%w+%-to%-%w+ map:") then
            rules           = {}
            maps[#maps + 1] = rules
        else
            local matches = { string.match(line, "(%d+) (%d+) (%d+)") }

            rules[#rules + 1] = {
                base = tonumber(matches[1]),
                min  = tonumber(matches[2]),
                max  = tonumber(matches[2]) + tonumber(matches[3]),
            }
        end
    end

    local function dfs(maps, seeds)
        if #maps == 0 then
            return seeds
        end

        local out = {}

        while #seeds > 0 do
            local seed = seeds[1]

            seeds = enum.slice(seeds, 2)

            for _, rule in ipairs(maps[1]) do
                if overlapped(seed.min, seed.max, rule.min, rule.max) then
                    if seed.min < rule.min and seed.max > rule.max then
                        seeds[#seeds + 1] = { min = seed.min,     max = rule.min - 1 }
                        seeds[#seeds + 1] = { min = rule.min,     max = rule.max }
                        seeds[#seeds + 1] = { min = rule.max + 1, max = seed.max }
                        seed              = nil
                    elseif seed.min < rule.min then
                        seeds[#seeds + 1] = { min = seed.min, max = rule.min - 1 }
                        seeds[#seeds + 1] = { min = rule.min, max = seed.max }
                        seed              = nil
                    elseif seed.max > rule.max then
                        seeds[#seeds + 1] = { min = seed.min,     max = rule.max }
                        seeds[#seeds + 1] = { min = rule.max + 1, max = seed.max }
                        seed              = nil
                    else
                        seed.min = seed.min - rule.min + rule.base
                        seed.max = seed.max - rule.min + rule.base
                    end

                    break
                end
            end

            out[#out + 1] = seed
        end

        return dfs(enum.slice(maps, 2), out)
    end

    local out = enum.min(dfs(maps, seeds), function(lhs, rhs)
        return lhs.min < rhs.min and lhs or rhs
    end)

    return out.min
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
