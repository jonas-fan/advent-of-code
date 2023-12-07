local enum = require "easy.enum"

local kind = {
    nothing   = -1,
    high      = 0,
    onePair   = 1,
    twoPair   = 2,
    three     = 3,
    fullHouse = 3.5,
    four      = 4,
    five      = 5,
}

local function comparer(order)
    return function(lhs, rhs)
        if lhs.kind ~= rhs.kind then
            return lhs.kind < rhs.kind
        end

        for i = 1, 5 do
            local left  = string.sub(lhs.token, i, i)
            local right = string.sub(rhs.token, i, i)

            if left ~= right then
                return string.find(order, left) < string.find(order, right)
            end
        end

        return false
    end
end

local function part1(lines)
    local hands = {}

    for _, line in ipairs(lines) do
        local matches = { string.match(line, "^(.+)%s(.+)$") }
        local hand    = {
            token = matches[1],
            bid   = tonumber(matches[2]),
            kind  = kind.nothing
        }

        local card = {}
        local diff = 0
        local most = 0

        for each in string.gmatch(hand.token, ".") do
            card[each] = (card[each] or 0) + 1

            if card[each] == 1 then
                diff = diff + 1
            end

            most = math.max(most, card[each])
        end

        if diff == 1 then
            hand.kind = kind.five
        elseif diff == 2 then
            hand.kind = most == 4 and kind.four or kind.fullHouse
        elseif diff == 3 then
            hand.kind = most == 3 and kind.three or kind.twoPair
        elseif diff == 4 then
            hand.kind = kind.onePair
        elseif string.find("23456789TJQKA", hand.token) then
            hand.kind = kind.highCard
        end

        hands[#hands+1] = hand
    end

    table.sort(hands, comparer("23456789TJQKA"))

    local out = {}

    for i, hand in ipairs(hands) do
        out[#out + 1] = i * hand.bid
    end

    return enum.sum(out)
end

local function part2(lines)
    local hands = {}

    for _, line in ipairs(lines) do
        local matches = { string.match(line, "^(.+)%s(.+)$") }
        local hand    = {
            token = matches[1],
            bid   = tonumber(matches[2]),
            kind  = kind.nothing
        }

        local card = {}
        local diff = 0
        local most = 0

        for each in string.gmatch(hand.token, ".") do
            card[each] = (card[each] or 0) + 1

            if card[each] == 1 then
                diff = diff + 1
            end

            if each ~= "J" then
                most = math.max(most, card[each])
            end
        end

        if card.J then
            diff = math.max(diff - 1, 1)
            most = most + card.J
        end

        if diff == 1 then
            hand.kind = kind.five
        elseif diff == 2 then
            hand.kind = most == 4 and kind.four or kind.fullHouse
        elseif diff == 3 then
            hand.kind = most == 3 and kind.three or kind.twoPair
        elseif diff == 4 then
            hand.kind = kind.onePair
        elseif string.find("J23456789TQKA", hand.token) then
            hand.kind = kind.highCard
        end

        hands[#hands+1] = hand
    end

    table.sort(hands, comparer("J23456789TQKA"))

    local out = {}

    for i, hand in ipairs(hands) do
        out[#out + 1] = i * hand.bid
    end

    return enum.sum(out)
end

local lines = enum.slice(io.lines())

print("Part1", part1(lines))
print("Part2", part2(lines))
