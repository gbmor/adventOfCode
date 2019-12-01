defmodule Fuel do
  def calc_naive(input) do
    input
    |> Enum.map(fn x -> floor(x / 3) - 2 end)
    |> Enum.sum()
  end

  def calc_for_each(input) do
    input
    |> Enum.map(fn x -> calc_individual(x) end)
    |> Enum.sum()
  end

  def calc_individual(mass, acc \\ 0) do
    tmp = floor(mass / 3) - 2

    if tmp < 1 do
      acc
    else
      acc = acc + tmp
      calc_individual(tmp, acc)
    end
  end
end

data =
  File.read!("input.txt")
  |> String.split("\n", trim: true)
  |> Enum.map(&String.to_integer/1)

Fuel.calc_naive(data) |> IO.puts()
Fuel.calc_for_each(data) |> IO.puts()
