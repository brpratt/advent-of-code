#!/usr/bin/gawk -f

BEGIN {
    FS = ""

    _items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for (i = 1; i <= length(_items); i++) {
        values[substr(_items, i, 1)] = i
    }
}

{
    sacks[NR % 3] = $0

    for (i = 1; i <= (NF / 2); i++) {
       items[$i] = 1
    }

    for (i = (NF / 2) + 1; i <= NF; i++) {
        if ($i in items) {
            part_a += values[$i]
            break
        }
    }

    delete items
}

NR % 3 == 0 {
    delete first
    delete second

    split(sacks[0], items)

    for (item in items) {
        first[items[item]] = 1
    }

    split(sacks[1], items)

    for (item in items) {
        if (items[item] in first) {
            second[items[item]] = 1
        }
    }

    split(sacks[2], items)

    for (item in items) {
        if (items[item] in second) {
            part_b += values[items[item]]
            delete sacks
            break
        }
    }
}

END {
    print part_a
    print part_b
}
