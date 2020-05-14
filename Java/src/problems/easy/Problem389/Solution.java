package problems.easy.Problem389;

public class Solution {
    public char findTheDifference(String s, String t) {
        int[] characters = new int[26];

        for (int i = 0; i < s.length(); i++) {
            characters[t.charAt(i) - 'a'] -= 1;
            characters[s.charAt(i) - 'a'] += 1;
        }

        characters[t.charAt(t.length() - 1) - 'a'] -= 1;

        for (int i = 0; i < characters.length; i++) {
            if (characters[i] < 0) {
                return ((char)(i + 'a'));
            }
        }

        return 'a';
    }
}
