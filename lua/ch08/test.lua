--
-- Created by IntelliJ IDEA.
-- User: BanMing
-- Date: 2020/8/4
-- Time: 08:18
-- To change this template use File | Settings | File Templates.
--

local function max(...)
    local args = { ... }
    local val, idx
    for i = 1, #args do
        if val == nil or args[i] > val then
            val, idx = args[i], i
        end
    end
    return val, idx
end

local function check(v)
    if not v then print("@@@@") end
end

local v1 = max(3, 9, 7, 128, 35)
--check(v1 == 128)

local v2, i2 = max(3, 9, 7, 128, 35)
--check(v2 == 128 and i2 == 4)

--local v3, i3 = max(max(3, 9, 7, 128, 35))
--check(v3 == 128 and i3 == 1)

--local t = { max(3, 9, 7, 128, 35) }
--check(t[1] == 128 and t[2] == 4)
