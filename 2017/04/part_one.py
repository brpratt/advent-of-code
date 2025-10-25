def is_valid(phrase):
    word_set = set()
    for word in phrase.split():
        if word in word_set:
            return False
        word_set.add(word)
    return True


def num_valid_phrases():
    count = 0
    with open("./2017/04/input.txt") as phrase_file:
        for line in phrase_file:
            if is_valid(line.strip()):
                count += 1
    return count


if __name__ == "__main__":
    print(num_valid_phrases())
