import org.apache.log4j.Logger;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.SocketTimeoutException;
import java.net.URL;
import java.util.concurrent.CountDownLatch;

/**
 * @program: CallTestGetStudentCourse
 * @description:
 * @author: Riter
 * @create: 2022-02-16 19:34
 **/
public class CallTestGetStudentCourse implements Runnable {
    private static final Logger log = Logger.getLogger(CallTestBookCourse.class);
    public static int successRequest = 0;
    public static int failRequest = 0;
    public static int timeOutRequest = 0;

    private final CountDownLatch begin;
    private final CountDownLatch end;
    private final String studentID;

    public static long costTime = 0;

    CallTestGetStudentCourse(String studentID, CountDownLatch begin, CountDownLatch end) {
        this.studentID = studentID;
        this.begin = begin;
        this.end = end;
    }

    private static synchronized void incrementSuccessCount() {
        successRequest++;
    }

    private static synchronized void incrementFailCount() {
        failRequest++;
    }

    private static synchronized void incrementTimeOutCount() {
        timeOutRequest++;
    }

    private static synchronized void incrementCostTime(long cost) {
        costTime += cost;
    }


    @Override
    public void run() {
        HttpURLConnection httpURLConnection = null;
        try {
            String local = "localhost:8000";
            String webURL = "180.184.74.1:80";
            begin.await();
            long startTime = System.currentTimeMillis();
            // 1. 获取访问地址URL
            URL url = new URL("http://" + webURL + "/api/v1/student/course?StudentID=" + studentID);
            // 2. 创建HttpURLConnection对象
            httpURLConnection = (HttpURLConnection) url.openConnection();
            /* 3. 设置请求参数等 */
            // 请求方式  默认 GET
            httpURLConnection.setRequestMethod("GET");
            // 超时时间
            // httpURLConnection.setConnectTimeout(300);
            // 设置是否输出
            // httpURLConnection.setDoOutput(true);
            // 设置是否读入
            httpURLConnection.setDoInput(true);
            // 设置文件字符集:
            httpURLConnection.setRequestProperty("Charset", "UTF-8");
            // 设置是否使用缓存
            httpURLConnection.setUseCaches(false);
            // 设置此 HttpURLConnection 实例是否应该自动执行 HTTP 重定向
            httpURLConnection.setInstanceFollowRedirects(true);
            // 连接
            httpURLConnection.connect();
            /* 4. 处理输入输出 */

            // StringBuilder msg = new StringBuilder();
            int code = httpURLConnection.getResponseCode();
            if (code == 200) {
                incrementSuccessCount();
                BufferedReader reader = new BufferedReader(
                        new InputStreamReader(httpURLConnection.getInputStream()));
//                String line;
//                while ((line = reader.readLine()) != null) {
//                    msg.append(line).append("\n");
//                }

                reader.close();
            } else {
                incrementFailCount();
            }
            long endTime = System.currentTimeMillis();
            incrementCostTime(endTime - startTime);
            // System.out.println(msg);
        } catch (SocketTimeoutException e) {
            incrementTimeOutCount();
            log.error(e.getMessage(), e);
        } catch (Exception e) {
            log.error(e.getMessage(), e);
        } finally {
            // 5. 断开连接
            if (null != httpURLConnection) {
                try {
                    httpURLConnection.disconnect();
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
            end.countDown();
        }
    }
}

