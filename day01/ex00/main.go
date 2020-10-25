package main

import (
	"errors"
	"flag"
	"fmt"
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

func main() {
	fileName := flag.String("f", "", "file with recipes to read")
	flag.Parse()

	fileExt, err := validateFileExt(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	recipes, err := recipesDB.NewRecipesDBByFileExt(fileExt)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = recipes.Read(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = recipes.Write()
	if err != nil {
		fmt.Println(err)
		return
	}
}
