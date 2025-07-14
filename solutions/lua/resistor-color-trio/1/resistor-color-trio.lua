return {
  decode = function(c1, c2, c3)
	local _colors = {
		black = 0,
		brown = 1,
		red = 2,
		orange = 3,
		yellow = 4,
		green = 5,
		blue = 6,
		violet = 7,
		grey = 8,
		white = 9,
	}

	local _units = { "ohms", "kiloohms", "megaohms", "gigaohms" }

	local exp = _colors[c3]
	local value = (((_colors[c1] * 10) + _colors[c2]) * 10 ^ exp)
	if value == 0
	then
		return 0, "ohms"
	end
	local idx = math.floor((math.log(value) / math.log(10)) / 3)
	return (value / 10 ^ (3 * idx)), _units[idx + 1]
end


}
