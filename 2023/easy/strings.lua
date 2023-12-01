local strings = {}

function strings.starts(str, prefix)
    return string.sub(str, 1, #prefix) == prefix
end

function strings.ends(str, suffix)
    return (suffix == "") or (string.sub(str, -#suffix) == suffix)
end

return strings
