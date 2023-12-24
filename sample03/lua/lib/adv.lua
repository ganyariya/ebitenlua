---メッセージを画面に表示する
---@param msg string
local function ShowMessage(msg)
    Call_Go_ShowMessage(msg)
    coroutine.yield()
end

return {
    ShowMessage = ShowMessage
}
