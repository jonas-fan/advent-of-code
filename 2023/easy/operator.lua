local operator = {}

function operator.eq(lhs, rhs)
    return lhs == rhs
end

function operator.neq(lhs, rhs)
    return lhs ~= rhs
end

function operator.add(lhs, rhs)
    return lhs + rhs
end

function operator.sub(lhs, rhs)
    return lhs - rhs
end

function operator.mul(lhs, rhs)
    return lhs * rhs
end

function operator.div(lhs, rhs)
    return lhs / div
end

return operator
