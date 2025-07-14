local cars = {}

-- returns the amount of working cars produced by the assembly line every hour
function cars.calculate_working_cars_per_hour(production_rate, success_rate)
  return (production_rate*success_rate) / 100
end

-- returns the amount of working cars produced by the assembly line every minute
function cars.calculate_working_cars_per_minute(production_rate, success_rate)
  return cars.calculate_working_cars_per_hour(production_rate, success_rate) // 60
end

cars.cost_per_car_per_qty = {}
cars.cost_per_car_per_qty[10] = 95000
cars.cost_per_car_per_qty[1] = 10000
-- returns the cost of producing the given number of cars
function cars.calculate_cost(cars_count)
  return cars.cost_per_car_per_qty[1] * (cars_count % 10) + cars.cost_per_car_per_qty[10] * (cars_count // 10)
end

return cars
