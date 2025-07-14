local EliudsEggs = {}

function EliudsEggs.egg_count(number)
  local n, e, count = number, 0, 0
  while (n ~= 0) do
    count =  count + ((n % 2 == 0) and 0 or 1)
    n = n // 2
    e = e + 1
  end
  return count
end

return EliudsEggs
