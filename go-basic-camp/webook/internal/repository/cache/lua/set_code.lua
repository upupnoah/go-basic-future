-- phone_code:login:176xxxx
local key = KEYS[1]

-- phone_code:login:176xxxx:cnt
local cntKey = key .. ":cnt"

-- 外面生成的验证码
local val = ARGV[1]

-- 过期时间
local ttl = tonumber(redis.call("ttl", key))
if ttl == -1 then
    -- key 存在, 但是没有过期时间
    return -2
elseif ttl == -2 or ttl < 540 then
    -- key 不存在, 或者过期时间小于 540 秒
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", cntKey, 600)
    -- 符合预期, 验证码和过期时间都设置成功
    return 0
else
    -- 发送频繁
    return -1
end
