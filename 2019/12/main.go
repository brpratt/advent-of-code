package day12

import "fmt"

type moon struct {
	posX int
	posY int
	posZ int
	velX int
	velY int
	velZ int
}

func (m moon) String() string {
	return fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>", m.posX, m.posY, m.posZ, m.velX, m.velY, m.velZ)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func applyGravity(moons []moon) {
	for i := 0; i < len(moons); i++ {
		for j := 0; j < len(moons); j++ {
			moons[i].velX += sgn(moons[j].posX - moons[i].posX)
			moons[i].velY += sgn(moons[j].posY - moons[i].posY)
			moons[i].velZ += sgn(moons[j].posZ - moons[i].posZ)
		}
	}
}

func applyVelocity(moons []moon) {
	for i := range moons {
		moons[i].posX += moons[i].velX
		moons[i].posY += moons[i].velY
		moons[i].posZ += moons[i].velZ
	}
}

func step(moons []moon) {
	applyGravity(moons)
	applyVelocity(moons)
}

func energy(m moon) int {
	return (abs(m.posX) + abs(m.posY) + abs(m.posZ)) * (abs(m.velX) + abs(m.velY) + abs(m.velZ))
}

func SolvePart01(moons []moon, steps int) int {
	for i := 0; i < steps; i++ {
		step(moons)
	}

	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += energy(moon)
	}

	return totalEnergy
}

func matchX(setA, setB []moon) bool {
	for i := 0; i < len(setA); i++ {
		if setA[i].posX != setB[i].posX {
			return false
		}
		if setA[i].velX != setB[i].velX {
			return false
		}
	}

	return true
}

func matchY(setA, setB []moon) bool {
	for i := 0; i < len(setA); i++ {
		if setA[i].posY != setB[i].posY {
			return false
		}
		if setA[i].velY != setB[i].velY {
			return false
		}
	}

	return true
}

func matchZ(setA, setB []moon) bool {
	for i := 0; i < len(setA); i++ {
		if setA[i].posZ != setB[i].posZ {
			return false
		}
		if setA[i].velZ != setB[i].velZ {
			return false
		}
	}

	return true
}

func SolvePart02(moons []moon) int {
	scratch := make([]moon, len(moons))
	copy(scratch, moons)

	cycle := 0
	cycleX := 0
	cycleY := 0
	cycleZ := 0

	for cycleX == 0 || cycleY == 0 || cycleZ == 0 {
		step(scratch)
		cycle++

		if cycleX == 0 && matchX(scratch, moons) {
			cycleX = cycle
		}

		if cycleY == 0 && matchY(scratch, moons) {
			cycleY = cycle
		}

		if cycleZ == 0 && matchZ(scratch, moons) {
			cycleZ = cycle
		}
	}

	return lcm(cycleX, lcm(cycleY, cycleZ))
}

func Solve(part int) int {
	moons := []moon{
		{posX: 1, posY: 3, posZ: -11},
		{posX: 17, posY: -10, posZ: -8},
		{posX: -1, posY: -15, posZ: 2},
		{posX: 12, posY: -4, posZ: -4},
	}

	if part == 1 {
		return SolvePart01(moons, 1000)
	}

	return SolvePart02(moons)
}
