package utils;

public class Helper {
    public static String convertIntToBinaryString(long number) {
        StringBuilder result = new StringBuilder();

        for (int i = 31; i >= 0; i--) {
            int mask = 1 << i;
            result.append((number & mask) != 0 ? "1" : "0");
        }

        return result.toString();
    }
}
