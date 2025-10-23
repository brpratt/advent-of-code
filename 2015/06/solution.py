import enum
from dataclasses import dataclass


@dataclass
class Point:
    x: int
    y: int


class InstructionType(enum.Enum):
    TURN_ON = 1
    TURN_OFF = 2
    TOGGLE = 3


class Instruction:
    def __init__(self, type: InstructionType, start: Point, stop: Point):
        self.type = type
        self.start = start
        self.stop = stop


def parse_point(raw: str) -> Point:
    s = raw.split(",")
    return Point(int(s[0]), int(s[1]))


def parse(line: str) -> Instruction:
    match line.split():
        case ["turn", "on", *rest]:
            type = InstructionType.TURN_ON
        case ["turn", "off", *rest]:
            type = InstructionType.TURN_OFF
        case ["toggle", *rest]:
            type = InstructionType.TOGGLE
        case _:
            raise Exception(f"could not parse line: {line}")
    return Instruction(type, parse_point(rest[0]), parse_point(rest[2]))


class Grid:
    lights: list[list[int]]

    def __init__(self) -> None:
        self.lights = [[0] * 1000 for _ in range(1000)]

    def do(self, instruction: Instruction) -> None:
        match instruction.type:
            case InstructionType.TURN_ON:
                self._turn_on(instruction.start, instruction.stop)
            case InstructionType.TURN_OFF:
                self._turn_off(instruction.start, instruction.stop)
            case InstructionType.TOGGLE:
                self._toggle(instruction.start, instruction.stop)

    def _turn_on(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                self.lights[y][x] = 1

    def _turn_off(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                self.lights[y][x] = 0

    def _toggle(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                self.lights[y][x] = 1 - self.lights[y][x]

    def brightness(self) -> int:
        total = 0
        for row in self.lights:
            total += sum(row)
        return total


def solve_part_1(lines: list[str]) -> int:
    grid = Grid()
    for line in lines:
        instruction = parse(line)
        grid.do(instruction)
    return grid.brightness()


class Grid2:
    lights: list[list[int]]

    def __init__(self) -> None:
        self.lights = [[0] * 1000 for _ in range(1000)]

    def do(self, instruction: Instruction) -> None:
        match instruction.type:
            case InstructionType.TURN_ON:
                self._turn_on(instruction.start, instruction.stop)
            case InstructionType.TURN_OFF:
                self._turn_off(instruction.start, instruction.stop)
            case InstructionType.TOGGLE:
                self._toggle(instruction.start, instruction.stop)

    def _turn_on(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                self.lights[y][x] += 1

    def _turn_off(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                if self.lights[y][x] > 0:
                    self.lights[y][x] -= 1

    def _toggle(self, start: Point, stop: Point) -> None:
        for y in range(start.y, stop.y + 1):
            for x in range(start.x, stop.x + 1):
                self.lights[y][x] += 2

    def brightness(self) -> int:
        total = 0
        for row in self.lights:
            total += sum(row)
        return total


def solve_part_2(lines: list[str]) -> int:
    grid = Grid2()
    for line in lines:
        instruction = parse(line)
        grid.do(instruction)
    return grid.brightness()


if __name__ == "__main__":
    with open("./2015/06/input.txt", "r") as f:
        lines = f.readlines()
        print(solve_part_1(lines))
        print(solve_part_2(lines))
