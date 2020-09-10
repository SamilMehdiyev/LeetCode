package problems.easy.problem1342;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void numberOfSteps() {
        problems.easy.problem1342.Solution solution = new Solution();
        assertEquals(12, solution.numberOfSteps(123));
    }
}