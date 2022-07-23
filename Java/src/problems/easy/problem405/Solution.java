package problems.easy.problem405;

public class Solution {
    public String toHex(int num) {
        var hexStr = "";
        var hexArray = new String[] {"a", "b", "c", "d", "e", "f"};

        long x = num;
        if (num < 0) {
            x += (long) Math.pow(2, 32);
        }

        while (x != 0) {
            var reminder = x % 16;
            if (reminder >= 10) {
                hexStr = hexArray[(int) reminder - 10] + hexStr;
            } else {
                hexStr = reminder + hexStr;
            }
            x /= 16;
        }

        return hexStr;
    }
}
