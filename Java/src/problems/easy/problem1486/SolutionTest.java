package problems.easy.problem1486;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void xorOperation() {
        Solution solution = new Solution();
        assertEquals(8, solution.xorOperation(5, 0));
    }
}