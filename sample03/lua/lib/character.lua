local function ShowCharacter(imagePath)
    Call_Go_DrawImage(imagePath)
    coroutine.yield()
end

return {
    ShowCharacter = ShowCharacter,
}
