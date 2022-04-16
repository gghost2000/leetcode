class Solution {
    public int[] sortedSquares(int[] nums) {
                int[] results = new int[nums.length];

        int leftIndex = 0;
        int rightIndex = nums.length-1;

        for (int i =nums.length-1; i >= 0; i--) {
            int leftValue = Math.abs(nums[leftIndex]);
            int rightValue = Math.abs(nums[rightIndex]);

            if (leftValue >= rightValue) {
                results[i] = leftValue * leftValue;
                leftIndex++;
            } else {
                results[i] = rightValue * rightValue;
                rightIndex--;
            }
        }

        return results;
    }
}