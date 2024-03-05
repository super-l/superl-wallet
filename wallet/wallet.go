// 项目地址:https://github.com/super-l/superl-wallet
// 作者:superl
// 邮箱:86717375@qq.com
// 博客:www.superl.org
// QQ交流群:235586685

package wallet

import (
	"log"

	"github.com/decred/dcrd/dcrec/secp256k1"
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
	pubKey := append(privkey.X.Bytes(), privkey.Y.Bytes()...)
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
	pubKey := append(w.PrivateKey.X.Bytes(), w.PrivateKey.Y.Bytes()...)
	return pubKey
}
