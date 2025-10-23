def solve_part_1(line: str) -> int:
    floor = 0
    for char in line:
        if char == "(":
            floor += 1
        if char == ")":
            floor -= 1
    return floor


def solve_part_2(line: str) -> int:
    floor = 0
    for idx, char in enumerate(line):
        if char == "(":
            floor += 1
        if char == ")":
            floor -= 1
        if floor < 0:
            return idx + 1
    return -1


if __name__ == "__main__":
    with open("./2015/01/input.txt", "r") as f:
        line = f.readline()
        print(solve_part_1(line))
        print(solve_part_2(line))
