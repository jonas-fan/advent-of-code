local enum = require "easy.enum"

local function copy(t)
    local out = {}

    for key, value in pairs(t) do
        out[key] = value
    end

    return out
end

local function part1(workflows, nums)
    local function eval(lhs, rhs, op)
        if op == "<" then
            return lhs < rhs
        end

        return lhs > rhs
    end

    local function dfs(workflows, num, target)
        if target == "R" then
            return 0
        end

        if target == "A" then
            local out = 0

            for _, value in pairs(num) do
                out = out + value
            end

            return out
        end

        local workflow = workflows[target]

        for _, rule in ipairs(workflow.rules) do
            if eval(num[rule.key], rule.value, rule.op) then
                return dfs(workflows, num, rule.target)
            end
        end

        return dfs(workflows, num, workflow.default)
    end

    return enum.reduce(nums, 0, function(num, acc)
        return acc + dfs(workflows, num, "in")
    end)
end

local function part2(workflows, num)
    local function dfs(workflows, num, target)
        if target == "R" then
            return 0
        end

        if target == "A" then
            local out = 1

            for _, value in pairs(num) do
                out = out * (value[2] - value[1] + 1)
            end

            return out
        end

        local out      = 0
        local workflow = workflows[target]

        for _, rule in ipairs(workflow.rules) do
            local include = {}
            local exclude = {}

            if rule.op == "<" then
                include[1], include[2] = num[rule.key][1], rule.value - 1
                exclude[1], exclude[2] = rule.value, num[rule.key][2]
            else
                include[1], include[2] = rule.value + 1, num[rule.key][2]
                exclude[1], exclude[2] = num[rule.key][1], rule.value
            end

            if include[1] <= include[2] then
                local copied = copy(num)

                copied[rule.key] = include

                out = out + dfs(workflows, copied, rule.target)
            end

            if exclude[1] <= exclude[2] then
                num           = copy(num)
                num[rule.key] = exclude
            end
        end

        out = out + dfs(workflows, num, workflow.default)

        return out
    end

    return dfs(workflows, num, "in")
end

local workflows = {}
local nums      = {}

for line in io.lines() do
    local values = string.match(line, "^{(.+)}$")

    if values then
        local num = {}

        for key, value in string.gmatch(values, "(%w)=(%d+)") do
            num[key] = tonumber(value)
        end

        nums[#nums + 1] = num
    elseif line ~= "" then
        local rules      = {}
        local default    = nil
        local name, rest = string.match(line, "^(%w+){(.+)}$")

        for key, op, value, target in string.gmatch(rest, "(%w+)([<>]?)(%d*):?([^,]*)") do
            if op == "" then
                default = key
            else
                rules[#rules + 1] = {
                    key    = key,
                    op     = op,
                    value  = tonumber(value),
                    target = target,
                }
            end
        end

        workflows[name] = {
            rules   = rules,
            default = default,
        }
    end
end

print("Part1", part1(workflows, nums))
print("Part2", part2(workflows, {
    x = { 1, 4000 },
    m = { 1, 4000 },
    a = { 1, 4000 },
    s = { 1, 4000 },
}))
