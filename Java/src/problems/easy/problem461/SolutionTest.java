package problems.easy.problem461;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void hammingDistance() {
        Solution solution = new Solution();
        assertEquals(2, solution.hammingDistance(1, 4));
    }
}