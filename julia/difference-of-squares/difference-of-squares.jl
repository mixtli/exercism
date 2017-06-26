# The descriptions for these are backward in the problem.

"Square the sum of the numbers up to the given number"
function sum_of_squares(n::Int)
  reduce(+, map((v) -> v^2, 1:n))
end

"Sum the squares of the numbers up to the given number"
function square_of_sum(n::Int)
  reduce(+, 1:n)^2
end

"Subtract sum of squares from square of sums"
function difference(n::Int)
  square_of_sum(n) - sum_of_squares(n)
end
