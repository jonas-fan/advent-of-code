local seq = {}

function seq.copy(iter)
    return table.move(iter, 1, #iter, 1, {})
end

function seq.first(iter)
    return iter[1]
end

function seq.last(iter)
    return iter[#iter]
end

function seq.filter(iter, fn)
    local out = {}

    for i = 1, #iter do
        local value = iter[i]

        if fn(value, i) then
            out[#out + 1] = value
        end
    end

    return out
end

function seq.map(iter, fn)
    local out = {}

    for i = 1, #iter do
        out[i] = fn(iter[i], i)
    end

    return out
end

function seq.reduce(iter, fn, initital)
    local out = initital

    for i = 1, #iter do
        out = fn(out, iter[i], i)
    end

    return out
end

function seq.sort(iter, cmp)
    table.sort(iter, cmp)

    return iter
end

function seq.sorted(iter, cmp)
    return seq.sort(seq.copy(iter), cmp)
end

function seq.reverse(iter)
    local size = #iter

    for i = 1, math.floor(size / 2) do
        iter[i], iter[size - i + 1] = iter[size - i + 1], iter[i]
    end

    return iter
end

function seq.reversed(iter)
    return seq.reverse(seq.copy(iter))
end

function seq.unique(iter)
    local out  = {}
    local seen = {}

    for i = 1, #iter do
        local value = iter[i]

        if not seen[value] then
            seen[value]   = true
            out[#out + 1] = value
        end
    end

    return out
end

function seq.uniqued(iter)
    return seq.unique(seq.copy(iter))
end

return seq
