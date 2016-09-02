all_the_letters = {
  "a" => nil,
  "b" => nil,
  "c" => nil,
  "d" => nil,
  "e" => nil,
  "f" => nil,
  "g" => nil,
  "h" => nil,
  "i" => nil,
  "j" => nil,
  "k" => nil,
  "l" => nil,
  "m" => nil,
  "n" => nil,
  "o" => nil,
  "p" => nil,
  "q" => nil,
  "r" => nil,
  "s" => nil,
  "t" => nil,
  "u" => nil,
  "v" => nil,
  "w" => nil,
  "x" => nil,
  "y" => nil,
  "z" => nil
}

File.open(ARGV[0], 'r') do |file|
  file.each_char do |c|
    if( !all_the_letters[c] )
      all_the_letters[c] = true
    end
  end
end

puts all_the_letters.select { |item, found| found.nil? }.keys
