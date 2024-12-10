package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
)

const emptyID = -1

type filesystem []int

func parseInput(r io.Reader) filesystem {
	var fs filesystem
	var id int
	var inFreeSpace bool

	reader := bufio.NewReader(r)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return fs
		}

		num := int(b - '0')
		for num > 0 {
			if inFreeSpace {
				fs = append(fs, emptyID)
			} else {
				fs = append(fs, id)
			}
			num--
		}

		if !inFreeSpace {
			id++
		}

		inFreeSpace = !inFreeSpace
	}
}

func compact(fs filesystem) {
	slot, file := 0, len(fs)-1

	for {
		for slot < len(fs) && fs[slot] != emptyID {
			slot++
		}

		for file >= 0 && fs[file] == emptyID {
			file--
		}

		if slot >= file {
			break
		}

		fs[slot], fs[file] = fs[file], fs[slot]
		slot++
		file--
	}
}

func checksum(fs filesystem) int {
	var cksum int

	for i, id := range fs {
		if id == emptyID {
			continue
		}

		cksum += i * id
	}

	return cksum
}

func part01(fs filesystem) int {
	compact(fs)
	return checksum(fs)
}

func compactWithoutFragmentation(fs filesystem) {
	fileEnd := len(fs) - 1

	for {
		for fileEnd >= 0 && fs[fileEnd] == emptyID {
			fileEnd--
		}

		if fileEnd < 0 {
			break
		}

		fileID := fs[fileEnd]
		file := fileEnd - 1

		for file >= 0 && fs[file] == fileID {
			file--
		}

		fileLen := fileEnd - file

		slot := 0
		for slot < file {
			for slot < len(fs) && fs[slot] != emptyID {
				slot++
			}

			if slot >= file {
				break
			}

			slotEnd := slot + 1
			for slotEnd < len(fs) && fs[slotEnd] == emptyID {
				slotEnd++
			}

			slotLen := slotEnd - slot

			if slotLen >= fileLen {
				for range fileLen {
					fs[slot] = fileID
					fs[fileEnd] = emptyID

					slot++
					fileEnd--
				}

				break
			}

			slot = slotEnd
		}

		fileEnd = file
	}
}

func part02(fs filesystem) int {
	compactWithoutFragmentation(fs)
	return checksum(fs)
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	fs := parseInput(input)

	fmt.Println(part01(fs))
	fmt.Println(part02(fs))
}
