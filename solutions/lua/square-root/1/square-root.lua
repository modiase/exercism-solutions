local SquareRoot = {}

function SquareRoot.square_root(radicand)
  if radicand == 1
    then
    return 1
  end
  
  local epsilon, e = 0.000001, 0
  while 10^e <= radicand
  do
    e = e + 1
  end
  e = e - 1

  local xprev = (radicand / 10^(e)) * 10^(e//2) * ((e % 2 == 0) and 1 or 3.3)
  local x = (xprev - radicand/xprev)
  while math.abs(x - xprev) > epsilon
  do
    xprev = x
    x = (x + radicand/x)/2
  end

  return x
  
end

return SquareRoot
