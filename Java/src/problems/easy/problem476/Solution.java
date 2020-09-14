package problems.easy.problem476;

import utils.Helper;

public class Solution {
    public int findComplement(int num) {
        var result = 0;
        var mask = 1;

        while (num != 0) {
            result += ((num & 1) ^ 1) * mask;
            num >>>= 1;
            mask <<= 1;
        }
        return result;
    }
}
