#!/usr/bin/gawk -f

BEGIN {
    FS = ""
}

{
    max_y = NR
    max_x = NF

    for (i = 1; i <= NF; i++) {
        trees[NR][i] = $i
    }
}

function grade(tx, ty, result) {
    _vis_up = 1
    _vis_left = 1
    _vis_down = 1
    _vis_right = 1

    _dist_up = 0
    _dist_left = 0
    _dist_down = 0
    _dist_right = 0

    # looking up
    for (_y = (ty - 1); _y >= 1; _y--) {
        _dist_up++

        if (trees[ty][tx] <= trees[_y][tx]) {
            _vis_up = 0
            break
        }
    }

    # looking left
    for (_x = (tx - 1); _x >= 1; _x--) {
        _dist_left++

        if (trees[ty][tx] <= trees[ty][_x]) {
            _vis_left = 0
            break
        }
    }

    # looking down
    for (_y = (ty + 1); _y <= max_y; _y++) {
        _dist_down++

        if (trees[ty][tx] <= trees[_y][tx]) {
            _vis_down = 0
            break
        }
    }

    # looking right
    for (_x = (tx + 1); _x <= max_x; _x++) {
        _dist_right++

        if (trees[ty][tx] <= trees[ty][_x]) {
            _vis_right = 0
            break
        }
    }

    result["visible"] = _vis_up || _vis_left || _vis_right || _vis_down
    result["score"] = _dist_up * _dist_left * _dist_right * _dist_down
}

END {
    part_a = 0
    part_b = 0

    for (y = 1; y <= max_y; y++) {
        for (x = 1; x <= max_x; x++) {
            grade(x, y, result)

            if (result["visible"]) {
                part_a++
            }

            if (result["score"] > part_b) {
                part_b = result["score"]
            }
        }
    }

    print part_a
    print part_b
}
