import enum
from dataclasses import dataclass
from typing import Sequence

Wire = str


class Op(enum.Enum):
    CONSTANT = 1
    NOT = 2
    AND = 3
    OR = 4
    LSHIFT = 5
    RSHIFT = 6


@dataclass
class Connection:
    out: Wire
    op: Op
    in1: str
    in2: str


def parse(line: str) -> Connection:
    match line.split():
        case [in1, "->", out]:
            return Connection(out, Op.CONSTANT, in1, "")
        case ["NOT", in1, "->", out]:
            return Connection(out, Op.NOT, in1, "")
        case [in1, "AND", in2, "->", out]:
            return Connection(out, Op.AND, in1, in2)
        case [in1, "OR", in2, "->", out]:
            return Connection(out, Op.OR, in1, in2)
        case [in1, "LSHIFT", in2, "->", out]:
            return Connection(out, Op.LSHIFT, in1, in2)
        case [in1, "RSHIFT", in2, "->", out]:
            return Connection(out, Op.RSHIFT, in1, in2)
        case _:
            raise Exception(f"could not parse line: {line}")


class Circuit:
    connections: dict[Wire, Connection]
    signals: dict[Wire, int]

    def __init__(self) -> None:
        self.connections = {}
        self.signals = {}

    def add_connection(self, connection: Connection) -> None:
        self.connections[connection.out] = connection

    def _signal(self, w: str) -> int:
        if w.isdigit():
            return int(w)
        else:
            return self.signal(w)

    def signal(self, wire: Wire) -> int:
        if wire in self.signals:
            return self.signals[wire]
        connection = self.connections[wire]
        match connection.op:
            case Op.CONSTANT:
                self.signals[wire] = self._signal(connection.in1)
            case Op.NOT:
                in1 = self._signal(connection.in1)
                self.signals[wire] = ~in1 & 0xFFFF
            case Op.AND:
                in1 = self._signal(connection.in1)
                in2 = self._signal(connection.in2)
                self.signals[wire] = (in1 & in2) & 0xFFFF
            case Op.OR:
                in1 = self._signal(connection.in1)
                in2 = self._signal(connection.in2)
                self.signals[wire] = (in1 | in2) & 0xFFFF
            case Op.LSHIFT:
                in1 = self._signal(connection.in1)
                in2 = int(connection.in2)
                self.signals[wire] = (in1 << in2) & 0xFFFF
            case Op.RSHIFT:
                in1 = self._signal(connection.in1)
                in2 = int(connection.in2)
                self.signals[wire] = (in1 >> in2) & 0xFFFF
        return self.signals[wire]


def solve_part_1(lines: Sequence[str], wire: str) -> int:
    circuit = Circuit()
    for line in lines:
        connection = parse(line)
        circuit.add_connection(connection)
    return circuit.signal(wire)


def solve_part_2(lines: Sequence[str], wire: str) -> int:
    circuit = Circuit()
    for line in lines:
        connection = parse(line)
        circuit.add_connection(connection)
    circuit.signals["b"] = 46065
    return circuit.signal(wire)


if __name__ == "__main__":
    with open("./2015/07/input.txt", "r") as f:
        lines = f.readlines()
        print(solve_part_1(lines, "a"))
        print(solve_part_2(lines, "a"))
