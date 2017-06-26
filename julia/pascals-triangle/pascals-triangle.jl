function triangle(n::Int)
  if n < 0
    throw(DomainError())
  end
  if n == 0
    return []
  end
  if n == 1
    return [[1]]
  end
  prev = triangle(n-1)
  lastrow = copy(prev[end])
  push!(lastrow,0)
  unshift!(lastrow,0)
  newrow = []
  for i in 1:(length(lastrow) - 1)
    push!(newrow, lastrow[i] + lastrow[i+1])
  end
  push!(prev, newrow)
  return prev
end
