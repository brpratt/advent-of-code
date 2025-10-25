#!/usr/bin/gawk -f

BEGIN { elf = 0 }
/^$/  { elf += 1 }
      { calories[elf] += $1 }

END {
    n = asort(calories)
    print calories[n]
    print calories[n] + calories[n-1] + calories[n-2]
}
