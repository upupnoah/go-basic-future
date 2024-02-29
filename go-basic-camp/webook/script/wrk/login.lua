-- 使用 wrk 发送 POST 请求进行用户登录测试
wrk.method = "POST"
wrk.headers["Content-Type"] = "application/json"
-- 替换为登录信息
wrk.body = '{"email":"test@gmail.com", "password": "TEST12345678!!"}'
