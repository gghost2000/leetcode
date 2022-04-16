class Solution {
    public void rotate(int[] nums, int k) {
               int[] tempArr = new int[nums.length];

       for (int i=0; i<nums.length; i++) {
           tempArr[i] = nums[i];
       }

        for (int i=0; i<nums.length; i++) {
            nums[(i+k)%(nums.length)] = tempArr[i];
        }
    }
}