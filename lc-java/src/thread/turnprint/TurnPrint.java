package thread.turnprint;

/**
 * 1114. 按序打印
 */
public class TurnPrint{

    public static void main(String[] args) {
        int[] nums = {2,1,3};
        Foo foo = new Foo();
        for(int n:nums){
            if(n==1){
                new Thread(new MyThread1(foo)).start();
            }else if(n==2){
                new Thread(new MyThread2(foo)).start();
            }else if(n==3){
                new Thread(new MyThread3(foo)).start();
            }
        }
        try {
            Thread.sleep(1*1000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    static class MyThread1 implements Runnable {
        private Foo foo;
        public MyThread1(Foo foo){
            this.foo = foo;
        }
        @Override
        public void run() {
            try {
                foo.first(new PrintThird("one"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    static class MyThread2 implements Runnable {
        private Foo foo;
        public MyThread2(Foo foo){
            this.foo = foo;
        }
        @Override
        public void run() {
            try {
                foo.second(new PrintThird("two"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }

    static class MyThread3 implements Runnable {
        private Foo foo;
        public MyThread3(Foo foo){
            this.foo = foo;
        }
        @Override
        public void run() {
            try {
                foo.third(new PrintThird("third"));
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
