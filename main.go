// 项目地址:https://github.com/Scorpio69t/superl-wallet
// 作者:Scorpio69t
// 邮箱:yangruitao6@gmail.com
// 博客:https://blog.csdn.net/qq_32019341?type=blog
// QQ: 3324729792

package main

import (
	"encoding/hex"
	"fmt"
	"superl-wallet/config"
	"superl-wallet/utils"
	"superl-wallet/wallet"
)

// [比特币]16进制公钥字符串，计算出比特币地址，验证是否有效
func BtcPublickeyToAddress(publicKey string) {
	//publicKey := "0384a6728273ddb4b769d9c4dc45fe73594bed3d8364ec50d156b5564f2301e1ca"
	fmt.Print("\n================通过[16进制公钥字符串]计算出[地址]的算法测试==================\n")
	fmt.Printf("当前录入的公钥:%s\n", publicKey)
	hex_data, _ := hex.DecodeString(publicKey)
	address := utils.GetAddress(hex_data, config.BitcoinMainNetVersion)
	fmt.Printf("比特币地址:%s\n", address)
	fmt.Printf("比特币地址是否有效:%v\n", utils.IsVaildBitcoinAddress(string(address)))
}

// [比特币]根据指定16进制私钥，计算出公钥和地址 算法测试
func BtcPrivkeyToPubkeyAndAddress(privkey string) {
	fmt.Print("\n================根据指定16进制私钥，计算出公钥和地址 算法测试==================\n")
	//privkey := "a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14"
	fmt.Printf("当前录入的原始16进制格式私钥:%s\n", privkey)
	pubkeyUncompressedTest, pubkeyCompressedTest, x, y := utils.PrivkeyToPubKey(privkey)
	fmt.Printf("计算出的未压缩公钥:%s\n", pubkeyUncompressedTest)
	fmt.Printf("计算出的压缩公钥:%s\n", pubkeyCompressedTest)

	fmt.Printf("计算出的椭圆曲线X坐标:%s\n", x)
	fmt.Printf("计算出的椭圆曲线Y坐标:%s\n", y)

	pubkeyCompressedTestBytes, _ := hex.DecodeString(pubkeyCompressedTest)
	fmt.Printf("根据计算出的压缩公钥，再计算出地址为:%s\n", utils.GetAddress(pubkeyCompressedTestBytes, config.BitcoinMainNetVersion))
}

// 根据指定的压缩WIF格式私钥，计算出16进制私钥 算法测试
func BtcWifPrivkeyToRawPrivkey(wifprivkey string) {
	fmt.Print("\n================根据指定的压缩WIF格式私钥，计算出16进制私钥 算法测试==================\n")
	//wifprivkey := "L2h5KiLDR5oPb7wiedHs2MQRCaA28LCkgLFKpamrHdVBtCjxkqum"
	//hex_data, _ := hex.DecodeString(wifprivkey)
	fmt.Printf("当前录入的压缩WIF格式私钥:%s\n", wifprivkey)
	wifToHexPrivateTest := utils.WifToHexPrivateKey([]byte(wifprivkey))
	fmt.Printf("计算出的16进制私钥:%s\n", wifToHexPrivateTest)
}

func main() {
	fmt.Print("\n================正在生成测试信息================\n")
	mywallrt := wallet.NewWallet()

	// 根据压缩公钥，获取地址
	compressedPubkeyAddress := mywallrt.BtcGetAddress(true, config.BitcoinMainNetVersion)

	// 根据未压缩公钥，获取地址
	UncompressedPubkeyAddress := mywallrt.BtcGetAddress(false, config.BitcoinMainNetVersion)

	pubkey := mywallrt.PublicKey
	pubkeyUncompressed := mywallrt.BtcGetUncompressedPubKey()
	pubkeycompressed := mywallrt.BtcGetCompressedPubKey()

	privkey := mywallrt.GetHashPrivKey()

	privkey_uncompressed := mywallrt.BtcGetUncompressedWifPrivKey() // WIF格式未压缩私钥
	privkey_compressed := mywallrt.BtcGetCompressedWifPrivKey()     // WIF格式压缩私钥

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

	fmt.Print("\n================签名算法测试==================\n")
	message := []byte("hello,dsa签名")
	verify_message := []byte("hello,dsa签名2")

	signature := utils.GetSign(mywallrt.PrivateKey, message)
	fmt.Printf("计算出的签名:%x\n", signature.Serialize())

	// 验证签名
	flag := utils.VerifySign(signature, message, mywallrt.PrivateKey.PubKey())
	if flag {
		fmt.Printf("【签名测试】:%s 【结果】:%s\n", message, "通过,数据未被修改")
	} else {
		fmt.Printf("【签名测试】:%s 结果】:%s\n", message, "异常,数据被修改")
	}

	flag = utils.VerifySign(signature, verify_message, mywallrt.PrivateKey.PubKey())
	if flag {
		fmt.Printf("【签名测试】:%s 结果】:%s\n", verify_message, "通过,数据未被修改")
	} else {
		fmt.Printf("【签名测试】:%s 结果】:%s\n", verify_message, "异常,数据被修改")
	}

	BtcPublickeyToAddress("0384a6728273ddb4b769d9c4dc45fe73594bed3d8364ec50d156b5564f2301e1ca")
	BtcPrivkeyToPubkeyAndAddress("a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14")
	BtcWifPrivkeyToRawPrivkey("L2h5KiLDR5oPb7wiedHs2MQRCaA28LCkgLFKpamrHdVBtCjxkqum")

	// 根据公钥，生成以太坊地址
	EthGetAddressFromPubkey("156f94de5208dd2ac61c301d51a4c93f1db00a92185b74f545984e8007de214132a0a20edb8019bd049f292f1fe8122e05d4d32ea206674a4fbd6d2376654da7")
	// 根据私钥，生成以太坊地址
	EthGetAddressFromPrivkey("6101eedcc7f11d60baeb4346a7e9810fc42dc9fc267254163bea19fe75bba16e")

}

// 根据公钥，生成以太坊地址
func EthGetAddressFromPubkey(pubkey string) {
	fmt.Print("\n================根据公钥，生成以太坊地址==================\n")
	fmt.Printf("当前录入的以太坊公钥字符串：%s\n", pubkey)
	address := wallet.EthGetAddressFromPubkey(pubkey)
	fmt.Printf("生成以太坊地址：%s\n\n", address)
}

// 根据私钥，生成公钥和以太坊地址
func EthGetAddressFromPrivkey(privkey string) {
	fmt.Print("\n================根据私钥，生成公钥和以太坊地址==================\n")
	fmt.Printf("当前录入的以太坊私钥字符串：%s\n", privkey)
	pubkey, address := wallet.EthGetAddressFromPrivkey(privkey)
	fmt.Printf("根据私钥计算出公钥结果：%x\n", pubkey)
	fmt.Printf("以太坊地址：%s\n\n", address)

}
