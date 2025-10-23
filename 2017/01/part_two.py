def has_match(digits, index):
    match_offset = len(digits) // 2
    match_index = (index + match_offset) % len(digits)
    return digits[index] == digits[match_index]


def find_matches(digits):
    matches = []
    index = 0
    while index < len(digits):
        if has_match(digits, index):
            matches.append(digits[index])
        index += 1
    return matches


def solve_captcha(captcha):
    digits = [int(x) for x in captcha]
    return sum(find_matches(digits))


def solve_captcha_from_file():
    with open("./2017/01/input.txt") as capcha_file:
        return solve_captcha(capcha_file.readline().strip())


if __name__ == "__main__":
    print(solve_captcha_from_file())
