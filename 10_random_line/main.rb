#!/usr/bin/env ruby -w

# Week 10: Random Line: How could you select one of ​_n_​ objects at random,
# where you see the objects sequentially but you don’t know the value of ​_n_
# beforehand? For concreteness, how would you read a text file, and select and
# print one random line, when you don’t know the number of lines in advance?

random_line = nil
count = 0

while line = gets
  count += 1
  if rand(count) == 0
    random_line = line
  end
end

puts random_line
