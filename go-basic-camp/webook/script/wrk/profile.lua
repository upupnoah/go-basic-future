-- Initialize variables
token = nil           -- Variable to store the JWT token once received
path = "/users/login" -- API path for the login endpoint
method = "POST"       -- HTTP method for the login request

-- Set headers for the HTTP request
wrk.headers["Content-Type"] = "application/json" -- Specify the content type as JSON
wrk.headers["User-Agent"] = ""                   -- Setting User-Agent header as empty

-- Function to prepare and return the HTTP request
request = function()
    -- Body of the request with login credentials
    body = '{"email": "test@gmail.com", "password": "TEST12345678!!"}'
    -- Format the request with method, path, headers, and body, then return it
    return wrk.format(method, path, wrk.headers, body)
end

-- Function to handle the response from the server
response = function(status, headers, body)
    -- Check if we haven't received a token yet and the status code is 200 (OK)
    if not token and status == 200 then
        -- Extract the JWT token from the response headers
        token = headers["X-Jwt-Token"]
        -- Update the path to the profile endpoint for subsequent requests
        path = "/users/profile"
        -- Change the method to GET for fetching profile information
        method = "GET"
        -- Update the Authorization header to include the received JWT token
        wrk.headers["Authorization"] = string.format("Bear %s", token)
    end
end
