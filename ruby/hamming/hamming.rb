module BookKeeping
  VERSION = 1 # Where the version number matches the one in the test.
end

module Hamming
  extend self
  def compute(str1, str2)
    distance = 0
    str1.split('').each_with_index do |c, i|
      distance += 1 if c != str2[i] 
    end
  end
end
