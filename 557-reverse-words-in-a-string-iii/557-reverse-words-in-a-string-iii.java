class Solution {
    public String reverseWords(String s) {
        StringBuilder result = new StringBuilder();

        String[] wordsList = s.split(" ");

        for (String word : wordsList) {
               result.append(revertWord(word)).append(" ");
        }

        return result.toString().trim();
    }
        private String revertWord(String word) {
        char[] chars = word.toCharArray();

        int left = 0;
        int right = chars.length-1;

        while (left<= right) {
            char temp = chars[left];
            chars[left] = chars[right];
            chars[right] = temp;

            left++;
            right--;
        }

        return String.valueOf(chars);
    }
}