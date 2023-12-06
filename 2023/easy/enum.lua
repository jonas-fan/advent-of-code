local enum = {}

local function add(lhs, rhs)
    return lhs + rhs
end

local function mul(lhs, rhs)
    return lhs * rhs
end

local function toslice(enumerable)
    local out = enumerable

    if type(out) == "function" then
        out = {}

        for value in enumerable do
            out[#out + 1] = value
        end
    end

    return out
end

local function each(enumerable)
    if type(enumerable) == "function" then
        return enumerable
    end

    local index = 0
    local size  = #enumerable

    return function()
        if index >= size then
            return nil
        end

        index = index + 1

        return enumerable[index]
    end
end

function enum.at(enumerable, index)
    enumerable = toslice(enumerable)

    if math.abs(index) > #enumerable then
        return nil
    end

    if index < 0 then
        index = index + #enumerable +  1
    end

    return enumerable[index]
end

function enum.into(enumerable, out)
    for value in each(enumerable) do
        out[#out + 1] = value
    end

    return out
end

function enum.map(enumerable, fn)
    local out = {}

    for value in each(enumerable) do
        out[#out + 1] = fn(value)
    end

    return out
end

function enum.max(enumerable, fn)
    return enum.reduce(enumerable, fn or math.max)
end

function enum.min(enumerable, fn)
    return enum.reduce(enumerable, fn or math.min)
end

function enum.product(enumerable)
    return enum.reduce(enumerable, mul)
end

function enum.reduce(enumerable, ...)
    local args = table.pack(...)
    local fn   = args[1]
    local out  = args[1]

    if type(fn) ~= "function" then
        fn = args[2]
    end

    for value in each(enumerable) do
        if out == fn then
            out = value
        else
            out = fn(value, out)
        end
    end

    return out
end

function enum.slice(enumerable, start, stop)
    enumerable = toslice(enumerable)
    start      = start or 1
    stop       = stop or #enumerable

    if start < 0 then
        start = start + #enumerable + 1
    end

    if stop < 0 then
        stop = stop + #enumerable + 1
    end

    return table.move(enumerable, start, stop, 1, {})
end

function enum.sum(enumerable)
    return enum.reduce(enumerable, add)
end

return enum
