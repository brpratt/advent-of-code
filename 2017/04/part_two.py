def is_anagram(base_str, test_str):
    letter_dict = {}
    for letter in base_str:
        if letter not in letter_dict:
            letter_dict[letter] = 0
        letter_dict[letter] += 1
    for letter in test_str:
        if letter not in letter_dict:
            letter_dict[letter] = 0
        letter_dict[letter] -= 1
    return all([letter_dict[letter] == 0 for letter in letter_dict])


def is_valid(phrase):
    word_set = set()
    for test_str in phrase.split():
        for base_str in word_set:
            if is_anagram(base_str, test_str):
                return False
        word_set.add(test_str)
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
