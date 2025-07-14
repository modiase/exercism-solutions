return {
  value = function(colors)
    _colors = {
      black = 0;
      brown = 1;
      red = 2;
      orange = 3;
      yellow = 4;
      green = 5;
      blue = 6;
      violet = 7;
      grey = 8;
      white = 9;
    }
    local first = _colors[colors[1]]
    local second = _colors[colors[2]]
    return 10 * first + second
  end
}
