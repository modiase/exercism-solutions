
function mod(a, b)
  return a - math.floor(a/b)*b
end

return function(n)
  result = ""
  if mod(n, 3) == 0
    then
    result = result .. "Pling"
  end

  if mod(n, 5) == 0
    then
    result = result .. "Plang"
  end

  if mod(n, 7) == 0
    then
    result = result .. "Plong"
  end

  if result == ""
    then
    return tostring(n)
  end
  return result
  
end
