package problems.easy.problem1486;

class Solution {
    public int xorOperation(int n, int start) {
        int result = start;

        for (int i = 1; i < n; i++) {
            var num = start + 2 * i;
            result ^= num;
        }
        return result;
    }
}