class Solution {
    public int[] sortByBits(int[] arr) {
  List<Integer> result = new ArrayList<>();

        Map<Integer, List<Integer>> integerMap = new HashMap<>();
        int maxBits = 0;

        for (int i : arr) {
            int countBits = getBits(i);
            List<Integer> integerList = integerMap.get(countBits) == null ? new ArrayList<>() : integerMap.get(countBits);
            integerList.add(i);


            integerMap.put(countBits, integerList);
            maxBits = Math.max(countBits, maxBits);
        }


        for (int i = 0; i <= maxBits; i++) {
            if (integerMap.get(i) != null) {
                Collections.sort(integerMap.get(i));
                result.addAll(integerMap.get(i));
            }
        }


        return result.stream().mapToInt(Integer::intValue).toArray();
    }
    
        private int getBits(int a) {
        int count =0;

        while (a > 0) {
            if (a%2==1) {
                count++;
            }
            a/=2;
        }

        return count;
    }
}