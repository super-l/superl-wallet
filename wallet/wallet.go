// 项目地址:https://github.com/super-l/superl-wallet
// 作者:superl
// 邮箱:86717375@qq.com
// 博客:www.superl.org
// QQ交流群:235586685

package wallet

import (
	"../config"
	"../crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"golang.org/x/crypto/ripemd160"
	"log"
	"math/big"
)

// 钱包数据类型
type Wallet struct {
	PrivateKey secp256k1.PrivateKey
	PublicKey  []byte
}

// 创建一个钱包数据
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

// 生成钱包的公钥和私钥数据 (公钥是通过私钥获取的)
func newKeyPair() (secp256k1.PrivateKey, []byte) {
	//比特币当中的曲线是secp256k1
	privkey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		log.Panic(err)
	}
	//拼接x和y坐标，就是公钥
	pubKey := append(privkey.PubKey().X().Bytes(), privkey.PubKey().Y().Bytes()...)
	return *privkey, pubKey
}

// 获取原始16进制格式私钥
// 结构体 PrivateKey 的 D 字段为椭圆曲线算法生成的私钥，然后通过 paddedAppend 方法转化为 32 bytes 的私钥
func (w Wallet) GetHashPrivKey() []byte {
	return w.PrivateKey.Serialize()
}

// 获取原始XY公钥
func (w Wallet) GetPubKey() []byte {
	//fmt.Printf("X:%x \n",w.PrivateKey.PubKey().X().Bytes())
	//fmt.Printf("Y:%x \n",w.PrivateKey.PubKey().Y().Bytes())
	pubKey := append(w.PrivateKey.PubKey().X().Bytes(), w.PrivateKey.PubKey().Y().Bytes()...)
	return pubKey
}

// 获取未压缩公钥
func (w *Wallet) GetUncompressedPubKey() []byte {
	pubkeyUncompressedStr := fmt.Sprintf("%s%x%x", "04", w.PrivateKey.PubKey().X().Bytes(), w.PrivateKey.PubKey().Y().Bytes())
	pubkeyUncompressed, _ := hex.DecodeString(pubkeyUncompressedStr)
	return pubkeyUncompressed
}

// 获取压缩公钥
func (w *Wallet) GetCompressedPubKey() []byte {
	//fmt.Printf("X:%x \n",w.PrivateKey.PubKey().X().Bytes())
	//fmt.Printf("Y:%x \n",w.PrivateKey.PubKey().Y().Bytes())
	data := big.NewInt(1)

	result := data.Mod(w.PrivateKey.PubKey().Y(), big.NewInt(2))

	var pubkeyCompressedStr string

	//如果Y坐标的值为偶数，在X坐标前添加前缀0x02
	//如果Y坐标的值为奇数，在X坐标前添加前缀0x03
	if result.Cmp(big.NewInt(0)) == 0 {
		//fmt.Print("偶数\n")
		pubkeyCompressedStr = fmt.Sprintf("%s%x", "02", w.PrivateKey.PubKey().X().Bytes())
	} else {
		//fmt.Print("奇数\n")
		pubkeyCompressedStr = fmt.Sprintf("%s%x", "03", w.PrivateKey.PubKey().X().Bytes())
	}
	pubkeyCompressed, _ := hex.DecodeString(pubkeyCompressedStr)

	return pubkeyCompressed
}

// 获取未压缩WIF私钥 WIF(wallet-import-format)格式，5开头  格式私钥51位
func (w *Wallet) GetUncompressedWifPrivKey() []byte {
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
func (w *Wallet) GetCompressedWifPrivKey() []byte {
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
func (w Wallet) GetAddress(is_compressed bool, netVersion byte) []byte {
	var publicSHA256 [32]byte
	if is_compressed {
		publicSHA256 = sha256.Sum256(w.GetCompressedPubKey())
	} else {
		publicSHA256 = sha256.Sum256(w.GetUncompressedPubKey())
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
