package recipes_db

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/LevOrlov5404/go-piscine/day01/models"
	"io/ioutil"
	"os"
)

type RecipesDB interface {
	Read(filename string) ([]*models.Cake, error)
	Write() error
}

type RecipesJSON struct {
	Cakes []*models.Cake `json:"cake"`
}

type RecipesXML struct {
	XMLName xml.Name       `xml:"recipes"`
	Cakes   []*models.Cake `xml:"cake"`
}

func (recipes *RecipesJSON) Read(filename string) ([]*models.Cake, error) {
	if filename == "" {
		return nil, errors.New("filename is not specified")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can not open file %s: %s", filename, err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can not read from file %s: %s", filename, err)
	}

	err = json.Unmarshal(fileBytes, recipes)
	if err != nil {
		return nil, fmt.Errorf("can not parse data from file %s: %s", filename, err)
	}

	return recipes.Cakes, nil
}

func (recipes *RecipesXML) Read(filename string) ([]*models.Cake, error) {
	if filename == "" {
		return nil, errors.New("filename is not specified")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can not open file %s: %s", filename, err)
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("can not read from file %s: %s", filename, err)
	}

	err = xml.Unmarshal(fileBytes, recipes)
	if err != nil {
		return nil, fmt.Errorf("can not unmarshal data from file %s: %s", filename, err)
	}

	return recipes.Cakes, nil
}

func (recipes *RecipesJSON) Write() error {
	prettyRecipes, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return errors.New("can not marshal data")
	}

	fmt.Println(string(prettyRecipes))
	return nil
}

func (recipes *RecipesXML) Write() error {
	prettyRecipes, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		return errors.New("can not marshal data")
	}

	fmt.Println(string(prettyRecipes))
	return nil
}

func NewRecipesDBByFileExt(fileExt string) (RecipesDB, error) {
	switch fileExt {
	case "json":
		return &RecipesJSON{}, nil
	case "xml":
		return &RecipesXML{}, nil
	default:
		return nil, errors.New("not supported ext: " + fileExt)
	}
}
