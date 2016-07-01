dim = $stdin.gets.to_i
src_grid = $stdin.readlines.map { |l| l.chomp.chars }

# Add sentinal values to make like easier later on
dim += 1
src_grid = src_grid.map { |l| l << "X" }
src_grid << Array.new(dim, "X")

puts "Initial map with sentinals"
puts src_grid.map { |l| l.map { |c| c.to_s.ljust(2) }.join }
puts

x_space_grid = Array.new(dim) { Array.new dim, 0 }

# Fill open grid with x dimension run lengths
dim.times do |y|
  first_open = nil
  dim.times do |x|
    if src_grid[y][x] == "-"
      first_open = x unless first_open
    else
      if first_open
        (first_open..x).each do |x2|
          x_space_grid[y][x2] = x-x2
        end
      end
      first_open = nil
    end
  end
end

puts "Open x run lengths"
puts x_space_grid.map { |l| l.map { |c| c.to_s.ljust(3) }.join }
puts

max_square_pos = nil
max_square_size = 0
space_grid = Array.new(dim) { Array.new dim, 0 }

debug = false

dim.times do |y|
  dim.times do |x|
    max_size_from_here = x_space_grid[y][x]
    debug = x==14 && y==1
    puts "max_size_from_here: #{max_size_from_here}" if debug
    next if max_size_from_here <= max_square_size
    size = 0
    puts (y..(y+max_size_from_here)) if debug
    (y..(y+max_size_from_here)).each do |y2|
      current_width = x_space_grid[y2][x]
      puts "x: #{x}, y2: #{y2}, current_width: #{current_width}, size: #{size}, max_size_from_here: #{max_size_from_here}" if debug
      break if current_width <= size || size == max_size_from_here
      size += 1
      max_size_from_here = current_width if current_width < max_size_from_here
    end

    if size > max_square_size
      max_square_size = size
      max_square_pos = [x,y]
    end
  end
end

src_grid[max_square_pos[1], max_square_size].each do |line|
  line[max_square_pos[0], max_square_size] = Array.new(max_square_size, "O")
end

puts "With largest open square marked"
puts src_grid.map { |l| l.map { |c| c.to_s.ljust(2) }.join }
puts

puts "Biggest square at #{max_square_pos.join(", ")}: #{max_square_size}"
