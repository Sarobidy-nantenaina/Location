Rental System CLI (Java)
A simple Java CLI for managing rentals. Users can list items, rent them with start/end dates, view active rentals, and exit.
Features

List rentable items (e.g., Voiture, Assiette, Appartement).
Rent an item (minimum 1 day, checks for conflicts).
View active rentals for an item.
Exit the app.

Project Structure
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

Install Java:

Download JDK from https://www.oracle.com/java/ (e.g., JDK 21 for Windows).
Install and add java to PATH.
Verify: java -version (e.g., openjdk 21).


Ensure Structure:

Place files as shown above.
Create data/ if missing: mkdir data.



Run
cd <project_directory>
javac -d . src/main/java/rental_system/*.java
java rental_system.Main


Menu:
1: List items.
2: Rent item (enter name, dates as YYYY-MM-DD).
3: View active rentals.
4: Exit.



Notes

Data in data/items.txt and data/rentals.txt.
Dates: YYYY-MM-DD format.
Contact team for issues.

