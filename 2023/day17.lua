local enum = require "easy.enum"

local function move(node)
    local nodes   = {
        { row = node.row, col = node.col - 1, direct = 0 },
        { row = node.row, col = node.col + 1, direct = 1 },
        { row = node.row - 1, col = node.col, direct = 2 },
        { row = node.row + 1, col = node.col, direct = 3 },
    }

    local index = 0

    local function iter()
        index = index + 1

        return nodes[index]
    end

    return iter
end

local function sol(grid, minMoves, maxMoves)
    local height  = #grid
    local width   = #grid[1]
    local visited = {}
    local queue   = {
        { row = 1, col = 1, direct = 1, cont = 0, cost = 0 },
    }

    while #queue > 0 do
        local node = queue[#queue]

        table.remove(queue)

        if node.row == height and node.col == width and node.cont >= minMoves then
            return node.cost
        end

        local key = string.format("%s|%s|%s|%s",
            node.row, node.col, node.direct, node.cont)

        if not visited[key] then
            visited[key] = true

            for n in move(node) do
                if n.row < 1 or n.row > height then
                    -- out of boundary
                elseif n.col < 1 or n.col > width then
                    -- out of boundary
                elseif node.direct ~= n.direct and math.floor(node.direct / 2) == math.floor(n.direct / 2) then
                    -- reverse direction
                elseif node.direct ~= n.direct and node.cont > 0 and node.cont < minMoves then
                    -- not reach the min moves
                elseif node.direct == n.direct and node.cont >= maxMoves then
                    -- reached the max moves
                else
                    n.cont = (node.direct == n.direct and node.cont or 0) + 1
                    n.cost = node.cost + grid[n.row][n.col]

                    local index = 1

                    if #queue > 0 then
                        for i = #queue, 1, -1 do
                            if n.cost < queue[i].cost then
                                index = i + 1
                                break
                            end
                        end
                    end

                    table.insert(queue, index, n)
                end
            end
        end
    end

    return nil
end

local grid = {}

for line in io.lines() do
    grid[#grid + 1] = enum.map(string.gmatch(line, "."), tonumber)
end

print("Part1", sol(grid, 0, 3))
print("Part2", sol(grid, 4, 10))
