package crypto

import (
	"testing"
)

func TestReadPrivateKey(t *testing.T) {
	_, err := readPrivateKey()
	if err != nil {
		t.Error("读取私钥失败")
	}
	// fmt.Println("输出结果↓")
	// fmt.Println(pk)
	// fmt.Println("输出结果↑")

}

func TestPrivateDecrypt(t *testing.T) {

	auth := "WHYpXRcGkh08iBbnfwxSJ2obycVYCrbUogAJvk2txAZKAWhhKYfPP8w4RJku8AR7ZQMu/jOgFfibJe78/AY1L+XiUtwMDYnMgOlrfec6f6swM6BCHfNCeWABDKBrL/N0OqIdbzbhYWtoim5BGTtvlIa0wIIkAzuT4y68oeRhNtU="

	plainText, err := PrivateDecrypt(auth)
	if err != nil {
		t.Error(err.Error())
	}

	accept := "xEdNRalVh4rXRn2mtzj3QrJofRTrBYpIvn/rN+rwMdk/IA5uYLTmYPhJp3qTyP1GC1a9xqqkLLs="
	if plainText != accept {
		t.Error("解密失败")
	}
}
