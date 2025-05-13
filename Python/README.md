Rental System CLI (Python)
A simple Python CLI for managing rentals. Users can list items, rent them with start/end dates, view active rentals, and exit.
Features

List rentable items (e.g., Voiture, Assiette, Appartement).
Rent an item (minimum 1 day, checks for conflicts).
View active rentals for an item.
Exit the app.

Project Structure
<project_directory>/
├── main.py
├── items.py
├── rentals.py
├── data/
│   ├── items.txt
│   └── rentals.txt

Setup

Install Python:

Download from https://www.python.org/ (e.g., Python 3.12 for Windows).
Install and add Python to PATH.
Verify: python --version (e.g., Python 3.12).


Ensure Structure:

Place files as shown above.
Create data/ if missing: mkdir data.



Run
cd <project_directory>
python main.py


Menu:
1: List items.
2: Rent item (enter name, dates as YYYY-MM-DD).
3: View active rentals.
4: Exit.



Notes

Data in data/items.txt and data/rentals.txt.
Dates: YYYY-MM-DD format.
Contact team for issues.

