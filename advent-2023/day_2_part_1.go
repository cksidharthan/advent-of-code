package advent_2023

import (
	"github.com/cksidharthan/advent-of-code/elf"
	"log"
	"strings"
)

type GameCubes struct {
	Red   int
	Blue  int
	Green int
}

type Games struct {
	ID    int
	Cubes map[int]GameCubes
}

func Day02PartOne(test bool, red, blue, green int) int {
	var file string
	var err error

	if test {
		file, err = elf.GetInputFile(true, 2023, 2, 1)
	} else {
		file, err = elf.GetInputFile(false, 2023, 2, 1)
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

	// get the possible games
	var possibleGames []Games
	for _, gameData := range gameDataList {
		// check game for possibility
		if isValidGame(gameData, red, blue, green) {
			possibleGames = append(possibleGames, gameData)
		}
	}

	// sum the ids of the possible games
	var sum int
	for _, game := range possibleGames {
		sum += game.ID
	}

	return sum

}

func getGameFromString(line string) Games {
	var games Games

	// split the string by : to get the game id and the cubes separtely
	gameSplit := strings.Split(line, ":")

	// get the game id
	gameID := strings.Replace(gameSplit[0], "Game ", "", -1)
	games.ID = elf.ConvertStringToInt(gameID)
	games.Cubes = make(map[int]GameCubes)

	// get the cubes
	cubesStrings := strings.Split(gameSplit[1], ";")
	for index, cubeString := range cubesStrings {
		// split the cube string by , to get the cube colors
		cubeColors := strings.Split(cubeString, ",")
		var gameCube GameCubes
		for _, cubeColor := range cubeColors {
			// split the cube color by : to get the color and the number
			cubeColor = strings.Trim(cubeColor, " ")
			cubeColorSplit := strings.Split(cubeColor, " ")
			number := cubeColorSplit[0]
			color := cubeColorSplit[1]

			switch color {
			case "red":
				gameCube.Red = elf.ConvertStringToInt(number)
			case "blue":
				gameCube.Blue = elf.ConvertStringToInt(number)
			case "green":
				gameCube.Green = elf.ConvertStringToInt(number)
			}
		}
		games.Cubes[index] = gameCube
	}

	return games
}

func isValidGame(game Games, red, blue, green int) bool {
	for _, cube := range game.Cubes {
		if cube.Red <= red && cube.Blue <= blue && cube.Green <= green {
			continue
		}

		return false
	}
	return true
}
