import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        for (int i = 0; i < T; i++) {
            int N = sc.nextInt(), n = sc.nextInt(), c = sc.nextInt(); // 容量 + 玩具个数 + 填充物个数
            int[] a = new int[n + c];
            for (int j = 0; j < n; j++) {
                a[j] = sc.nextInt();
            }
            for (int j = n; j < n + c; j++) {
                a[j] = 1;
            }
            boolean res = checkFull(N, n, a);

            if (res) {
                System.out.println("YES");
            } else {
                System.out.println("NO");
            }
        }
        sc.close();
    }

    private static boolean checkFull(int N, int c, int[] a) {
        if (N <= c)
            return true;
        boolean[][] dp = new boolean[a.length][N + 1];
        if (a[0] <= N) {
            dp[0][a[0]] = true;
        }
        for (int i = 0; i < a.length; i++) {
            dp[i][0] = true;
        }
        for (int i = 1; i < a.length; i++) {
            for (int j = 1; j <= N; j++) {
                if (a[i] <= j) {
                    dp[i][j] = dp[i - 1][j - a[i]];
                }
                dp[i][j] = dp[i][j] || dp[i - 1][j];
            }
        }

        return dp[a.length - 1][N];
    }
}