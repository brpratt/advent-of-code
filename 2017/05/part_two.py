class Program:
    def __init__(self, instructions):
        self.instructions = instructions
        self.index = 0

    def next_index(self):
        return self.index + self.instructions[self.index]

    def can_jump(self):
        return self.next_index() < len(self.instructions)

    def jump(self):
        next_index = self.next_index()
        if self.instructions[self.index] >= 3:
            self.instructions[self.index] -= 1
        else:
            self.instructions[self.index] += 1
        self.index = next_index


def count_program_steps(instructions):
    count = 0
    program = Program(instructions)
    while program.can_jump():
        program.jump()
        count += 1
    return count + 1


def count_program_steps_from_file():
    with open("./2017/05/input.txt") as program_file:
        lines = program_file.readlines()
        instructions = [int(x) for x in lines if x.strip()]
        return count_program_steps(instructions)


if __name__ == "__main__":
    print(count_program_steps_from_file())
