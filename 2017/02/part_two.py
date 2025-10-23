def row_quotient(row):
    i = 0
    while i < len(row) - 1:
        j = i + 1
        while j < len(row):
            if row[i] >= row[j]:
                x, y = row[i], row[j]
            else:
                x, y = row[j], row[i]
            if x % y == 0:
                return x // y
            j += 1
        i += 1
    return 0


def calculate_checksum(sheet):
    results = [row_quotient(x) for x in sheet]
    return sum(results)


def calculate_checksum_from_file():
    with open("./2017/02/input.txt") as sheet_file:
        lines = sheet_file.readlines()
        sheet = [[int(x) for x in row.strip().split()] for row in lines]
        return calculate_checksum(sheet)


if __name__ == "__main__":
    print(calculate_checksum_from_file())
