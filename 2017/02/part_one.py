def row_difference(row):
    return max(row) - min(row)


def calculate_checksum(sheet):
    diffs = [row_difference(x) for x in sheet]
    return sum(diffs)


def calculate_checksum_from_file():
    with open("./2017/02/input.txt") as sheet_file:
        lines = sheet_file.readlines()
        sheet = [[int(x) for x in row.strip().split()] for row in lines]
        return calculate_checksum(sheet)


if __name__ == "__main__":
    print(calculate_checksum_from_file())
