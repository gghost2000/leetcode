class Solution {
    public int lengthOfLongestSubstring(String s) {
        int longestLength = 0;
        int position = 0;

        Set<Character> characterHashSet = new HashSet<>();

        while (position + longestLength <= s.length()-1) {
            String subString = s.substring(position, position + characterHashSet.size() + 1);
            Character firstCharacter = subString.charAt(0);
            Character lastCharacter = subString.charAt(subString.length()-1);

            if (isNonRepeatingChars(characterHashSet, lastCharacter)) {
                characterHashSet.add(lastCharacter);
                longestLength = Math.max(longestLength, subString.length());
            } else {
                characterHashSet.remove(firstCharacter);
                position++;
            }
        }

        return longestLength;
    }
    
    
    private boolean isNonRepeatingChars(Set<Character> characterHashSet, Character lastCharacter) {
        return !characterHashSet.contains(lastCharacter);
    }

}