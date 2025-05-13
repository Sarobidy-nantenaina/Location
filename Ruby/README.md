Rental System CLI (Ruby)
A simple Ruby CLI for managing rentals. Users can list items, rent them with start/end dates, view active rentals, and exit.
Features

List rentable items (e.g., Voiture, Assiette, Appartement).
Rent an item (minimum 1 day, checks for conflicts).
View active rentals for an item.
Exit the app.

Project Structure
<project_directory>/
├── main.rb
├── items.rb
├── rentals.rb
├── data/
│   ├── items.txt
│   └── rentals.txt

Setup

Install Ruby:

Download from https://rubyinstaller.org/ (e.g., Ruby 3.3.5 for Windows).
Install and add Ruby to PATH.
Verify: ruby -v (e.g., ruby 3.3.5).


Ensure Structure:

Place files as shown above.
Create data/ if missing: mkdir data.



Run
cd <project_directory>
ruby main.rb


Menu:
1: List items.
2: Rent item (enter name, dates as YYYY-MM-DD).
3: View active rentals.
4: Exit.



Notes

Data in data/items.txt and data/rentals.txt.
Dates: YYYY-MM-DD format.
Contact team for issues.

