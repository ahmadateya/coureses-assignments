package regex_assignment;

import java.util.Scanner;
import java.util.regex.Pattern;
import java.util.regex.Matcher;

/**
 * Simple utility class that compiles a given regex pattern, and try to extract
 * the matched pattern from the given text
 */
public class RegexCompiler {
    public static void main(String[] args) {
        Scanner console = new Scanner(System.in);
        while (true) {
            // System.out.println("Enter your regex...");
            // String regex = console.nextLine();

            /*==============================================================
             * Assignment RegEx
             * ===========================================================
             * 
             * Email Regex = ^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$
             * Fax Regex = ^(\+?\d{1,}(\s?|\-?)\d*(\s?|\-?)\(?\d{2,}\)?(\s?|\-?)\d{3,}\s?\d{3,})$
             * IP Regex = \b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b
             * Date Pattern RegEx = ^(0[1-9]|[12][0-9]|3[01])[-/.](0[1-9]|1[012])[-/.](19|20)\d\d$
             * <title></title> RegEx = ^(<title>).*(</title>)$
             * HTML Tags (any tag) RegEx = <.+?>
             * 
             * ==============================================================
             */
            String regex = "[0-2][0-4]:[0-5][0-9]:[0-5][0-9]";
            Pattern pattern = Pattern.compile(regex);
            System.out.println("Enter your input text...");
            String text = console.nextLine();
            Matcher matcher = pattern.matcher(text);
            boolean found = false;
            System.out.println("==========================================");
            while (matcher.find()) {
                System.out.format("I found the text \"%s\" starting " + "at index %d and ending at " + "index %d.%n",
                        matcher.group(), matcher.start(), matcher.end());
                found = true;
            }
            if (!found)
                System.out.format("No match found.%n");
            System.out.println("==========================================");
        }
    }
}
