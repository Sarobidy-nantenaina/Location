package rentals

import (
	"fmt"
	"os"
	"strings"
	"time"
	"rental_system/items"
)

// Rental représente une location
type Rental struct {
	Item      string
	StartDate time.Time
	EndDate   time.Time
}

// InitRentalsFile crée rentals.txt s'il n'existe pas
func InitRentalsFile() error {
	if _, err := os.Stat("data/rentals.txt"); os.IsNotExist(err) {
		_, err := os.Create("data/rentals.txt")
		if err != nil {
			return fmt.Errorf("création de rentals.txt : %v", err)
		}
	}
	return nil
}

// RentItem ajoute une location après validation
func RentItem(item, startDateStr, endDateStr string) error {
	// Vérifier si l'objet existe
	exists, err := items.ItemExists(item)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("l'objet %s n'existe pas", item)
	}

	// Parser les dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return fmt.Errorf("date de début invalide : %v", err)
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return fmt.Errorf("date de fin invalide : %v", err)
	}

	// Vérifier la durée minimale (1 jour)
	if endDate.Before(startDate.Add(24 * time.Hour)) {
		return fmt.Errorf("la location doit durer au moins un jour")
	}

	// Vérifier les conflits
	rentals, err := loadRentals()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, r := range rentals {
		if r.Item == item && r.EndDate.After(now) {
			if startDate.Before(r.EndDate) && endDate.After(r.StartDate) {
				return fmt.Errorf("%s est déjà loué jusqu'à %s", item, r.EndDate.Format("2006-01-02"))
			}
		}
	}

	// Ajouter la location
	file, err := os.OpenFile("data/rentals.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("ouverture de rentals.txt : %v", err)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s,%s,%s\n", item, startDateStr, endDateStr))
	if err != nil {
		return fmt.Errorf("écriture dans rentals.txt : %v", err)
	}
	return nil
}

// ListRentals affiche les locations actives d'un objet
func ListRentals(item string) error {
	exists, err := items.ItemExists(item)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("l'objet %s n'existe pas", item)
	}

	rentals, err := loadRentals()
	if err != nil {
		return err
	}

	now := time.Now()
	found := false
	for _, r := range rentals {
		if r.Item == item && r.EndDate.After(now) {
			fmt.Printf("Loué du %s au %s\n", r.StartDate.Format("2006-01-02"), r.EndDate.Format("2006-01-02"))
			found = true
		}
	}
	if !found {
		fmt.Printf("Aucune location active pour %s.\n", item)
	}
	return nil
}

// loadRentals charge les locations depuis rentals.txt
func loadRentals() ([]Rental, error) {
	data, err := os.ReadFile("data/rentals.txt")
	if err != nil {
		return nil, fmt.Errorf("lecture de rentals.txt : %v", err)
	}
	var rentals []Rental
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		startDate, err := time.Parse("2006-01-02", parts[1])
		if err != nil {
			continue
		}
		endDate, err := time.Parse("2006-01-02", parts[2])
		if err != nil {
			continue
		}
		rentals = append(rentals, Rental{
			Item:      parts[0],
			StartDate: startDate,
			EndDate:   endDate,
		})
	}
	return rentals, nil
}