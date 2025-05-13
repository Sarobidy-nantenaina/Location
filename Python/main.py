import os
from items import Items
from rentals import Rentals

# Initialiser les fichiers
os.makedirs('data', exist_ok=True)
Items.init_items_file()
Rentals.init_rentals_file()

while True:
    print("\n=== Système de Location ===")
    print("1. Voir tous les objets à louer")
    print("2. Louer un objet")
    print("3. Afficher les locations d'un objet")
    print("4. Quitter")
    choice = input("Choisissez une option (1-4) : ")

    if choice == '1':
        print("\nObjets disponibles :")
        Items.list_items()
    elif choice == '2':
        item = input("Nom de l'objet : ")
        start_date = input("Date de début (YYYY-MM-DD) : ")
        end_date = input("Date de fin (YYYY-MM-DD) : ")
        print(Rentals.rent_item(item, start_date, end_date))
    elif choice == '3':
        item = input("Nom de l'objet : ")
        print(f"\nLocations actives pour {item} :")
        Rentals.list_rentals(item)
    elif choice == '4':
        print("Au revoir !")
        break
    else:
        print("Option invalide, veuillez choisir entre 1 et 4.")