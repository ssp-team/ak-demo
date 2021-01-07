# AccessToken 生成算法说明

```
accessKeyId 系统生成的密钥ID，AK开头的32位字符串
ts 是毫秒数的时间戳，13位数字
accessKeySecret 是系统生成的密钥值，16位字符串
mediaId 是媒体app的唯一标识

accessToken 的计算规则：

accessToken = encrypt("accessKeyId" + "|" + ts + "|" + mediaId, accessKeySecret)

encrypt 加密算法： xxtea算法加密后的字节使用标准的base64编码，然后替换+/=为-_*

```

算法参考：

* java版本：https://github.com/xxtea/xxtea-java
* c版本： https://github.com/xxtea/xxtea-c
* golang版本：https://github.com/xxtea/xxtea-go
* php版本：https://github.com/xxtea/xxtea-php

