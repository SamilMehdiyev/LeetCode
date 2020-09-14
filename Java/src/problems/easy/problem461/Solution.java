package problems.easy.problem461;

class Solution {
    public int hammingDistance(int x, int y) {
        var result = 0;

        var z = x ^ y;

        while (z != 0) {
            if ((z & 1) == 1) {
                result++;
            }
            z >>>= 1;
        }

        return result;
    }
}