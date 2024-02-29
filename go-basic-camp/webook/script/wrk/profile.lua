wrk.method = "GET"
wrk.headers["Content-Type"] = "application/json"
wrk.headers["User-Agent"] = "RapidAPI/4.2.0 (Macintosh; OS X/14.3.1) GCDHTTPRequest"
-- 记得修改这个，你在登录页面登录一下，然后复制一个过来这里
wrk.headers["Authorization"] =
"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MywiVXNlckFnZW50IjoiUmFwaWRBUEkvNC4yLjAgKE1hY2ludG9zaDsgT1MgWC8xNC4zLjEpIEdDREhUVFBSZXF1ZXN0IiwiZXhwIjoxNzA5MTU5MzQzfQ.wj5TIVM4yvlLJiRl1aGxDs5chITHbT3JvadLSKYBrOc"
