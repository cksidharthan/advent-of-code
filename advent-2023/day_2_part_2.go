package advent_2023

import (
	"github.com/cksidharthan/advent-of-code/elf"
	"log"
)

func Day02PartTwo(test bool) int {
	var file string
	var err error

	if test {
		file, err = elf.GetInputFile(true, 2023, 2, 2)
	} else {
		file, err = elf.GetInputFile(false, 2023, 2, 2)
	}

	if err != nil {
		log.Fatal(err)
	}

	contents, err := elf.GetContentsFromFile(file)
	if err != nil {
		log.Fatal(err)
	}

	var gameDataList []Games

	for _, line := range contents {
		gameData := getGameFromString(line)
		gameDataList = append(gameDataList, gameData)
	}

	var resultPower int
	for _, game := range gameDataList {
		gamePower := getMaxCubePowerInGame(game)
		resultPower += gamePower
	}
	return resultPower

}

func getMaxCubePowerInGame(game Games) int {
	var redMax, blueMax, greenMax int
	for _, cubes := range game.Cubes {
		if cubes.Red > redMax {
			redMax = cubes.Red
		}
		if cubes.Blue > blueMax {
			blueMax = cubes.Blue
		}
		if cubes.Green > greenMax {
			greenMax = cubes.Green
		}
	}

	return redMax * blueMax * greenMax
}
