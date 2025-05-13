Rental System CLI (Go)
A simple Go CLI for managing rentals. Users can list items, rent them with start/end dates, view active rentals, and exit.
Features

List rentable items (e.g., Voiture, Assiette, Appartement).
Rent an item (minimum 1 day, checks for conflicts).
View active rentals for an item.
Exit the app.

Project Structure
<project_directory>/
├── main.go
├── items/
│   └── items.go
├── rentals/
│   └── rentals.go
├── data/
│   ├── items.txt
│   └── rentals.txt
├── go.mod

Setup

Install Go:

Download from https://go.dev/dl/ (e.g., Windows: go1.23.2.windows-amd64.msi).
Install and add Go to PATH.
Verify: go version (e.g., go1.23.2).


Ensure Structure:

Place files as shown above.
Create data/ if missing: mkdir data.


Initialize Module:
cd <project_directory>
go mod init rental_system



Run
cd <project_directory>
go run .


Menu:
1: List items.
2: Rent item (enter name, dates as YYYY-MM-DD).
3: View active rentals.
4: Exit.



Notes

Data in data/items.txt and data/rentals.txt.
Dates: YYYY-MM-DD format.
Contact team for issues.

