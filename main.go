package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pokemon struct {
	Id             int
	Name           string
	Species_id     int
	Height         int
	BaseExperience int
}

func main() {
	file, err := os.Open("./pokemon.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pokemon := []Pokemon{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		id, err := strconv.Atoi(s[0])
		speciesId, err := strconv.Atoi(s[2])
		height, err := strconv.Atoi(s[3])
		baseExperiment, err := strconv.Atoi(s[4])

		if err != nil {
			log.Fatal(err)
		}

		pokemon = append(pokemon, Pokemon{id, s[1], speciesId, height, baseExperiment})
	}

	fmt.Println(pokemon)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
