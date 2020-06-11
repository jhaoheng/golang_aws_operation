package cloudfrontHnadler

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
)

func SignUrl(b []byte) {

	//
	// 	-----BEGIN PRIVATE KEY-----
	// MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQClHYNDPVSF‌​FmWF
	// oKGTqd/n7Dt2+tGXh97KJjVLAqCBZZHlQJ534v2OzFjTgzuMNehD9Y6HnkYF‌​dkRb
	// QzYi2YDROOzRl1bhyyWPA35OGf50r7LiNvSvNPNtswsCuK7ywOcH0yEMKSiW‌​4q5R
	// GKYi42w961EcTQQPrfihavY+c2FYPv4+pXymzaIz9hGBPLHwaHq/QTAyHxPC‌​fkOo
	// s/x3mxUVd7Ni2bz1VJGlyqcNEeU88wTAYMmv8oQ3y2NfKExtYn+W6TCDiq/+‌​ZkOp
	// wacuAU0J7tCNgcXvkq39KH5uza2uSiTniye6uhlkvYWD3s9riIIiekTEiHk/‌​kkc6
	// jMg8HN/7AgMBAAECggEBAJ12u8vQHV6esUrymaTdCG+BVmRtZpyA ...
	// -----END RSA PRIVATE KEY-----
	//

	block, _ := pem.Decode(b)
	PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	keyID := ""
	signer := sign.NewURLSigner(keyID, PrivateKey)

	location, _ := time.LoadLocation("Asia/Taipei")
	obj := ""
	cloudfront_dns := ""
	rawURL := cloudfront_dns + "/" + obj
	signedURL, err := signer.Sign(rawURL, time.Now().In(location).AddDate(0, 0, 1))
	if err != nil {
		log.Fatalf("Failed to sign url, err: %s\n", err.Error())
	}

	fmt.Println(signedURL)
}
