package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"strings"

	"github.com/kevin-kho/aoc-utilities/common"
)

type Food struct {
	Ingredients map[string]bool
	Allergens   map[string]bool
}

func GetFoods(data []byte) []Food {
	var res []Food

	for entry := range bytes.Lines(data) {
		entry = bytes.TrimSpace(entry)

		ingred, aller, _ := strings.Cut(string(entry), " (contains ")
		ingredArr := strings.Split(ingred, " ")
		aller = strings.TrimSuffix(aller, ")")
		allerArr := strings.Split(aller, ", ")

		for i, val := range ingredArr {
			val = strings.TrimSpace(val)
			ingredArr[i] = val
		}

		for i, val := range allerArr {
			val = strings.TrimSpace(val)
			allerArr[i] = val
		}

		ingredients := make(map[string]bool)
		for _, val := range ingredArr {
			ingredients[val] = true
		}
		allergens := make(map[string]bool)
		for _, val := range allerArr {
			allergens[val] = true
		}

		res = append(res, Food{
			Ingredients: ingredients,
			Allergens:   allergens,
		})

	}

	return res
}

func SolvePartOne(foods []Food) (int, map[string]int) {

	allergenSet := make(map[string]bool)
	for _, f := range foods {
		maps.Copy(allergenSet, f.Allergens)
	}

	ingredientCt := make(map[string]int)
	for _, food := range foods {
		for ingred := range food.Ingredients {
			ingredientCt[ingred]++
		}
	}

	for allergen := range allergenSet {

		var ingredients map[string]bool // Ingredients that contain the allergen
		for _, f := range foods {
			if f.Allergens[allergen] {
				if ingredients == nil {
					ingredients = f.Ingredients
					continue
				}

				// Find intersection of allergens
				newIngred := make(map[string]bool)
				for i := range ingredients {
					if f.Ingredients[i] {
						newIngred[i] = true
					}
				}
				ingredients = newIngred
			}

		}

		// Remove ingredients that contain the allergen
		for i := range ingredients {
			delete(ingredientCt, i)
		}

	}

	var count int
	for _, val := range ingredientCt {
		count += val
	}

	return count, ingredientCt

}

func SolvePartTwo(foods []Food, inertIngredients map[string]int) {

	allergenSet := make(map[string]bool)
	for _, f := range foods {
		maps.Copy(allergenSet, f.Allergens)
	}

	activeIngredients := make(map[string]bool)
	for _, f := range foods {
		maps.Copy(activeIngredients, f.Ingredients)
	}
	for ingre := range inertIngredients {
		delete(activeIngredients, ingre)
	}

	mp := make(map[string]map[string]int)
	for ingre := range activeIngredients {
		mp[ingre] = map[string]int{}
	}

	for _, f := range foods {
		if len(f.Allergens) == 0 {
			continue
		}

		for ingre := range f.Ingredients {
			if _, ok := inertIngredients[ingre]; ok {
				continue
			}

			for a := range f.Allergens {
				mp[ingre][a]++
			}
		}
	}

	fmt.Println(mp)

}

func main() {
	// data, err := common.ReadInput("inputExample.txt")
	data, err := common.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data = common.TrimNewLineSuffix(data)

	foods := GetFoods(data)

	res, inertIngredients := SolvePartOne(foods)
	fmt.Println(res)

	SolvePartTwo(foods, inertIngredients)

}
