package thread.turnprint;

public class Foo {

    boolean firstFinshed = false;
    boolean secondFinsed = false;
    Object lock = new Object();

    public Foo() {

    }

    public void first(Runnable printFirst) throws InterruptedException {

        synchronized(lock) {
            // printFirst.run() outputs "first". Do not change or remove this line.
            printFirst.run();
            firstFinshed = true;
            lock.notifyAll();
        }
    }

    public void second(Runnable printSecond) throws InterruptedException {

        synchronized (lock) {
            while(!firstFinshed){
                lock.wait();
            }
            // printSecond.run() outputs "second". Do not change or remove this line.
            printSecond.run();
            secondFinsed = true;
            lock.notifyAll();
        }
    }

    public void third(Runnable printThird) throws InterruptedException {

        synchronized (lock) {
            // printThird.run() outputs "third". Do not change or remove this line.
            while(!secondFinsed) {
                lock.wait();
            }
            printThird.run();
        }
    }


}
