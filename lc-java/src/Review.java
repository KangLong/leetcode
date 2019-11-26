import java.util.Arrays;

public class Review {

    public static void main(String[] args) {
        ListNode l1 = new ListNode(1);
        ListNode l2 = new ListNode(2);
        mergeTwoLists(l1,l2);
    }

    /**
     * 26. 删除排序数组中的重复项
     * @param nums
     * @return
     */
    public int removeDuplicates(int[] nums){
        int index = 0;
        for(int i=0;i<nums.length;i++){
            if(nums[i]!=nums[i+1]){
                nums[++index] = nums[i+1];
            }
        }
        return index++;
    }

    /**
     * 189. 旋转数组
     * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数
     * @param nums
     * @param k
     */
    public static void rotate(int[] nums, int k){
        //复制一个数组，分别将需要移动的元素重新赋值到原数组中
        int[] tmpArr = Arrays.copyOf(nums,nums.length);
        int j = tmpArr.length-k;
        for(int i=0;i<k;i++){
            nums[i] = tmpArr[j++];
        }
        j = 0;
        for(int i=k;i<nums.length;i++){
            nums[i]=tmpArr[j++];
        }
    }

    public void rotate2(int[]nums,int k){

    }

    /**
     * 21.将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
     * @param l1
     * @param l2
     * @return
     */
    public static ListNode mergeTwoLists(ListNode l1, ListNode l2) {

        ListNode startNode = new ListNode(0);
        ListNode curNode = startNode;
        while(l1!=null){
            while(l2!=null){
                if(l1.val<l2.val){
                    curNode.next = l1;

                    l1 = l1.next;
                }else{
                    curNode.next = l2;
                    l2 = l2.next;
                }
                curNode = curNode.next;
                if(l2==null||l1==null){
                    break;
                }
            }
            if(l2==null||l1==null){
                break;
            }
        }
        if(l1!=null&&l2==null){
            curNode.next = l1;
        }else if(l2!=null&&l1==null){
            curNode.next = l2;
        }
        return startNode.next;
    }

    /*
     641. 设计循环双端队列
     设计实现双端队列。
     你的实现需要支持以下操作：

     MyCircularDeque(k)：构造函数,双端队列的大小为k。
     insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
     insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
     deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
     deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
     getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
     getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
     isEmpty()：检查双端队列是否为空。
     isFull()：检查双端队列是否满了。
     */

    /**
     * 42. 接雨水
     * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
     * @param height
     * @return
     */
    public int trap(int[] height){
        return 0;
    }

    static class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    /**
     * 242. 有效的字母异位词
     * @param s
     * @param t
     * @return
     */
    public static boolean isAnagram(String s, String t) {
        return false;
    }

    /**
     * 283. 移动零
     * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
     * @param nums
     */
    public void moveZeroes(int[] nums) {

    }
}
