-- 使用 wrk 发送 POST 请求进行用户注册测试
-- 设置请求方法和头部
wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"

-- 生成随机 UUID
local function uuid()
    local template = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'
    return string.gsub(template, '[xy]', function(c)
        local v = (c == 'x') and math.random(0, 0xf) or math.random(8, 0xb)
        return string.format('%x', v)
    end)
end

local cnt, prefix

-- 初始化函数，每个线程执行
function init(args)
    cnt = 0
    prefix = uuid() -- 为每个线程生成唯一前缀
end

-- 构造请求体，每次请求递增 cnt 以生成唯一邮箱
function request()
    local body = string.format(
        '{"email":"%s%d@gmail.com", "password":"TEST12345678!!", "confirmPassword": "TEST12345678!!"}', prefix, cnt)
    cnt = cnt + 1
    return wrk.format(nil, wrk.path, wrk.headers, body)
end
