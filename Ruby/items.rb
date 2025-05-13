module Items
  def self.init_items_file
    return if File.exist?('data/items.txt')
    File.write('data/items.txt', "Voiture\nAssiette\nAppartement\n")
  end

  def self.list_items
    if File.exist?('data/items.txt') && !File.empty?('data/items.txt')
      puts File.read('data/items.txt').split("\n")
    else
      puts "Aucun objet disponible."
    end
  end

  def self.item_exists?(item)
    return false unless File.exist?('data/items.txt')
    File.read('data/items.txt').split("\n").include?(item)
  end
end