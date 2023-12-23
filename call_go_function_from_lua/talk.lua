local function setImage(i)
    call_go_set_image(i)
    coroutine.yield(100 + i)
end

function MainEntry()
    for i = 1, 100 do
        print("lua", "double", call_go_double(i))
        print("lua", "calcurate", call_go_calculate(i, 10))

        setImage(i)
        -- coroutine.yield(call_go_double(i), call_go_double(i))
    end
end
