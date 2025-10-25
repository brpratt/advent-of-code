#!/usr/bin/gawk -f

$2 == "X" && $1 == "A" { ascore += 4; bscore += 3; }
$2 == "X" && $1 == "B" { ascore += 1; bscore += 1; }
$2 == "X" && $1 == "C" { ascore += 7; bscore += 2; }

$2 == "Y" && $1 == "A" { ascore += 8; bscore += 4; }
$2 == "Y" && $1 == "B" { ascore += 5; bscore += 5; }
$2 == "Y" && $1 == "C" { ascore += 2; bscore += 6; }

$2 == "Z" && $1 == "A" { ascore += 3; bscore += 8; }
$2 == "Z" && $1 == "B" { ascore += 9; bscore += 9; }
$2 == "Z" && $1 == "C" { ascore += 6; bscore += 7; }

END { 
    print ascore
    print bscore
}
