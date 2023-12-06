local enum = require "easy.enum"

local races = {
    {},
    {},
}

for line in io.lines() do
    local key = string.lower(string.sub(line, 1, 1))
    local i   = 1

    for token in string.gmatch(line, "%d+") do
        races[1][i]      = races[1][i] or {}
        races[1][i][key] = tonumber(token)
        races[2][1]      = races[2][1] or {}
        races[2][1][key] = (races[2][1][key] or "") .. token

        i = i + 1
    end

    races[2][1][key] = tonumber(races[2][1][key])
end

local function sol(races)
    local out = {}

    for _, r in ipairs(races) do
        local first = 1
        local last  = r.t - 1

        while first * (r.t - first) <= r.d do
            first = first + 1
        end

        while last * (r.t - last) <= r.d do
            last = last - 1
        end

        out[#out + 1] = last - first + 1
    end

    return enum.product(out)
end

print("Part1", sol(races[1]))
print("Part2", sol(races[2]))
