package problems.easy.solution401;

import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<String> readBinaryWatch(int num) {
        var list = new ArrayList<String>();

        for (int hour = 1; hour < 12; hour++) {
            for (int minute = 1; minute < 60; minute++) {
                if (bitCounts(hour) + bitCounts(minute) == num) {
                    var str = hour + (minute < 10 ? ":0" : ":") + minute;
                    list.add(str);
                }
            }
        }

        return list;
    }

    private int bitCounts(int num) {
        var result = 0;

        while (num != 0) {
            result += num & 1;
            num >>>= 1;
        }

        return result;
    }
}