from datetime import datetime
from items import Items
import os

class Rentals:
    @staticmethod
    def init_rentals_file():
        if not os.path.exists('data/rentals.txt'):
            open('data/rentals.txt', 'w').close()

    @staticmethod
    def rent_item(item, start_date_str, end_date_str):
        if not Items.item_exists(item):
            return f"Erreur : L'objet {item} n'existe pas."

        try:
            start_date = datetime.strptime(start_date_str, '%Y-%m-%d').date()
            end_date = datetime.strptime(end_date_str, '%Y-%m-%d').date()
        except ValueError:
            return "Erreur : Format de date invalide (utilisez YYYY-MM-DD)."

        if (end_date - start_date).days < 1:
            return "Erreur : La location doit durer au moins un jour."

        rentals = Rentals.load_rentals()
        now = datetime.now().date()
        for r in rentals:
            if r['item'] == item and r['end_date'] >= now:
                if start_date <= r['end_date'] and end_date >= r['start_date']:
                    return f"Erreur : {item} est déjà loué jusqu'à {r['end_date']}."

        with open('data/rentals.txt', 'a') as f:
            f.write(f"{item},{start_date_str},{end_date_str}\n")
        return f"Succès : {item} loué du {start_date_str} au {end_date_str}."

    @staticmethod
    def list_rentals(item):
        if not Items.item_exists(item):
            print(f"Erreur : L'objet {item} n'existe pas.")
            return

        rentals = Rentals.load_rentals()
        now = datetime.now().date()
        found = False
        for r in rentals:
            if r['item'] == item and r['end_date'] >= now:
                print(f"Loué du {r['start_date']} au {r['end_date']}")
                found = True
        if not found:
            print(f"Aucune location active pour {item}.")

    @staticmethod
    def load_rentals():
        rentals = []
        try:
            with open('data/rentals.txt', 'r') as f:
                for line in f:
                    parts = line.strip().split(',')
                    if len(parts) == 3:
                        try:
                            rentals.append({
                                'item': parts[0],
                                'start_date': datetime.strptime(parts[1], '%Y-%m-%d').date(),
                                'end_date': datetime.strptime(parts[2], '%Y-%m-%d').date()
                            })
                        except ValueError:
                            continue
        except FileNotFoundError:
            pass
        return rentals