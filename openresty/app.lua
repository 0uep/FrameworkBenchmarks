local mysql = require "resty.mysql"
local math = require "math"

local encode = require("cjson").encode
local random = math.random
local insert = table.insert


local mysqlconn = {
	host = "127.0.0.1",
	port = 3306,
	database = "hello_world",
	user = "benchmarkdbuser",
	password = "benchmarkdbpass"
}
local db = mysql:new()
return function()
	db:connect(mysqlconn)
	local num_queries = tonumber(ngx.var.arg_queries) or 1
	local worlds = {}
	for i=1, num_queries do
		local wid = random(1, 10000)
		insert(worlds, db:query('SELECT * FROM World WHERE id = '..wid)[1])
	end
	ngx.print( encode(worlds) )
	db:set_keepalive(0, 256)
end
