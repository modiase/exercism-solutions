return function(n)
  if n < 1
    then
    error("Only positive numbers are allowed") 
  end
  steps, current = 0, n
  while true
  do
    if current == 1
      then
      return steps
    end
    if current % 2 == 0
      then
      current = current // 2
    else
      current = current * 3 + 1
    end
    steps = steps + 1
  end
end
