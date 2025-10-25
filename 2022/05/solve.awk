#!/usr/bin/gawk -f

function stack_top(s) {
    return s[s[0]]
}

function stack_len(s) {
    return s[0]
}

function stack_push(s, v) {
    n = s[0]
    s[n+1] = v
    s[0] += 1
}

function stack_pop(s) {
    n = s[0]
    v = s[n]
    s[0] -= 1
    return v
}

BEGIN {
    FS = ""

    init_height = 0

    while (getline < ARGV[1]) {
        if ($2 == 1) {
            num_stacks = $(NF-1)
            break
        }
        init_height += 1
    }

    for (i = 1; i <= num_stacks; i++) {
        stacks_a[i][0] = init_height
        stacks_b[i][0] = init_height
    }

    load = 1
    set = 0
    height = init_height
}

load == 1 {
    for (i = 1; i <= num_stacks; i++) {
        stacks_a[i][height] = $((i * 4) - 2)
        stacks_b[i][height] = $((i * 4) - 2)
    }

    height -= 1
    if (height == 0) {
        load = 0
        set = 1
    }
}

set == 1 {
    for (i = 1; i <= num_stacks; i++) {
        for (j = 1; j <= init_height; j++) {
            if (stacks_a[i][j] == " ") {
                stacks_a[i][0] = j - 1
                stacks_b[i][0] = j - 1
                break
            }
        }
    }

    set = 0
    FS = " "
}

$1 == "move" {
    temp_b[0] = 0

    for (i = 0; i < $2; i++) {
        v = stack_pop(stacks_a[$4])
        stack_push(stacks_a[$6], v)

        v = stack_pop(stacks_b[$4])
        stack_push(temp_b, v)
    }

    for (i = 0; i < $2; i++) {
        v = stack_pop(temp_b)
        stack_push(stacks_b[$6], v)
    }
}

END {
    for (i = 1; i <= num_stacks; i++) {
        printf("%c", stack_top(stacks_a[i]))
    }
    print ""

    for (i = 1; i <= num_stacks; i++) {
        printf("%c", stack_top(stacks_b[i]))
    }
    print ""
}
