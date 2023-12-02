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

return strings
