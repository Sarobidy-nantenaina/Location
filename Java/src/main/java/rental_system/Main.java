import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        // Initialiser les fichiers
        Items.initItemsFile();
        Rentals.initRentalsFile();

        Scanner scanner = new Scanner(System.in);
        while (true) {
            System.out.println("\n=== Système de Location ===");
            System.out.println("1. Voir tous les objets à louer");
            System.out.println("2. Louer un objet");
            System.out.println("3. Afficher les locations d'un objet");
            System.out.println("4. Quitter");
            System.out.print("Choisissez une option (1-4) : ");

            String choice = scanner.nextLine();
            switch (choice) {
                case "1":
                    System.out.println("\nObjets disponibles :");
                    Items.listItems();
                    break;
                case "2":
                    System.out.print("Nom de l'objet : ");
                    String item = scanner.nextLine();
                    System.out.print("Date de début (YYYY-MM-DD) : ");
                    String startDate = scanner.nextLine();
                    System.out.print("Date de fin (YYYY-MM-DD) : ");
                    String endDate = scanner.nextLine();
                    System.out.println(Rentals.rentItem(item, startDate, endDate));
                    break;
                case "3":
                    System.out.print("Nom de l'objet : ");
                    String itemForRentals = scanner.nextLine();
                    System.out.println("\nLocations actives pour " + itemForRentals + " :");
                    Rentals.listRentals(itemForRentals);
                    break;
                case "4":
                    System.out.println("Au revoir !");
                    scanner.close();
                    return;
                default:
                    System.out.println("Option invalide, veuillez choisir entre 1 et 4.");
            }
        }
    }
}