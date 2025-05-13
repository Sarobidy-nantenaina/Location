Rental System CLI (Go, Ruby, Python, Java)
A simple command-line interface (CLI) for managing rentals, implemented in Go, Ruby, Python, and Java. Users can list rentable items, rent items with start/end dates, view active rentals, and exit. Each language has a separate implementation with identical functionality.
Features

List rentable items (e.g., Voiture, Assiette, Appartement).
Rent an item (minimum 1 day, checks for scheduling conflicts).
View active rentals for an item.
Exit the application.

Project Structure
Each language has its own directory with the following structure:
Go
<project_directory>/
├── main.go
├── items/
│   └── items.go
├── rentals/
│   └── rentals.go
├── data/
│   ├── items.txt
│   └── rentals.txt
├── README.md

Ruby
<project_directory>/
├── main.rb
├── items.rb
├── rentals.rb
├── data/
│   ├── items.txt
│   └── rentals.txt
├── README.md

Python
<project_directory>/
├── main.py
├── items.py
├── rentals.py
├── data/
│   ├── items.txt
│   └── rentals.txt
├── README.md

Java
<project_directory>/
├── src/
│   ├── main/
│   │   ├── java/
│   │   │   ├── rental_system/
│   │   │   │   ├── Main.java
│   │   │   │   ├── Items.java
│   │   │   │   └── Rentals.java
├── data/
│   ├── items.txt
│   └── rentals.txt
├── README.md

Setup

Install Dependencies:

Go:
Download: https://go.dev/dl/ (e.g., Go 1.23.2 for Windows).
Verify: go version (e.g., go1.23.2).


Ruby:
Download: https://rubyinstaller.org/ (e.g., Ruby 3.3.5 for Windows).
Verify: ruby -v (e.g., ruby 3.3.5).


Python:
Download: https://www.python.org/ (e.g., Python 3.12 for Windows).
Verify: python --version (e.g., Python 3.12).


Java:
Download JDK: https://www.oracle.com/java/ (e.g., JDK 21 for Windows).
Verify: java -version (e.g., openjdk 21).




Ensure Structure:

Place files as shown for the desired language.
Create data/ if missing: mkdir data.



Run
Navigate to the project directory for the chosen language and run:

Go:
cd <project_directory>
go mod init rental_system
go run .


Ruby:
cd <project_directory>
ruby main.rb


Python:
cd <project_directory>
python main.py


Java:
cd <project_directory>
javac -d . src/main/java/rental_system/*.java
java rental_system.Main


Menu (same for all languages):

1: List items.
2: Rent item (enter name, dates as YYYY-MM-DD).
3: View active rentals.
4: Exit.



Notes

Data Storage: Items in data/items.txt, rentals in data/rentals.txt (format: item,start_date,end_date).
Date Format: YYYY-MM-DD.
Troubleshooting: Ensure language is installed, files are in place, and data/ exists. Contact team for issues.
Collaboration: Each language implementation is independent; choose one based on preference.

