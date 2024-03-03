-- 获取调用脚本时传递的第一个键名（验证码的键）
local key = KEYS[1]
-- 构造一个新的键名，用于存储尝试验证次数
local cntKey = key .. ":cnt"
-- 获取调用脚本时传递的第一个参数值（期望的验证码）
local expectedCode = ARGV[1]

-- 从Redis获取尝试验证次数，并转换为数字
local cnt = tonumber(redis.call("get", cntKey))
-- 从Redis获取存储的验证码
local code = redis.call("get", key)

-- 如果尝试次数已经用尽（小于等于0），返回-1表示无法再尝试
if cnt <= 0 then
    return -1
end

-- 如果存储的验证码与传入的期望验证码相同
if code == expectedCode then
    -- 设置尝试次数键的值为-1，表示验证成功
    redis.call("set", cntKey, -1)
    return 0 -- 返回0表示验证成功
else
    -- 如果验证码不匹配，减少一次尝试机会
    -- 注意：原代码中错误地使用了decr命令的语法，应该是redis.call("decr", cntKey)
    redis.call("decr", cntKey)
    return -2 -- 返回-2表示验证码错误
end
