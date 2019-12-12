package thread.turnprint;

public class PrintThird implements Runnable {
    String val;

    public PrintThird(String val){
        this.val = val;
    }

    @Override
    public void run() {
        System.out.println(val);
    }
}
