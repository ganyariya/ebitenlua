---@param i number
---@param isRoutine boolean
local function displayImage(i, isRoutine)
    call_go_display_image(i)
    if isRoutine then
        coroutine.yield(100 + i)
    end
end

local function hello(msg)
    print("hello", msg)
end

return {
    displayImagge = displayImage,
    hello = hello
}
