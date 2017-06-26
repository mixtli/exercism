module BookKeeping
  VERSION = 5 # Where the version number matches the one in the test.
end

class Gigasecond
  def self.from(time)
    return time + 1000000000
  end
end
