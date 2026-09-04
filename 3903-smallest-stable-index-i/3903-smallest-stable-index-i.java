class Solution {
    public int firstStableIndex(int[] nums, int k) {
        // max(nums[0..i]) - min(nums[i..n - 1])
        // An index i is called stable if its instability score is less than or equal to k.
        // return smallest stable index

        // [5, 5, 5, 5]
        // pMax[i] - sMin[i]
        // [0, 0, 1, 4]

        //take care of edge case


        int n = nums.length;

        int[] pMax = new int[n];
        int[] sMin = new int[n];

        pMax[0] = nums[0];

        for(int i = 1; i<n; i++){
            pMax[i] = Math.max(pMax[i-1], nums[i]);
        }

        sMin[n-1] = nums[n-1];

        for(int i = n-2; i>=0; i--){
            sMin[i] = Math.min(sMin[i+1], nums[i]);
        }

        int smallestStableIndex = -1;
        for(int i = 0; i<n; i++){
            int currStabilityScore = pMax[i] - sMin[i];
            if(currStabilityScore<=k){
                smallestStableIndex = i;
                break;
            }
        }

        return smallestStableIndex;
    }
}