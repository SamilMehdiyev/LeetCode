package problems.medium.problem918;

public class Solution {
    public int maxSubArraySumCircular(int[] A) {
        int sumMax = -30000;
        int sumMin = 30000;

        int currMax = 0, currMin = 0, sum = 0;

        for (int i = 0; i <A.length; i++) {
            currMax = Math.max(currMax + A[i], A[i]);
            sumMax = Math.max(currMax, sumMax);

            currMin = Math.min(currMin + A[i], A[i]);
            sumMin = Math.min(currMin, sumMin);

            sum += A[i];
        }

        if (sum != sumMin) {
            return Math.max(sum - sumMin, sumMax);
        }

        return sumMax;
    }
}
