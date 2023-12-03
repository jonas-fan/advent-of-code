local strings = {}

function strings.starts(str, prefix)
    return string.sub(str, 1, #prefix) == prefix
end

function strings.ends(str, suffix)
    return (suffix == "") or (string.sub(str, -#suffix) == suffix)
end

function strings.split(str, sep)
    local out = {}

    for token in string.gmatch(str, "[^" .. sep .. "]+") do
        out[#out + 1] = token
    end

    return out
end

function strings.enum(str, pattern)
    pattern = pattern or "."

    local index = 0
    local token

    return function()
        local from, to = string.find(str, pattern)

        if from == nil then
            return nil
        end

        token = string.sub(str, from, to)
        str   = string.sub(str, to + 1)
        index = index + to

        return index - to + from, token
    end
end

return strings
