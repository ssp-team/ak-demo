

public class AccessToken {


    public static void main(String[] args) {

        String accessKeyId = "accessKeyId";
        String accessKeySecret = "accessKeySecret";
        String mediaId = "mediaId";

        String createAccessToken = createAccessToken(accessKeyId, mediaId, accessKeySecret);

        System.out.println("createAccessToken: " + createAccessToken);
    }

    public static String createAccessToken(String accessKeyId, String mediaId, String accessKeySecret) {

        String ts = String.valueOf(System.currentTimeMillis());

        String tempStr = accessKeyId + "|" + ts + "|" + mediaId;

        String createAccessToken = XXTEA.encryptToBase64String(tempStr, accessKeySecret);
        createAccessToken = createAccessToken.replace("+", "-").replace("/", "_").replace("=", "*");

        return createAccessToken;
    }
}
