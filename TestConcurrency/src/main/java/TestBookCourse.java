import org.apache.log4j.Logger;

import java.util.Random;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * @program: TestBookCourse
 * @description:
 * @author: Riter
 * @create: 2022-02-16 19:34
 **/
public class TestBookCourse {

    private static final Logger log = Logger.getLogger(TestBookCourse.class);

    public static void main(String[] args) {
        CountDownLatch begin = new CountDownLatch(1);

        int allRequestSize = 20000;
        log.info("all request size is " + allRequestSize);
        ExecutorService exec = Executors.newFixedThreadPool(1000);
        CountDownLatch end = new CountDownLatch(allRequestSize);

        Random random = new Random();
        int user_id = 3;
        int user_cnt = 24999;
        int course_id = 7;
        int course_cnt = 199;

        for (int i = 0; i < allRequestSize; i++) {

            exec.execute(new CallTestBookCourse(String.valueOf(user_id + random.nextInt(user_cnt)), String.valueOf(course_id + random.nextInt(course_cnt)), begin, end));
        }

        long startTime = System.currentTimeMillis();
        begin.countDown();
        try {
            end.await();
        } catch (InterruptedException e) {
            e.printStackTrace();
        } finally {
            log.info("all url requests is done!");
            log.info("the success size :" + CallTestBookCourse.successRequest);
            log.info("the fail size: " + CallTestBookCourse.failRequest);
            log.info("the timeout size: " + CallTestBookCourse.timeOutRequest);
            double successRate = (double) CallTestBookCourse.successRequest / allRequestSize;
            log.info("the success rate is: " + successRate * 100 + "%");
            long endTime = System.currentTimeMillis();
            long costTime = endTime - startTime;
            log.info("the total time cost is: " + costTime + " ms");
            log.info("qps is:" + (costTime < 1000 ? allRequestSize : allRequestSize / (costTime / 1000)));
            log.info("sum request time cost is:" + CallTestBookCourse.costTime + " ms");
        }
        exec.shutdown();
        log.info("main method end");
    }


}
