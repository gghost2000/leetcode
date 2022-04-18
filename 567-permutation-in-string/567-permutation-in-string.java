class Solution {
    public boolean checkInclusion(String s1, String s2) {
        Map<Character,Integer> characterIntegerMap = resetHashSetByString(s1);

        int startPoint = 0;
        int endPoint = 0;

        while (endPoint < s2.length()) {
            char endChar = s2.charAt(endPoint);
            char startChar = s2.charAt(startPoint);

            if (characterIntegerMap.containsKey(endChar) && characterIntegerMap.get(endChar) > 0) {
                characterIntegerMap.put(endChar, characterIntegerMap.get(endChar)-1);
                endPoint++;
            } else {
                if (characterIntegerMap.containsKey(startChar)) {
                    characterIntegerMap.put(startChar, characterIntegerMap.get(startChar)+1);
                }
                startPoint++;
            }
            endPoint = Math.max(startPoint, endPoint);
            if (isAllZero(characterIntegerMap)) {
                return true;
            }
        }


        return false;
    }
    private boolean isAllZero(Map<Character,Integer> characterIntegerMap) {
        for (Character character : characterIntegerMap.keySet()) {
            if (characterIntegerMap.get(character) !=0) {
                return false;
            }
        }
        return true;
    }

    private Map<Character,Integer> resetHashSetByString(String text) {
        Map<Character,Integer> characterIntegerMap = new HashMap<>();

        for (char c : text.toCharArray()) {
            if (characterIntegerMap.containsKey(c)) {
                characterIntegerMap.put(c, characterIntegerMap.get(c)+1);
            } else {
                characterIntegerMap.put(c,1);
            }
        }
        return characterIntegerMap;
    }
}