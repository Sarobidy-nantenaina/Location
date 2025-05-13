package main

import (
	"fmt"
	"time"
)

// Item représente un objet louable
type Item struct {
	Name         string
	IsReserved   bool
	ReservationEnd time.Time
}

// RentalSystem gère les objets louables
type RentalSystem struct {
	Items []Item
}

// NewRentalSystem crée un nouveau système de location
func NewRentalSystem() *RentalSystem {
	return &RentalSystem{
		Items: []Item{
			{Name: "Voiture"},
			{Name: "Assiette"},
			{Name: "Appartement"},
		},
	}
}

// Reserve tente de réserver un objet pour un nombre de jours
func (rs *RentalSystem) Reserve(itemName string, days int) error {
	if days < 1 {
		return fmt.Errorf("la durée de réservation doit être d'au moins 1 jour")
	}

	for i, item := range rs.Items {
		if item.Name == itemName {
			if item.IsReserved && time.Now().Before(item.ReservationEnd) {
				return fmt.Errorf("%s est déjà réservé jusqu'à %v", itemName, item.ReservationEnd)
			}
			rs.Items[i].IsReserved = true
			rs.Items[i].ReservationEnd = time.Now().Add(time.Duration(days) * 24 * time.Hour)
			return nil
		}
	}
	return fmt.Errorf("objet %s non trouvé", itemName)
}

// CheckReservations vérifie et met à jour le statut des réservations
func (rs *RentalSystem) CheckReservations() {
	for i, item := range rs.Items {
		if item.IsReserved && time.Now().After(item.ReservationEnd) {
			rs.Items[i].IsReserved = false
		}
	}
}

func main() {
	rs := NewRentalSystem()

	// Exemple d'utilisation
	fmt.Println("Tentative de réservation de la voiture pour 2 jours...")
	err := rs.Reserve("Voiture", 2)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Voiture réservée avec succès!")
	}

	// Tentative de réservation de la même voiture
	fmt.Println("\nTentative de réservation de la voiture à nouveau...")
	err = rs.Reserve("Voiture", 1)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Voiture réservée avec succès!")
	}

	// Simuler l'attente pour que la réservation expire (dans un vrai système, cela se ferait avec le temps réel)
	fmt.Println("\nAprès quelques jours (simulation)...")
	rs.CheckReservations()

	// Nouvelle tentative de réservation
	fmt.Println("Tentative de réservation de la voiture après expiration...")
	err = rs.Reserve("Voiture", 1)
	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Println("Voiture réservée avec succès!")
	}
}