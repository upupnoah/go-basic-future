-- 1. 定义限流键、窗口大小、阈值和当前时间
local key = KEYS[1]                 -- 限流对象的键
local window = tonumber(ARGV[1])    -- 窗口大小，单位毫秒
local threshold = tonumber(ARGV[2]) -- 阈值，即窗口时间内允许的最大请求数
local now = tonumber(ARGV[3])       -- 当前时间戳，单位毫秒

-- 2. 计算窗口的起始时间
local min = now - window -- 窗口的起始时间

-- 3. 删除窗口开始之前的所有请求记录
redis.call('ZREMRANGEBYSCORE', key, '-inf', min)

-- 4. 计算当前窗口内的请求数量
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf') -- 计算当前键中的成员数量

-- 5. 判断当前请求数量是否超过阈值
if cnt >= threshold then
    -- 如果超过阈值，执行限流操作
    return "true"
else
    -- 如果未超过阈值，记录当前请求
    -- 使用当前时间戳作为score和member，添加到有序集合中
    redis.call('ZADD', key, now, now)
    -- 设置键的过期时间，以保持滑动窗口的大小
    redis.call('PEXPIRE', key, window)
    return "false"
end
