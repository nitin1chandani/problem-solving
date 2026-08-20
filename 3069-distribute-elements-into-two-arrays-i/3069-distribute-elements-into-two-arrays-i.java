class Solution {
    public int[] resultArray(int[] nums) {
        int n = nums.length;

        List<Integer> arr1 = new ArrayList<>();
        List<Integer> arr2 = new ArrayList<>();
        
        arr1.add(nums[0]);
        arr2.add(nums[1]);

        for(int i = 2; i<n; i++){
            int arr1LastElement = arr1.get(arr1.size()-1);
            int arr2LastElement = arr2.get(arr2.size()-1);
            if(arr1LastElement>arr2LastElement){
                arr1.add(nums[i]);
            }else{
                arr2.add(nums[i]);
            }
        }

        // List<Integer> result = new ArrayList<>(arr1);
        int[] result = new ArrayList<>(arr1){{
            addAll(arr2);
        }}.stream().mapToInt(Integer::intValue).toArray();
        return result;
        
    }
}