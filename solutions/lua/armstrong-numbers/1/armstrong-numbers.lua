local ArmstrongNumbers = {}

function ArmstrongNumbers.is_armstrong_number(number)
  if number == 0
    then
    return true
  end
  
  local digits, d, n = {}, 0, number
  
  while (n ~= 0) do
    digits[d] = n % 10
    d = d + 1
    n = n // 10
  end

  sum = 0
  for _, k in pairs(digits) 
  do
    sum = sum + k^d
  end
  
  return sum == number
end

return ArmstrongNumbers
