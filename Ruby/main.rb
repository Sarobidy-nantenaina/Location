require_relative 'items'
require_relative 'rentals'
require 'fileutils'

# Initialiser les fichiers de données
FileUtils.mkdir_p('data')
Items.init_items_file
Rentals.init_rentals_file

loop do
  puts "\n=== Système de Location ==="
  puts "1. Voir tous les objets à louer"
  puts "2. Louer un objet"
  puts "3. Afficher les locations d'un objet"
  puts "4. Quitter"
  print "Choisissez une option (1-4) : "

  choice = gets.chomp
  case choice
  when '1'
    puts "\nObjets disponibles :"
    Items.list_items
  when '2'
    print "Nom de l'objet : "
    item = gets.chomp
    print "Date de début (YYYY-MM-DD) : "
    start_date = gets.chomp
    print "Date de fin (YYYY-MM-DD) : "
    end_date = gets.chomp
    result = Rentals.rent_item(item, start_date, end_date)
    puts result
  when '3'
    print "Nom de l'objet : "
    item = gets.chomp
    puts "\nLocations actives pour #{item} :"
    Rentals.list_rentals(item)
  when '4'
    puts "Au revoir !"
    exit
  else
    puts "Option invalide, veuillez choisir entre 1 et 4."
  end
end