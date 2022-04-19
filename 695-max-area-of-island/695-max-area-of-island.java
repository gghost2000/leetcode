class Solution {
        int n;
        int m;

    public int maxAreaOfIsland(int[][] grid) {
        int total = 0;
        n = grid.length;
        m = grid[0].length;

        for (int i =0; i<n; i++) {
            for(int j= 0; j<m; j++) {
                if (grid[i][j] !=0) {
                    total = Math.max(total, checkCount(i,j,grid));
                }
            }
        }

        return total;

    }

    private int checkCount(int i, int j, int[][] grid) {
        if (i < 0 || j <0 || i>=n || j >=m || grid[i][j] ==0) {
            return 0;
        }
        grid[i][j] = 0;

        return 1+ checkCount(i-1, j , grid) + checkCount(i+1, j, grid) + checkCount(i, j-1, grid) + checkCount(i, j+1, grid);
    }
}