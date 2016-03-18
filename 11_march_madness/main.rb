def game_count_by_team_count(team_count)
  if team_count > 2
    game_count = team_count / 2
    game_count + game_count_by_team_count(game_count)
  elsif team_count == 2
    1
  else
    raise "Invalid team_count"
  end
end

game_count = game_count_by_team_count 64
puts "Games: #{game_count}"

combinations = 2**game_count
puts "Combinations: #{combinations}"
