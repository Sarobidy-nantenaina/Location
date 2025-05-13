import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Items {
    public static void initItemsFile() {
        File file = new File("data/items.txt");
        if (!file.exists()) {
            try {
                Files.createDirectories(Paths.get("data"));
                try (PrintWriter writer = new PrintWriter(file)) {
                    writer.println("Voiture");
                    writer.println("Assiette");
                    writer.println("Appartement");
                }
            } catch (IOException e) {
                System.out.println("Erreur lors de l'initialisation de items.txt : " + e.getMessage());
            }
        }
    }

    public static void listItems() {
        try (BufferedReader reader = new BufferedReader(new FileReader("data/items.txt"))) {
            String line;
            boolean hasItems = false;
            while ((line = reader.readLine()) != null) {
                if (!line.trim().isEmpty()) {
                    System.out.println(line);
                    hasItems = true;
                }
            }
            if (!hasItems) {
                System.out.println("Aucun objet disponible.");
            }
        } catch (IOException e) {
            System.out.println("Erreur lors de la lecture de items.txt : " + e.getMessage());
        }
    }

    public static boolean itemExists(String item) {
        try (BufferedReader reader = new BufferedReader(new FileReader("data/items.txt"))) {
            String line;
            while ((line = reader.readLine()) != null) {
                if (line.trim().equals(item)) {
                    return true;
                }
            }
        } catch (IOException e) {
            System.out.println("Erreur lors de la vérification de l'objet : " + e.getMessage());
        }
        return false;
    }
}