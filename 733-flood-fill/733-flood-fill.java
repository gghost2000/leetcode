class Solution {
    public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
        Queue<PointInfo> pointInfos = new LinkedList<>();
        boolean[][] checkPoint = new boolean[image.length][image[0].length];
        int targetColor = image[sr][sc];
        image[sr][sc]  = newColor;
        checkPoint[sr][sc] = true;

        pointInfos.add(new PointInfo(sr,sc));

        while (!pointInfos.isEmpty()) {
            PointInfo pointInfo = pointInfos.poll();

            checkPoint[pointInfo.x][pointInfo.y] = true;

            List<PointInfo> filledList = fillColor(image, pointInfo.x, pointInfo.y, targetColor, newColor, checkPoint);
            pointInfos.addAll(filledList);
        }

        return image;
    }
    private List<PointInfo> fillColor(int[][] image, int sr, int sc, int targetColor, int newColor, boolean[][] checkPoint) {
        List<PointInfo> infos = new ArrayList<>();

        if (sr-1 >=0 &&image[sr-1][sc] == targetColor && !checkPoint[sr-1][sc]) {
            infos.add(new PointInfo(sr-1,sc));
            image[sr-1][sc] = newColor;
        }

        if (sr+1 < image.length && image[sr+1][sc] == targetColor&& !checkPoint[sr+1][sc]) {
            infos.add(new PointInfo(sr+1,sc));
            image[sr+1][sc] = newColor;
        }

        if (sc-1 >= 0 && image[sr][sc-1] == targetColor&& !checkPoint[sr][sc-1]) {
            infos.add(new PointInfo(sr,sc-1));
            image[sr][sc-1] = newColor;
        }

        if (sc+1 < image[0].length && image[sr][sc+1] == targetColor&& !checkPoint[sr][sc+1]) {
            infos.add(new PointInfo(sr,sc+1));
            image[sr][sc+1] = newColor;
        }
        return infos;
    }    
    
}


class PointInfo{
    int x;
    int y;
    public PointInfo(int x, int y) {
        this.x = x;
        this.y = y;
    }
}