package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max) // 一意の番号であることが望まれるが、自己署名証明書であるためランダムな大きな整数で十分
	subject := pkix.Name{
		Organization:   []string{"Manning Publications Go."},
		OrganizationalUnit: []string{"Books"},
		CommonName: "Go Web Programming",
	}

	// 構造体Certificateを用意する
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:    subject,
		NotBefore:  time.Now(),
		NotAfter:   time.Now().Add(365 * 24 * time.Hour), // 有効期間は作成から一年
		KeyUsage:   x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // X.509証明書がサーバ認証に使用されることを示す
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // X.509証明書がサーバ認証に使用されることを示す
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}, // IPアドレス:127.0.0.1を対象にする
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) // RSA の秘密鍵を生成する

	// 証明書を作成
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // DER形式のバイトデータのスライスを生成
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}) //証明書データを符号化して、 サーバ証明書ファイルを生成
	certOut.Close()

	// 秘密鍵を生成
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type:"RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) // RSAの秘密鍵を符号化して鍵ファイルを生成
	keyOut.Close()
}