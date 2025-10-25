#!/usr/bin/gawk -f

function contains(range1, range2) {
    if (range1[1] <= range2[1] && range1[2] >= range2[2]) {
        return 1
    }

    if (range2[1] <= range1[1] && range2[2] >= range1[2]) {
        return 1
    }

    return 0
}

function overlaps(range1, range2) {
    if (range1[2] < range2[1] || range2[2] < range1[1]) {
        return 0
    }

    return 1
}

BEGIN { FS = "," }

{
    split($1, range1, "-")
    split($2, range2, "-")

    if (contains(range1, range2)) {
        partA += 1
    }

    if (overlaps(range1, range2)) {
        partB += 1
    }
}

END {
    print partA
    print partB
}
