local lasagna = {}

-- add a 'oven_time' field to the lasagna table
lasagna.oven_time = 40
-- returns the remaining minutes based on the actual minutes in the oven
function lasagna.remaining_oven_time(actual_minutes_in_oven)
  return lasagna.oven_time - actual_minutes_in_oven
end

-- calculates and returns the time needed to prepare the lasagna on the amount
-- of layers
lasagna.time_per_layer_minutes = 2
function lasagna.preparation_time(number_of_layers)
  return number_of_layers * lasagna.time_per_layer_minutes
end

-- calculates the time elapsed cooking the lasagna - including preparation time
-- and baking time.
function lasagna.elapsed_time(number_of_layers, actual_minutes_in_oven)
  return lasagna.preparation_time(number_of_layers) + actual_minutes_in_oven
end

return lasagna
