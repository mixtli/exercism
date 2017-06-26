module BookKeeping
  VERSION = 4
end

class Complement
  DNAMAP = {
      'G' => 'C',
      'C' => 'G',
      'T' => 'A',
      'A' => 'U'
  }
  def self.of_dna(strand)
    strand.chars.map {|c| return "" unless DNAMAP[c]; DNAMAP[c] }.join('') 
  end
end
