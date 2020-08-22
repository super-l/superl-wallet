// 项目地址:https://github.com/super-l/superl-wallet
// 作者:superl
// 邮箱:86717375@qq.com
// 博客:www.superl.org
// QQ交流群:235586685

package utils

import (
	"../config"
	"../crypto"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/decred/dcrd/dcrec/secp256k1/ecdsa"
	"golang.org/x/crypto/ripemd160"
	"log"
)

// WIF压缩格式私钥转换成16进制格式私钥
func WifToHexPrivateKey(privatekey []byte) string {
	compressed := true
	if len(privatekey) == 51 {
		privatekeyStr := string(privatekey)
		if privatekeyStr[0:1] == "5" {
			// 非压缩格式WIF私钥
			compressed = false
		}
	}
	privatekeyDecode := crypto.Base58Decode(privatekey)
	//fmt.Printf("Base58解码结果:%x \n",privatekeyDecode)

	privatekeyDecodeStr := fmt.Sprintf("%x", privatekeyDecode)

	var privatekeyHexStr string
	//fmt.Print(len(privatekeyDecodeStr))
	if compressed {
		privatekeyHexStr = privatekeyDecodeStr[2 : len(privatekeyDecodeStr)-10]
	} else {
		privatekeyHexStr = privatekeyDecodeStr[2 : len(privatekeyDecodeStr)-8]
	}
	return privatekeyHexStr
}

// 通过公钥，计算比特币地址
func GetAddress(publickey []byte, netVersion byte) []byte {

	publicSHA256 := sha256.Sum256(publickey)
	//fmt.Printf("进行SHA-256哈希计算，得到结果:%X \n",publicSHA256)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)
	//fmt.Printf("计算RIPEMD-160哈希值，得到结果:%X\n",publicSHA256)

	versionedPayload := append([]byte{netVersion}, publicRIPEMD160...)
	//fmt.Printf("加入主网版本号，得到结果:%X\n",versionedPayload)

	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	//fmt.Printf("再进行两次SHA-256计算哈希值，得到结果:%X\n",versionedPayload)

	checksum := secondSHA[:config.AddressChecksumLen]
	//fmt.Printf("取上一步结果的前4个字节（8位十六进制）:得到结果:%X\n",checksum)

	fullPayload := append(versionedPayload, checksum...)
	//fmt.Printf("把4个字节加在第五步的结果后面作为校验位，得到结果:%X\n",fullPayload)

	address := crypto.Base58Encode(fullPayload)
	//fmt.Printf("用Base58编码变换一下地址后得到地址结果:%s\n",address)
	return address
}

// 根据16进制原始私钥，计算出压缩公钥和未压缩公钥
func PrivkeyToPubKey(privkey string) (pubkeyUncompressed string, pubkeyCompressed string, x string, y string) {
	privateKeyBytes, _ := hex.DecodeString(privkey)
	publicKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)

	x = hex.EncodeToString(publicKey.PubKey().X().Bytes())
	y = hex.EncodeToString(publicKey.PubKey().Y().Bytes())

	pubkeyUncompressed = hex.EncodeToString(publicKey.PubKey().SerializeUncompressed())
	pubkeyCompressed = hex.EncodeToString(publicKey.PubKey().SerializeCompressed())

	return pubkeyUncompressed, pubkeyCompressed, x, y
}

// 检测比特币地址是否有效
func IsVaildBitcoinAddress(address string) bool {
	adddressByte := []byte(address)
	fullHash := crypto.Base58Decode(adddressByte)
	if len(fullHash) != 25 {
		return false
	}
	prefixHash := fullHash[:len(fullHash)-config.AddressChecksumLen]
	tailHash := fullHash[len(fullHash)-config.AddressChecksumLen:]

	firstSHA := sha256.Sum256(prefixHash)
	secondSHA := sha256.Sum256(firstSHA[:])
	tailHash2 := secondSHA[:config.AddressChecksumLen]

	if bytes.Compare(tailHash, tailHash2[:]) == 0 {
		return true
	} else {
		return false
	}
}

// 获取签名
func GetSign(key secp256k1.PrivateKey, message []byte) *ecdsa.Signature {
	fmt.Printf("待签名字符:%s\n", message)
	signature := ecdsa.Sign(&key, message)
	return signature
}

// 验证签名
func VerifySign(Signature *ecdsa.Signature, message []byte, pubKey *secp256k1.PublicKey) bool {
	fmt.Printf("待验证字符:%s\n", message)
	flag := Signature.Verify(message, pubKey)
	return flag
}
