import hashlib


def md5(s: bytes) -> str:
    h = hashlib.md5()
    h.update(s)
    return h.hexdigest()


def secret(key: str, num: int) -> bytes:
    return f"{key}{num}".encode("ascii")


def solve_part_1(line: str) -> int:
    number = 0
    while not md5(secret(line, number)).startswith("00000"):
        number += 1
    return number


def solve_part_2(line: str) -> int:
    number = 0
    while not md5(secret(line, number)).startswith("000000"):
        number += 1
    return number


if __name__ == "__main__":
    with open("./2015/04/input.txt", "r") as f:
        line = f.readline().strip()
        print(solve_part_1(line))
        print(solve_part_2(line))
