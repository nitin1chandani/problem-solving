class Solution {
    int[][] dp = new int[501][501];

    int solve(int l, int r, int[] prefixSum) {

        if (l >= r) {
            return 0;
        }

        if (dp[l][r] != -1) {
            return dp[l][r];
        }

        int score = 0;

        for (int mid = l; mid < r; mid++) {

            int leftSum = prefixSum[mid] - (l > 0 ? prefixSum[l - 1] : 0);
            int rightSum = prefixSum[r] - prefixSum[mid];

            if (leftSum < rightSum) {

                score = Math.max(
                    score,
                    leftSum + solve(l, mid, prefixSum)
                );

            } else if (leftSum > rightSum) {

                score = Math.max(
                    score,
                    rightSum + solve(mid + 1, r, prefixSum)
                );

            } else {

                score = Math.max(
                    score,
                    Math.max(
                        leftSum + solve(l, mid, prefixSum),
                        rightSum + solve(mid + 1, r, prefixSum)
                    )
                );
            }
        }

        return dp[l][r] = score;
    }

    public int stoneGameV(int[] stoneValue) {

        int n = stoneValue.length;

        for (int i = 0; i < 501; i++) {
            Arrays.fill(dp[i], -1);
        }

        int[] prefixSum = new int[n];

        prefixSum[0] = stoneValue[0];

        for (int i = 1; i < n; i++) {
            prefixSum[i] = prefixSum[i - 1] + stoneValue[i];
        }

        return solve(0, n - 1, prefixSum);
    }
}