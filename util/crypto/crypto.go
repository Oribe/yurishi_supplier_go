package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"path"
)

const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC0HvosNomhSfTbfHbQdQeDARLtm1853PwqmtaRSyoyB40g26iI
nN/U4img4hXIpaK9cBVCQu46OKZXsqa6ebnyUpeER1U4irbQDZbq56YVGKgSsKDQ
zDg7bcCYzp4cXVfZn4hajj5gu73U88RlJgbOZ33cqSXTe37aK9uyCSjHdQIDAQAB
AoGAK8VTWicuruk//Y5zeRjXaHh8Vw1oyLDw/pF+DvTLHjlDjHaUsA4fPqZvI0+N
p6LIt2xjXiTRq4hUs/8QEUS/cMFIt8o8yxSTIVcFEw/LrzRND+TpSNU3IUndySTF
y5GapfNe9ebsofxukbq8ztu/e7vJZLomROdzbaAZGMqmyokCQQDbsLfb1a/w0a8Q
bOVL8KO7+MWGRxN4zjgNJBtp/plXh8xQdLQrTy2uJgys4EZwXGQo0sNm37luF9+Y
0APVd8sTAkEA0eQMcdqvo8sfN+Sky2fJpIArFpehQILjX5jv7gssGjEFA6QrhBps
vi09LkGt29pTLHx5jTJH/1DX3IaiiTasVwJBALraQnbS+BMOdSS2Sgxd/xU4kOAV
geVFGH1s0XhEmK3PDaL8r+UKMiMlr7A2DRMyMepa4OGVtcSCv6XcfvZILksCQFnz
yV2WIoYpDFUQ+YtvDJ0ijNTe4S3bqoSS7+juAtWqwoJf/oJcLNSIYNsNLOy3McQ0
CIf5z59dT1XkoY36z3sCQQCOtWWkSy+kKRIsyjW5Eq21sBwZYCSlH4J+Z7xxE1eK
GxLQcxx+92+82TNdA+0GWPs8o/tYsZiLAejvxHDuCBOI
-----END RSA PRIVATE KEY-----
`

// readPrivateKey 获取私钥
func readPrivateKey() (string, error) {
	keyPath := path.Join("/util/crypto/rsa_private_key.pem")
	pkByte, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}
	privateKey := string(pkByte)
	return privateKey, nil
}

// PrivateDecrypt 私钥解密
func PrivateDecrypt(auth string) (string, error) {

	// privateKey, err := readPrivateKey()
	// if err != nil {
	// 	return "", err
	// }

	authBase64, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return "", errors.New("token错误")
	}

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("私钥错误")
	}

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.New("私钥解码错误")
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, authBase64)
	if err != nil {
		return "", errors.New("token解密失败")
	}

	return base64.StdEncoding.EncodeToString(plainText), nil
}
