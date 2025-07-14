local triangle = {}
function triangle.validate(a, b, c)
  has_no_sides = (a + b + c) == 0
  has_negative_side = a < 0 or b < 0 or c < 0
  violates_triangle_inequality = (a > b + c) or (b > a + c) or (c > a + b)
  if has_negative_side or violates_triangle_inequality or has_no_sides
    then
    error('Input Error')
  end
end
function triangle.kind(a, b, c)
  triangle.validate(a,b,c)
  ab = a == b
  bc = b == c
  ac = a == c
  if not ab and not bc and not ac
    then
    return 'scalene'
  elseif ab and bc and ac
      then
      return 'equilateral'
  end
  return 'isosceles'
end

return triangle
