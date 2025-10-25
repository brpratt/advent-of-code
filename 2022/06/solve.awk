#!/usr/bin/gawk -f

function set_add(s, v) {
    if (s[v] == "") {
        s[0] += 1
    }

    s[v] += 1
}

function set_count(s) {
    return s[0]
}

function set_init(s) {
    s[0] = 0
}

BEGIN {
    FS = ""
}

{
    for (i = 1; i <= NF; i++) {
        delete chars
        set_init(chars)

        for (j = 0; j < 4; j++) {
            set_add(chars, $(i + j))
        }

        if (set_count(chars) == 4) {
            print i + 3
            break
        }
    }

    for (i = 1; i <= NF; i++) {
        delete chars
        set_init(chars)

        for (j = 0; j < 14; j++) {
            set_add(chars, $(i + j))
        }

        if (set_count(chars) == 14) {
            print i + 13
            break
        }
    }
}
