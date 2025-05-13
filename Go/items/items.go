package items

import (
	"fmt"
	"os"
	"strings"
)

// InitItemsFile crée items.txt avec des objets par défaut s'il n'existe pas
func InitItemsFile() error {
	if _, err := os.Stat("../data/items.txt"); os.IsNotExist(err) {
		defaultItems := []string{"Voiture", "Assiette", "Appartement"}
		file, err := os.Create("../data/items.txt")
		if err != nil {
			return fmt.Errorf("création de items.txt : %v", err)
		}
		defer file.Close()
		for _, item := range defaultItems {
			if _, err := file.WriteString(item + "\n"); err != nil {
				return fmt.Errorf("écriture dans items.txt : %v", err)
			}
		}
	}
	return nil
}

// ListItems affiche tous les objets
func ListItems() error {
	data, err := os.ReadFile("../data/items.txt")
	if err != nil {
		return fmt.Errorf("lecture de items.txt : %v", err)
	}
	items := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(items) == 0 || items[0] == "" {
		fmt.Println("Aucun objet disponible.")
		return nil
	}
	for _, item := range items {
		fmt.Println(item)
	}
	return nil
}

// ItemExists vérifie si un objet existe
func ItemExists(item string) (bool, error) {
	data, err := os.ReadFile("../data/items.txt")
	if err != nil {
		return false, fmt.Errorf("lecture de items.txt : %v", err)
	}
	items := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, i := range items {
		if i == item {
			return true, nil
		}
	}
	return false, nil
}