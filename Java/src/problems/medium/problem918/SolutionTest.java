package problems.medium.problem918;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void maxSubArraySumCircular1() {
        Solution solution = new Solution();

        assertEquals(3, solution.maxSubArraySumCircular(new int[] {1, -2, 3, -2}));
    }

    @Test
    void maxSubArraySumCircular2() {
        Solution solution = new Solution();

        assertEquals(10, solution.maxSubArraySumCircular(new int[] {5, -3, 5}));
    }

    @Test
    void maxSubArraySumCircular3() {
        Solution solution = new Solution();

        assertEquals(4, solution.maxSubArraySumCircular(new int[] {3, -1, 2, -1}));
    }

    @Test
    void maxSubArraySumCircular4() {
        Solution solution = new Solution();

        assertEquals(3, solution.maxSubArraySumCircular(new int[] {3, -2, 2, -3}));
    }
}