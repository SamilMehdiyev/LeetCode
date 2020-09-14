package problems.easy.problem1356;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void sortByBits() {
        Solution solution = new Solution();

        var result = solution.sortByBits(new int[]{0,1,2,3,4,5,6,7,8});

        assertArrayEquals(new int[]{0,1,2,4,8,3,5,6,7}, result);
    }
}