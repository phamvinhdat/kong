local Auth = {
    PRIORITY = 2000,
    VERSION = "0.1.0"
}

local get_header = kong.request.get_header
local set_header = kong.service.request.set_header

-- Run this when the client request hits the service
function Auth:access(conf)
    authorization = get_header(conf.request_header)
    if authorization ~= conf.default_uid then
        local headers = {
            ["Content-Type"] = "application/json; charset=utf-8"
        }

        body = [[
{
  "_errorMessage": "unauthorized"
}]]

        return kong.response.exit(401, body, headers)
    end

    set_header(conf.forward_header, authorization)
end

return Auth
