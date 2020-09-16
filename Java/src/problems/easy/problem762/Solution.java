package problems.easy.problem762;

public class Solution {
    public int countPrimeSetBits(int L, int R) {
        var result = 0;
        int[] primes = {0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1};

        for (int i = L; i <= R; i++) {
            var ones = getBitsCount(i);
            if (primes[ones] == 1) {
                result++;
            }
        }

        return result;
    }

    private int getBitsCount(int x) {
        var result = 0;

        while(x != 0) {
            result += x & 1;
            x >>>= 1;
        }

        return result;
    }
}
