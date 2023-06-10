package wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"superl-wallet/config"
	"superl-wallet/crypto"

	"golang.org/x/crypto/ripemd160"
)

// 获取未压缩公钥
func (w *Wallet) BtcGetUncompressedPubKey() []byte {
	pubkeyUncompressedStr := fmt.Sprintf("%s%x%x", "04", w.PrivateKey.PubKey().GetX().Bytes(), w.PrivateKey.PubKey().GetY().Bytes())
	pubkeyUncompressed, _ := hex.DecodeString(pubkeyUncompressedStr)
	return pubkeyUncompressed
}

// 获取压缩公钥
func (w *Wallet) BtcGetCompressedPubKey() []byte {
	//fmt.Printf("X:%x \n",w.PrivateKey.PubKey().X().Bytes())
	//fmt.Printf("Y:%x \n",w.PrivateKey.PubKey().Y().Bytes())
	data := big.NewInt(1)

	result := data.Mod(w.PrivateKey.PubKey().GetY(), big.NewInt(2))

	var pubkeyCompressedStr string

	//如果Y坐标的值为偶数，在X坐标前添加前缀0x02
	//如果Y坐标的值为奇数，在X坐标前添加前缀0x03
	if result.Cmp(big.NewInt(0)) == 0 {
		//fmt.Print("偶数\n")
		pubkeyCompressedStr = fmt.Sprintf("%s%x", "02", w.PrivateKey.PubKey().GetX().Bytes())
	} else {
		//fmt.Print("奇数\n")
		pubkeyCompressedStr = fmt.Sprintf("%s%x", "03", w.PrivateKey.PubKey().GetX().Bytes())
	}
	pubkeyCompressed, _ := hex.DecodeString(pubkeyCompressedStr)

	return pubkeyCompressed
}

// 获取未压缩WIF私钥 WIF(wallet-import-format)格式，5开头  格式私钥51位
func (w *Wallet) BtcGetUncompressedWifPrivKey() []byte {
	privatekey := w.GetHashPrivKey()
	//privatekey,_ := hex.DecodeString("1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd")
	//fmt.Printf("私钥:%x \n",privatekey)

	key := append([]byte{config.BitcoinKeyWIFPrefix}, privatekey...)
	//fmt.Printf("增加0x80版本号到私钥的前面，得到结果:%x \n",key)

	firstSHA := sha256.Sum256(key)
	//fmt.Printf("进行第一次sha-256哈希运算，得到结果:%x \n",firstSHA)

	secondSHA := sha256.Sum256(firstSHA[:])
	//fmt.Printf("进行第二次sha-256哈希运算，得到结果:%x \n",secondSHA)

	checksum := secondSHA[:config.AddressChecksumLen]
	//fmt.Printf("上一步结果的前四个字节作为效验位:%x \n",checksum)

	sumstr := append(key, checksum...)
	//fmt.Printf("效验位加在(增加0x80版本号到私钥)结果的后面:%x \n",sumstr)

	wif_privatekey := crypto.Base58Encode([]byte(sumstr))
	//fmt.Printf("WIF格式私钥(未压缩):%s \n",wif_privatekey)

	return wif_privatekey
}

// 获取压缩WIF私钥  WIF-compressed（WIF压缩格式）， 以“L"/"K"开头，52位 最后一位为是否压缩
func (w *Wallet) BtcGetCompressedWifPrivKey() []byte {
	privatekey := w.GetHashPrivKey()
	////字符串转化为16进制的字节
	//privatekey,_ := hex.DecodeString("1e99423a4ed27608a15a2616a2b0e9e52ced330ac530edcc32c8ffc6a526aedd")
	//fmt.Printf("私钥:%x \n",privatekey)

	privatekeyCompressedStr := fmt.Sprintf("%x%s", privatekey, "01")
	privatekeyCompressed, _ := hex.DecodeString(privatekeyCompressedStr)
	//fmt.Printf("私钥+01:%x \n",privatekeyCompressed)

	//wifcode := []byte{01}
	//wifkey := [][]byte{ privatekey,wifcode}
	//privatekey_wif := bytes.Join(wifkey,[]byte{})
	//fmt.Printf("私钥+01:%x \n",privatekey_wif)

	key := append([]byte{config.BitcoinKeyWIFPrefix}, privatekeyCompressed...)
	//fmt.Printf("增加0x80版本号到私钥+01的前面，得到结果:%x \n",key)

	firstSHA := sha256.Sum256(key)
	//fmt.Printf("进行第一次sha-256哈希运算，得到结果:%x \n",firstSHA)

	secondSHA := sha256.Sum256(firstSHA[:])
	//fmt.Printf("进行第二次sha-256哈希运算，得到结果:%x \n",secondSHA)

	checksum := secondSHA[:config.AddressChecksumLen]
	//fmt.Printf("上一步结果的前四个字节作为效验位:%x \n",checksum)

	sumstr := append(key, checksum...)
	//fmt.Printf("效验位加在(增加0x80版本号到私钥)结果的后面:%x \n",sumstr)

	wif_privatekey := crypto.Base58Encode(sumstr)
	//fmt.Printf("WIF格式私钥(压缩):%s \n",wif_privatekey)

	return wif_privatekey
}

// 通过压缩公钥，生成比特币地址
func (w Wallet) BtcGetAddress(is_compressed bool, netVersion byte) []byte {
	var publicSHA256 [32]byte
	if is_compressed {
		publicSHA256 = sha256.Sum256(w.BtcGetCompressedPubKey())
	} else {
		publicSHA256 = sha256.Sum256(w.BtcGetUncompressedPubKey())
	}

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	versionedPayload := append([]byte{netVersion}, publicRIPEMD160...)

	firstSHA := sha256.Sum256(versionedPayload)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:config.AddressChecksumLen]

	fullPayload := append(versionedPayload, checksum...)
	address := crypto.Base58Encode(fullPayload)
	return address
}
