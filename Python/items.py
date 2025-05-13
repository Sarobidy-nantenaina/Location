import os

class Items:
    @staticmethod
    def init_items_file():
        if not os.path.exists('data/items.txt'):
            with open('data/items.txt', 'w') as f:
                f.write("Voiture\nAssiette\nAppartement\n")

    @staticmethod
    def list_items():
        try:
            with open('data/items.txt', 'r') as f:
                items = [line.strip() for line in f if line.strip()]
            if items:
                for item in items:
                    print(item)
            else:
                print("Aucun objet disponible.")
        except FileNotFoundError:
            print("Erreur : items.txt introuvable.")

    @staticmethod
    def item_exists(item):
        try:
            with open('data/items.txt', 'r') as f:
                return item in [line.strip() for line in f]
        except FileNotFoundError:
            return False