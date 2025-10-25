def distribute(banks, offset, value):
    while value > 0:
        banks[offset] += 1
        value -= 1
        offset = (offset + 1) % len(banks)


def redistribute(banks):
    max_index = 0
    index = 1
    while index < len(banks):
        if banks[index] > banks[max_index]:
            max_index = index
        index += 1
    value = banks[max_index]
    banks[max_index] = 0
    distribute(banks, (max_index + 1) % len(banks), value)


def loop_count(banks):
    count = 0
    dist_set = set()
    curr_set = tuple(banks)
    while curr_set not in dist_set:
        dist_set.add(curr_set)
        redistribute(banks)
        curr_set = tuple(banks)
        count += 1
    return count


def cycle_count(banks):
    count_to_loop = loop_count(list(banks))
    while count_to_loop > 0:
        redistribute(banks)
        count_to_loop -= 1
    count = 0
    dist_set = set()
    curr_set = tuple(banks)
    while curr_set not in dist_set:
        dist_set.add(curr_set)
        redistribute(banks)
        curr_set = tuple(banks)
        count += 1
    return count


def cycle_count_from_file():
    with open("./2017/06/input.txt") as banks_file:
        banks = [int(x) for x in banks_file.readline().strip().split()]
        return cycle_count(banks)


if __name__ == "__main__":
    print(cycle_count_from_file())
