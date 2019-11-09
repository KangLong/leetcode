import java.util.Arrays;

public class Solution {

    /**
     * 189. 旋转数组
     * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
     * @param nums
     * @param k
     */
    public static void rotate(int[] nums,int k){
        //1.旋转即将数组最后的元素移动到数组的最前面
        //2.移动 k 个元素，也就是将前面 length-k个元素移动到数组最后
        //3.创建一个嵌套循环，内循环将数组第一个元素移动至末尾，外循环为需要被移动的元素个数
        //4. 空间复杂度为 O(1) 时间复杂度为0(n^2)
        if(nums.length<2||k==0)
            return;
        if(k>nums.length){
            k = k%nums.length;
        }
        int tmp = 0;
        for (int i=0;i<nums.length-k;i++){
            for(int j=0;j<nums.length-1;j++){
                tmp = nums[j];
                nums[j] = nums[j+1];
                nums[j+1] = tmp;
            }
        }
    }

    public static void rotate2(int[] nums,int k){
        //1.旋转即将数组最后的元素移动到数组的最前面
        //2.复制一个新的数组，用两个并列循环分别处理 k 个元素的赋值问题
        if(nums.length<2||k==0)
            return;
        if(k>nums.length){
            k = k%nums.length;
        }
        int[] arrNums = Arrays.copyOf(nums,nums.length);
        int j= 0;
        for(int i=0;i<arrNums.length-k;i++){
            nums[k+i] = arrNums[j++];
        }
        for(int i=0;i<k;i++){
            nums[i] = arrNums[arrNums.length-k+i];
        }
    }

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
        int[] nums = {1,2,3,4,5,6,7};
        rotate2(nums,2);
        for(int n:nums){
            System.out.print(n);
        }
//        System.out.println(removeDuplicates(nums));
    }
}
