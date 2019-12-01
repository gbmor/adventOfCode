{:ok, input} = File.read("input.txt")

input
|> String.split("\n", trim: true)
|> Enum.map(&String.to_integer/1)
|> Enum.sum()
|> IO.puts()
