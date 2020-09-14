package problems.easy.problem476;

import utils.Helper;

public class Solution {
    public int findComplement(int num) {
        var mask = 1;
        var result = 0;

        if (num == 0 ) {
            result += ((num & 1) ^ 1) * mask;
            return result;
        }

        while (num != 0) {
            result += ((num & 1) ^ 1) * mask;
            num >>>= 1;
            mask <<= 1;
        }

        return result;
    }
}
