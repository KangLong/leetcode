public class Solution {

    /**
     * 删除排序数组中的重复项
     * 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度
     * 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成
     * @param nums
     * @return
     */
    public static int removeDuplicates(int[] nums) {
        //1.有序数组 若出现第一个不相同的元素，则后续元素必然不同
        //2.len 作为“新数组”的索引，当出现不同元素时，直接将元素放入“新数组”的末尾
        int len = 0;
        for(int i = 0; i<nums.length;i++){
            if(nums[len]!=nums[i]){
                nums[++len] = nums[i];
            }
        }
        return len+1;
    }

    public static void main(String[] args) {
        int[] nums = {0,0,1,1,1,2,2,3,3,4};
        System.out.println(removeDuplicates(nums));
    }
}