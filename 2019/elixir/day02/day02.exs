defmodule Opcodes do
  def get_input do
    File.read!("input.txt")
    |> String.trim()
    |> String.split(",", trim: true)
    |> Enum.map(&String.to_integer/1)
  end

  def exec(input, index \\ 0) do
    a_indx = Enum.at(input, index + 1, 0)
    b_indx = Enum.at(input, index + 2, 0)
    dest = Enum.at(input, index + 3, 0)

    case Enum.at(input, index, nil) do
      1 ->
        n = Enum.at(input, a_indx, nil) + Enum.at(input, b_indx, nil)
        new_input = List.update_at(input, dest, fn _ -> n end)

        exec(new_input, index + 4)

      2 ->
        n = Enum.at(input, a_indx, nil) * Enum.at(input, b_indx, nil)
        new_input = List.update_at(input, dest, fn _ -> n end)

        exec(new_input, index + 4)

      99 ->
        Enum.at(input, 0, nil)
    end
  end
end

Opcodes.get_input() |> Opcodes.exec() |> IO.puts()
