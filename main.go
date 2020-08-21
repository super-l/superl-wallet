// 项目地址:https://github.com/super-l/superl-wallet
// 作者:superl
// 邮箱:86717375@qq.com
// 博客:www.superl.org
// QQ交流群:235586685

package main

import (
	"./config"
	"./utils"
	"./wallet"
	"encoding/hex"
	"fmt"
)

// 16进制公钥字符串，计算出比特币地址，验证是否有效
func BtcPublickeyToAddress(publicKey string) {
	hex_data, _ := hex.DecodeString(publicKey)

	fmt.Print("================通过公钥计算出地址的算法测试==================\n")
	fmt.Printf("当前录入的公钥:%x\n", hex_data)
	address := utils.GetAddress(hex_data, config.BitcoinMainNetVersion)
	fmt.Printf("比特币地址:%s\n", address)
	fmt.Printf("比特币地址是否有效:%v\n", utils.IsVaildBitcoinAddress(string(address)))
}

func main() {
	fmt.Print("\n================正在创建钱包信息================\n")
	mywallrt := wallet.NewWallet()

	// 根据压缩公钥，获取地址
	compressedPubkeyAddress := mywallrt.GetAddress(true, config.BitcoinMainNetVersion)

	// 根据未压缩公钥，获取地址
	UncompressedPubkeyAddress := mywallrt.GetAddress(false, config.BitcoinMainNetVersion)

	pubkey := mywallrt.PublicKey
	pubkeyUncompressed := mywallrt.GetUncompressedPubKey()
	pubkeycompressed := mywallrt.GetCompressedPubKey()

	privkey := mywallrt.GetHashPrivKey()

	privkey_uncompressed := mywallrt.GetUncompressedWifPrivKey() // WIF格式未压缩私钥
	privkey_compressed := mywallrt.GetCompressedWifPrivKey()     // WIF格式压缩私钥

	fmt.Printf("原始公钥XY:%x\n", pubkey)
	fmt.Printf("公钥(未压缩):%x\n", pubkeyUncompressed)
	fmt.Printf("公钥(压缩):%x\n\n", pubkeycompressed)

	fmt.Printf("根据压缩公钥获取地址:%s\n", compressedPubkeyAddress)
	fmt.Printf("根据未压缩公钥获取地址:%s\n\n", UncompressedPubkeyAddress)

	fmt.Printf("私钥(原始16进制格式):%x\n", privkey)
	fmt.Printf("私钥(未压缩WIF格式):%s\n", privkey_uncompressed)
	fmt.Printf("私钥(压缩WIF格式):%s\n\n", privkey_compressed)

	wifToHexPrivate1 := utils.WifToHexPrivateKey(privkey_uncompressed)
	fmt.Printf("逆向解密测试(未压缩WIF格式私钥转原始16进制):%s\n", wifToHexPrivate1)

	wifToHexPrivate2 := utils.WifToHexPrivateKey(privkey_compressed)
	fmt.Printf("逆向解密测试(压缩WIF格式私钥转原始16进制):%s\n\n", wifToHexPrivate2)

	fmt.Print("\n================通过公钥计算出地址的算法测试==================\n")
	fmt.Printf("当前录入的公钥:%x\n", pubkeycompressed)
	address := utils.GetAddress(pubkeycompressed, config.BitcoinMainNetVersion)
	fmt.Printf("比特币地址:%s\n", address)
	fmt.Printf("比特币地址是否有效:%v\n", utils.IsVaildBitcoinAddress(string(address)))

	// 根据私钥，计算出压缩公钥和未压缩公钥
	//privkey := "e8fafc057d16764ae8bc0f7f480c6334fbb0932ee5b378aad27840d43409ea83"
	fmt.Print("\n================通过私钥计算出公钥的算法测试==================\n")
	fmt.Printf("当前录入的原始16进制格式私钥:%x\n", privkey)
	pubkeyUncompressed2, pubkeyCompressed2, x, y := utils.PrivkeyToPubKey(hex.EncodeToString(privkey))
	fmt.Printf("计算出的未压缩公钥:%s\n", pubkeyUncompressed2)
	fmt.Printf("计算出的压缩公钥:%s\n", pubkeyCompressed2)

	fmt.Printf("计算出的椭圆曲线X坐标:%s\n", x)
	fmt.Printf("计算出的椭圆曲线Y坐标:%s\n", y)
}
