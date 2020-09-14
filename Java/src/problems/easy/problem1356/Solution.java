package problems.easy.problem1356;

import java.util.*;

public class Solution {
    public int[] sortByBits(int[] arr) {
        var map = new HashMap<Integer, List<Integer>>();

        for (int i = 0; i < arr.length; i++) {
            var bits = bitCounts(arr[i]);
            if (map.containsKey(bits)) {
                map.get(bits).add(arr[i]);
                continue;
            }
            var list = new ArrayList<Integer>();
            list.add(arr[i]);
            map.put(bits, list);
        }

        var result = new int[arr.length];
        int idx = 0;
        for (List<Integer> list: map.values()){
            Collections.sort(list);
            for (int num: list) {
                result[idx] = num;
                idx++;
            }
        }

        return result;
    }

    private int bitCounts(int x) {
        int result = 0;

        while (x != 0) {
            result += x & 1;
            x >>>= 1;
        }

        return result;
    }
}
