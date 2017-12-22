package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"math/big"
)

func main() {
	var err error

	rsaInit()

	privateKey := loadRsaPrivateKey()
	// privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Validating...")
	keyErr := privateKey.Validate()
	if keyErr != nil {
		panic(keyErr)
	}

	msg := []byte("Hello world")

	fmt.Println("Signing...")
	sign := signRsaMsg(privateKey, msg)
	fmt.Printf("Sign: %s\n", base64Encode(sign))

	fmt.Println("Verifying...")
	err = verifyRsaMsgSign(&privateKey.PublicKey, msg, sign)
	fmt.Printf("VerifyResult:(error) %v\n", err)

	fmt.Println("Verifying2...")
	err = verifyRsaMsgSign(loadRsaPubKey(), msg, sign)
	fmt.Printf("VerifyResult2:(error) %v\n", err)
}

func loadRsaPrivateKey() *rsa.PrivateKey {
	d := base64Decode("o/C1h62UgqQqlLSopmAHfoEUIQJrPP831Bw9vC1oZCeDsC8nv+D1M/nvY4RbQyduCuplZ3jWld2N17ulI0U11/h01gZ9AWIJ+oX2yaUHFZAq+302NLrqsaNIUY5GNryRAhRGKHj5635k19TUevt0X4jKwT6lCTRU6B4DxIxTNhILujhjNkCGRFOkUutLnAn3B2mjdg2CTxkVEQZMAc8TgrVrV3jJv1cjd8PzXnkYH9Z5KZYP/pG9GV/So6FtMytYqNmqFBdF0A7tNzBc4izWPIZomB516zo3Jm1Hq7ZhdZrYGgWJi46J8CcRlE5lIRpJ6y2Psz/xZEm6ACtxMyXqcQ==")
	n := base64Decode("zCAqc73a1TKwvaQcteK4URcST+w3OqnBnXjJ8TNUR9iLhyQ9LjGf42+ugtxtCpnxAIhDNKuF4FaNfaaZmX6kIrfrVOlpx5Ya0jLK/5/lq79kuucP+hdaRI7xhl8zXE1K3C9pj2bJQHW939vzrypOWVn+gI3REJFHXfTNoSsYjwZcn3zeQ0wSbY9abHNI3ToeI75nxZOFpEx/jMa+h8Z0XmZKtSTB6DpXZNXR8/06jcMmYXGFugsjDUOotK2wvWtbSsru7OdbODGt0KdFWJ+bbxTeZZTPXrcVLsnQRM+iDhjdEm2uMXPVk/Ehrz30FxF+QfDlY6Pyqcmpl992GdJSXw==")
	e := 65537
	primes := [][]byte{
		base64Decode("4ZMJOz348vie5GEDttiR24TcRkfuIJsX4M4Xx/7YTfVjxpJtalWgPo1O1ep/e1ZOJ5w5orzpgGdqU/tAMBnM/wf2odyzp/bXK1y3dtEKq1nkUf4Wo0Sz9Zo369DRhgynkK9H4HBf3bfzplP2n9+XMoGiHunh3qS2FuWCoEMV65c="),
		base64Decode("56iEXGRi/ZT7UE37i20gW92+QcSIzx1kjkahoar+v4yXxwp9UncVciO9OoERRi8m3DB2IRhPmoehmlELBMIZ/ID+zhuUzdunupz/feWX4hpQFYEL4CHkBevZ+gENxQYJivgOVC5dVGC5lM5vqo8sOpNEhfM1/cMygTFdqQvYyHk="),
	}

	key := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: newBigIntFromBytes(n),
			E: e,
		},
		D:      newBigIntFromBytes(d),
		Primes: newBigIntArrayFromBytesArray(primes),
	}

	return key
}

func loadRsaPubKey() *rsa.PublicKey {
	n := base64Decode("zCAqc73a1TKwvaQcteK4URcST+w3OqnBnXjJ8TNUR9iLhyQ9LjGf42+ugtxtCpnxAIhDNKuF4FaNfaaZmX6kIrfrVOlpx5Ya0jLK/5/lq79kuucP+hdaRI7xhl8zXE1K3C9pj2bJQHW939vzrypOWVn+gI3REJFHXfTNoSsYjwZcn3zeQ0wSbY9abHNI3ToeI75nxZOFpEx/jMa+h8Z0XmZKtSTB6DpXZNXR8/06jcMmYXGFugsjDUOotK2wvWtbSsru7OdbODGt0KdFWJ+bbxTeZZTPXrcVLsnQRM+iDhjdEm2uMXPVk/Ehrz30FxF+QfDlY6Pyqcmpl992GdJSXw==")
	e := 65537
	return &rsa.PublicKey{
		N: newBigIntFromBytes(n),
		E: e,
	}
}

func newBigIntFromBytes(bytes []byte) *big.Int {
	x := big.NewInt(0)
	x.SetBytes(bytes)
	return x
}

func newBigIntArrayFromBytesArray(bytesArr [][]byte) []*big.Int {
	a := make([]*big.Int, len(bytesArr))
	for i, bytes := range bytesArr {
		a[i] = newBigIntFromBytes(bytes)
	}
	return a
}

func rawSha1(msg []byte) []byte {
	s := newSha1()
	s.Write(msg)
	return s.Sum(make([]byte, 0))
}

func newSha1() hash.Hash {
	return sha1.New()
}

func rsaInit() {
	crypto.RegisterHash(crypto.SHA1, newSha1)
}

func signRsaMsg(privateKey *rsa.PrivateKey, msg []byte) (sign []byte) {
	digest := rawSha1(msg)
	return signRsaDigest(privateKey, digest)
}

func signRsaDigest(privateKey *rsa.PrivateKey, digest []byte) (sign []byte) {
	sign, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA1, digest, nil)
	if err != nil {
		panic(err)
	}

	return
}

func verifyRsaMsgSign(pubKey *rsa.PublicKey, msg []byte, sign []byte) (err error) {
	digest := rawSha1(msg)
	return verifyRsaDigestSign(pubKey, digest, sign)
}

func verifyRsaDigestSign(pubKey *rsa.PublicKey, digest []byte, sign []byte) (err error) {
	return rsa.VerifyPSS(pubKey, crypto.SHA1, digest, sign, nil)
}

func base64Encode(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func base64Decode(str string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return bytes
}
