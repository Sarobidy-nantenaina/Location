require 'date'

module Rentals
  def self.init_rentals_file
    File.write('data/rentals.txt', '') unless File.exist?('data/rentals.txt')
  end

  def self.rent_item(item, start_date_str, end_date_str)
    unless Items.item_exists?(item)
      return "Erreur : L'objet #{item} n'existe pas."
    end

    begin
      start_date = Date.parse(start_date_str)
      end_date = Date.parse(end_date_str)
    rescue ArgumentError
      return "Erreur : Format de date invalide (utilisez YYYY-MM-DD)."
    end

    if (end_date - start_date).to_i < 1
      return "Erreur : La location doit durer au moins un jour."
    end

    rentals = load_rentals
    now = Date.today
    rentals.each do |r|
      if r[:item] == item && r[:end_date] >= now
        if start_date <= r[:end_date] && end_date >= r[:start_date]
          return "Erreur : #{item} est déjà loué jusqu'à #{r[:end_date]}."
        end
      end
    end

    File.open('data/rentals.txt', 'a') do |f|
      f.puts "#{item},#{start_date_str},#{end_date_str}"
    end
    "Succès : #{item} loué du #{start_date_str} au #{end_date_str}."
  end

  def self.list_rentals(item)
    unless Items.item_exists?(item)
      puts "Erreur : L'objet #{item} n'existe pas."
      return
    end

    rentals = load_rentals
    now = Date.today
    found = false
    rentals.each do |r|
      if r[:item] == item && r[:end_date] >= now
        puts "Loué du #{r[:start_date]} au #{r[:end_date]}"
        found = true
      end
    end
    puts "Aucune location active pour #{item}." unless found
  end

  def self.load_rentals
    return [] unless File.exist?('data/rentals.txt')
    rentals = []
    File.readlines('data/rentals.txt').each do |line|
      item, start_date, end_date = line.chomp.split(',')
      begin
        rentals << {
          item: item,
          start_date: Date.parse(start_date),
          end_date: Date.parse(end_date)
        }
      rescue ArgumentError
        next
      end
    end
    rentals
  end
end