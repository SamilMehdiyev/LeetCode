package problems.easy.problem762;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void countPrimeSetBits() {
        Solution solution = new Solution();
        assertEquals(16, solution.countPrimeSetBits(244, 269));
    }
}