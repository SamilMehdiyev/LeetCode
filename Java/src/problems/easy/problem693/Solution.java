package problems.easy.problem693;

public class Solution {
    public boolean hasAlternatingBits(int n) {
        var rightBit = n & 1;
        while(n != 0) {
            n >>>= 1;
            if (rightBit != (n & 1)) {
                rightBit = n & 1;
                continue;
            }
            return false;
        }
        return true;
    }
}
