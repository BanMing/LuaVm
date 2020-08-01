--
-- Created by IntelliJ IDEA.
-- User: BanMing
-- Date: 2020/7/25
-- Time: 22:18
-- To change this template use File | Settings | File Templates.
--

--local t = {}
--
--local p = { x = 100, y = 200 }
--
--t[false] = nil
--assert(t[false] == nil)

local t = { "a", "b", "c" }
t[2] = "B"
t["foo"] = "Bar"
local s = t[3] .. t[2] .. t[1] .. t["foo"] .. #t