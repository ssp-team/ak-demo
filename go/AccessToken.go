package main

import (
    "fmt"
    "time"
    "strings"
    "strconv"
    "encoding/base64"
    "github.com/xxtea/xxtea-go/xxtea"
)

func createAccessToken(accessKeyId string, mediaId string, accessKeySecret string) string{

	ts := strconv.FormatInt(time.Now().Unix() * 1000,10)
        tempStr := accessKeyId + "|" + string(ts) + "|" + mediaId
        createAccessToken := xxtea.Encrypt( []byte(tempStr), []byte(accessKeySecret))
	var createAccessTokenStr = base64.StdEncoding.EncodeToString(createAccessToken)
        createAccessTokenStr = strings.Replace(createAccessTokenStr,"+", "-",-1)
        createAccessTokenStr = strings.Replace(createAccessTokenStr,"/", "_",-1)
        createAccessTokenStr = strings.Replace(createAccessTokenStr,"=", "*",-1)

        return createAccessTokenStr
}

func main() {
    accessKeyId := "accessKeyId"
    mediaId := "mediaId"
    accessKeySecret := "accessKeySecret"

    fmt.Println("accessKeyId: ", accessKeyId)
		
    accessToken := createAccessToken(accessKeyId,mediaId,accessKeySecret)

    fmt.Println("accessToken: ", accessToken)
}
