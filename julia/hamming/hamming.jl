function distance(s1::AbstractString, s2::AbstractString)
  distance = 0
  if length(s1) != length(s2) 
    throw(ArgumentError("Strings must be same length"))
  end

  for (index, value) in enumerate(s1)
    if value != s2[index]
      distance += 1
    end
  end
  distance
end
