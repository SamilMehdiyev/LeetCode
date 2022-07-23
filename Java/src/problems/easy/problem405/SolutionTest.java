package problems.easy.problem405;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void toHex() {
        Solution solution = new Solution();
        assertEquals("ffffffff", solution.toHex(-1));
    }
}