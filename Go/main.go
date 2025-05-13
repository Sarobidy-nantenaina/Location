package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"rental_system/items"
	"rental_system/rentals"
)

func main() {
	// Initialiser les fichiers de données
	if err := items.InitItemsFile(); err != nil {
		fmt.Println("Erreur d'initialisation:", err)
		return
	}
	if err := rentals.InitRentalsFile(); err != nil {
		fmt.Println("Erreur d'initialisation:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== Système de Location ===")
		fmt.Println("1. Voir tous les objets à louer")
		fmt.Println("2. Louer un objet")
		fmt.Println("3. Afficher les locations d'un objet")
		fmt.Println("4. Quitter")
		fmt.Print("Choisissez une option (1-4) : ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("\nObjets disponibles :")
			if err := items.ListItems(); err != nil {
				fmt.Println("Erreur:", err)
			}
		case "2":
			fmt.Print("Nom de l'objet : ")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())
			fmt.Print("Date de début (YYYY-MM-DD) : ")
			scanner.Scan()
			startDate := strings.TrimSpace(scanner.Text())
			fmt.Print("Date de fin (YYYY-MM-DD) : ")
			scanner.Scan()
			endDate := strings.TrimSpace(scanner.Text())
			if err := rentals.RentItem(item, startDate, endDate); err != nil {
				fmt.Println("Erreur:", err)
			} else {
				fmt.Printf("Succès : %s loué du %s au %s.\n", item, startDate, endDate)
			}
		case "3":
			fmt.Print("Nom de l'objet : ")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())
			fmt.Printf("\nLocations actives pour %s :\n", item)
			if err := rentals.ListRentals(item); err != nil {
				fmt.Println("Erreur:", err)
			}
		case "4":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide, veuillez choisir entre 1 et 4.")
		}
	}
}