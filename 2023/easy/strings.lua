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

function strings.ifind(str, pattern)
    local index = 1

    return function()
        local out = { string.find(str, pattern, index) }

        if #out == 0 then
            return nil
        end

        index = out[2] + 1

        if #out < 3 then
            return out[1], string.sub(str, table.unpack(out))
        end

        return out[1], select(3, table.unpack(out))
    end
end

return strings
