package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/LevOrlov5404/go-piscine/day01/models"
	recipesDB "github.com/LevOrlov5404/go-piscine/day01/recipes-db"
	"strings"
)

func validateFileExt(fileName string) (ext string, err error) {
	if fileName == "" {
		return "", errors.New("file is not specified")
	}

	fileNameParts := strings.Split(fileName, ".")
	if len(fileNameParts) != 2 || (fileNameParts[1] != "json" && fileNameParts[1] != "xml") {
		return "", errors.New("not valid file name")
	}

	return fileNameParts[1], nil
}

func getCakesFromFile(fileName string) ([]*models.Cake, error) {
	fileExt, err := validateFileExt(fileName)
	if err != nil {
		return nil, err
	}

	recipes, err := recipesDB.NewRecipesDBByFileExt(fileExt)
	if err != nil {
		return nil, err
	}

	cakes, err := recipes.Read(fileName)
	if err != nil {
		return nil, err
	}

	return cakes, nil
}

func compareCakesNames(old []*models.Cake, new []*models.Cake) (commonNames []string, oldCakesMap map[string]*models.Cake, newCakesMap map[string]*models.Cake) {
	oldCakesMap = make(map[string]*models.Cake)
	newCakesMap = make(map[string]*models.Cake)

	for i := range new {
		newCakesMap[new[i].Name] = new[i]
	}

	for i := range old {
		oldCakesMap[old[i].Name] = old[i]

		if newCakesMap[old[i].Name] == nil {
			fmt.Printf("REMOVED cake \"%s\"\n", old[i].Name)
		} else {
			commonNames = append(commonNames, old[i].Name)
		}
	}

	for i := range new {
		if oldCakesMap[new[i].Name] == nil {
			fmt.Printf("ADDED cake \"%s\"\n", new[i].Name)
		}
	}

	return
}

func compareCakes(old []*models.Cake, new []*models.Cake) {
	//commonNames, oldCakesMap, newCakesMap := compareCakesNames(old, new)
	compareCakesNames(old, new)
}

func main() {
	oldFileName := flag.String("old", "", "old file with recipes")
	newFileName := flag.String("new", "", "new file with recipes")
	flag.Parse()

	oldCakes, err := getCakesFromFile(*oldFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	newCakes, err := getCakesFromFile(*newFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	compareCakes(oldCakes, newCakes)
}
