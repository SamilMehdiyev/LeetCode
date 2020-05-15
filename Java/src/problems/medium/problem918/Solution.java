package problems.medium.problem918;

public class Solution {
    public int maxSubArraySumCircular(int[] A) {
        int currMax = Integer.MIN_VALUE;
        int totalMax = Integer.MIN_VALUE;
        int currMin = Integer.MAX_VALUE;
        int totalMin = Integer.MAX_VALUE;
        int sum = 0;

        for (int i = 0; i <A.length; i++) {
            currMax = Math.max(currMax, 0) + A[i];
            totalMax = Math.max(currMax, totalMax);

            currMin = Math.min(currMin, 0) + A[i];
            totalMin = Math.min(currMin, totalMin);

            sum += A[i];
        }

        if (sum != totalMin) {
            return Math.max(sum - totalMin, totalMax);
        }

        return totalMax;
    }
}
