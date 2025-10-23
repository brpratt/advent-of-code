def find_matches(digits):
    matches = []
    index = 0
    while index < len(digits) - 1:
        if digits[index] == digits[index + 1]:
            matches.append(digits[index])
        index += 1
    if digits[-1] == digits[0]:
        matches.append(digits[-1])
    return matches


def solve_captcha(captcha):
    digits = [int(x) for x in captcha]
    return sum(find_matches(digits))


def solve_captcha_from_file():
    with open("./2017/01/input.txt") as capcha_file:
        return solve_captcha(capcha_file.readline().strip())


if __name__ == "__main__":
    print(solve_captcha_from_file())
