import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Solution {


    public static void main(String[] args) {
//        int[] nums = {0,1,0,2,1,0,1,3,2,1,2,1};
//        System.out.println(trap(nums));
        String s = "rat", t = "art";
        System.out.println(isAnagram(s,t));
    }

    /**
     * 242. 有效的字母异位词
     * @param s
     * @param t
     * @return
     */
    public static boolean isAnagram(String s, String t) {
        //1.将两个字符串直接转换成 ascii数组，并排序，若相等则认为异位词
        char[] c1 = s.toCharArray();
        char[] c2 = t.toCharArray();
        Arrays.sort(c1);
        Arrays.sort(c2);
        s = new String(c1);
        t = new String(c2);
        return s.equals(t);
    }

    /**
     * 42. 接雨水
     * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
     * @param height
     * @return
     */
    public static int trap(int[] height){
        //双指针方法，从数组前后同时遍历
        //能装对少水只和木桶的短板相关，故找到短板，再减去底部凸起部分，就能知道水桶能装多少水
        int left = 0, right = height.length - 1; //表示数组左右索引值
        int leftMax =0, rightMax = 0; //记录左右两边的最大值
        int ans = 0; //结果
        while(left < right){
            if(height[left]<height[right]){ //当左指针元素小右指针元素时
                if(height[left]>=leftMax){ //当左侧元素大于记录值时， 刷新最大值
                    leftMax = height[left];
                }else{
                    //当前位置所能装载的水量, 是由左侧最高点和当前元素的差值所决定
                    ans += (leftMax-height[left]);
                }
                left++;
            }else{
                if(height[right]>=rightMax){
                    rightMax = height[right];
                }else{
                    //当前位置所能装载的水量, 是由右侧最高点和当前元素的差值所决定
                    ans += (rightMax - height[right]);
                }
                right--;
            }

        }
        return ans;
    }


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

}
