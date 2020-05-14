package problems.easy.Problem389;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void findTheDifference() {
        Solution solution = new Solution();

        assertEquals('e', solution.findTheDifference("abcd", "abcde"));
    }
}