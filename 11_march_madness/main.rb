def game_count_by_team_count(team_count)
  team_count - 1
end

game_count = game_count_by_team_count 64
puts "Games: #{game_count}"

combinations = 2**game_count
puts "Combinations: #{combinations}"
