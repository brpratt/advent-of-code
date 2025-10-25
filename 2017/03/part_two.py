import math


class Memory:
    def __init__(self, rings):
        self.rings = rings
        self.length = (rings * 2) + 1
        self.grid = [[0 for y in range(self.length)] for x in range(self.length)]

    def getval(self, pos):
        y, x = pos
        return self.grid[y][x]

    def setval(self, pos, value):
        y, x = pos
        self.grid[y][x] = value

    def start_pos(self):
        return (self.rings, self.rings)

    def next_pos(self, pos):
        y, x = pos
        y_off, x_off = y - self.rings, x - self.rings
        ring = max(abs(y_off), abs(x_off))
        if x_off == ring:
            if y_off == ring:
                x_off += 1
            elif y_off == -ring:
                x_off -= 1
            else:
                y_off -= 1
        elif x_off == -ring:
            if y_off == ring:
                x_off += 1
            else:
                y_off += 1
        elif y_off == ring:
            x_off += 1
        else:
            x_off -= 1
        return y_off + self.rings, x_off + self.rings

    def neighbors(self, pos):
        y, x = pos
        return [
            (y - 1, x - 1),
            (y - 1, x),
            (y - 1, x + 1),
            (y, x - 1),
            (y, x + 1),
            (y + 1, x - 1),
            (y + 1, x),
            (y + 1, x + 1),
        ]


def memory_accum_until(value):
    memory = Memory(math.floor(math.log2(value)) // 2)
    pos = memory.start_pos()
    memory.setval(pos, 1)
    while memory.getval(pos) < value:
        pos = memory.next_pos(pos)
        neighbors = memory.neighbors(pos)
        acc = sum([memory.getval(x) for x in neighbors])
        memory.setval(pos, acc)
    return memory.getval(pos)


if __name__ == "__main__":
    print(memory_accum_until(289326))
