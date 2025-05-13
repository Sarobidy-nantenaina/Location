import java.io.*;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.LocalDate;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

public class Rentals {
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd");

    public static void initRentalsFile() {
        File file = new File("data/rentals.txt");
        if (!file.exists()) {
            try {
                Files.createDirectories(Paths.get("data"));
                file.createNewFile();
            } catch (IOException e) {
                System.out.println("Erreur lors de l'initialisation de rentals.txt : " + e.getMessage());
            }
        }
    }

    public static String rentItem(String item, String startDateStr, String endDateStr) {
        if (!Items.itemExists(item)) {
            return "Erreur : L'objet " + item + " n'existe pas.";
        }

        LocalDate startDate, endDate;
        try {
            startDate = LocalDate.parse(startDateStr, DATE_FORMATTER);
            endDate = LocalDate.parse(endDateStr, DATE_FORMATTER);
        } catch (DateTimeParseException e) {
            return "Erreur : Format de date invalide (utilisez YYYY-MM-DD).";
        }

        if (endDate.isBefore(startDate.plusDays(1))) {
            return "Erreur : La location doit durer au moins un jour.";
        }

        try {
            BufferedReader reader = new BufferedReader(new FileReader("data/rentals.txt"));
            String line;
            LocalDate now = LocalDate.now();
            while ((line = reader.readLine()) != null) {
                String[] parts = line.split(",");
                if (parts.length == 3 && parts[0].equals(item)) {
                    try {
                        LocalDate existingEnd = LocalDate.parse(parts[2], DATE_FORMATTER);
                        if (existingEnd.isAfter(now) || existingEnd.isEqual(now)) {
                            LocalDate existingStart = LocalDate.parse(parts[1], DATE_FORMATTER);
                            if (startDate.isBefore(existingEnd.plusDays(1)) && endDate.isAfter(existingStart.minusDays(1))) {
                                reader.close();
                                return "Erreur : " + item + " est déjà loué jusqu'à " + existingEnd + ".";
                            }
                        }
                    } catch (DateTimeParseException ignored) {
                    }
                }
            }
            reader.close();

            try (PrintWriter writer = new PrintWriter(new FileWriter("data/rentals.txt", true))) {
                writer.println(item + "," + startDateStr + "," + endDateStr);
            }
            return "Succès : " + item + " loué du " + startDateStr + " au " + endDateStr + ".";
        } catch (IOException e) {
            return "Erreur lors de l'ajout de la location : " + e.getMessage();
        }
    }

    public static void listRentals(String item) {
        if (!Items.itemExists(item)) {
            System.out.println("Erreur : L'objet " + item + " n'existe pas.");
            return;
        }

        boolean found = false;
        try (BufferedReader reader = new BufferedReader(new FileReader("data/rentals.txt"))) {
            String line;
            LocalDate now = LocalDate.now();
            while ((line = reader.readLine()) != null) {
                String[] parts = line.split(",");
                if (parts.length == 3 && parts[0].equals(item)) {
                    try {
                        LocalDate endDate = LocalDate.parse(parts[2], DATE_FORMATTER);
                        if (endDate.isAfter(now) || endDate.isEqual(now)) {
                            System.out.println("Loué du " + parts[1] + " au " + parts[2]);
                            found = true;
                        }
                    } catch (DateTimeParseException ignored) {
                    }
                }
            }
        } catch (IOException e) {
            System.out.println("Erreur lors de la lecture des locations : " + e.getMessage());
        }
        if (!found) {
            System.out.println("Aucune location active pour " + item + ".");
        }
    }
}