#!/usr/bin/gawk -f

function stack_push(s, v) {
    _n = s[0]
    s[_n+1] = v
    s[0] += 1
}

function stack_pop(s) {
    _n = s[0]
    _v = s[_n]
    s[0] -= 1
    return _v
}

function path(s) {
    _n = s[0]
    if (_n == 0) {
        return "/"
    }

    _p = ""
    for (_i = 1; _i <= _n; _i++) {
        _p = _p "/" s[_i]
    }
    return _p
}

BEGIN {
    FS = " "
    pwd[0] = 0
}

$1 == "$" && $2 == "cd" && $3 == "/" {
    next
}

$1 == "$" && $2 == "cd" && $3 == ".." {
    stack_pop(pwd)
    next
}

$1 == "$" && $2 == "cd" {
    stack_push(pwd, $3)
}

$1 ~ /[0-9]+/ {
    sizes["/"] += $1
    dir = ""

    for (i = 1; i <= pwd[0]; i++) {
        dir = dir "/" pwd[i]
        sizes[dir] += $1
    }
}

END {
    part_a = 0
    part_b = 70000000

    unused = 70000000 - sizes["/"]
    needed = 30000000 - unused

    for (d in sizes) {
        if (sizes[d] < 100000) {
            part_a += sizes[d]
        }
        if ((sizes[d] >= needed) && (sizes[d] < part_b)) {
            part_b = sizes[d]
        }
    }

    print part_a
    print part_b
}
