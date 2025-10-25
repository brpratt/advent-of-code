import math


def ring_level(num):
    ring = math.floor(math.sqrt(num))
    if ring**2 == num:
        return ring
    else:
        return ring + 1


def ring_span(ring):
    if ring % 2 == 0:
        return ring // 2
    else:
        return (ring - 1) // 2


def ring_max(ring):
    return ring**2


def ring_min(ring):
    return ((ring - 1) ** 2) + 1


def ring_mid(ring):
    return (ring_max(ring) + ring_min(ring)) // 2


def steps(num):
    ring = ring_level(num)
    mid = ring_mid(ring)
    span = ring_span(ring)
    offset = abs(num - mid)
    if offset > span:
        offset -= (offset - span) * 2
    return (span * 2) - offset


if __name__ == "__main__":
    print(steps(289326))
