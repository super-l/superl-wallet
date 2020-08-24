package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"golang.org/x/crypto/sha3"
)

// 根据字符串私钥计算以太坊地址
func EthGetAddressFromPrivkey(strPrivKey string) (pubkey []byte, address string) {
	var (
		pubKeyHash256 []byte
		buf           []byte
	)

	buf, _ = hex.DecodeString(strPrivKey)
	privKey := secp256k1.PrivKeyFromBytes(buf)

	//拿公钥（非压缩公钥）来hash，计算公钥的 Keccak-256 哈希值（32bytes）
	pubkey = append(privKey.PubKey().X().Bytes(), privKey.PubKey().Y().Bytes()...)

	//计算公钥的Keccak256哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubkey)

	pubKeyHash256 = hash.Sum(nil)

	//取上一步结果取后20bytes即以太坊地址
	address = string("0x") + hex.EncodeToString(pubKeyHash256[len(pubKeyHash256)-20:])
	return pubkey, address
}

//拿公钥（非压缩公钥）来hash，计算公钥的 Keccak-256 哈希值（32bytes）
func EthGetAddressFromPubkey(strPubKey string) string {
	var (
		uncompressPubKey []byte
		pubKeyHash256    []byte
	)

	uncompressPubKey, _ = hex.DecodeString(strPubKey)

	//计算公钥的Keccak256哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(uncompressPubKey)

	pubKeyHash256 = hash.Sum(nil)

	//取上一步结果取后20bytes即以太坊地址
	address := string("0x") + hex.EncodeToString(pubKeyHash256[len(pubKeyHash256)-20:])
	return address
}

//拿公钥（非压缩公钥）来hash，计算公钥的 Keccak-256 哈希值（32bytes）
func EthGetAddressFromPubkeyBytes(uncompressPubKey []byte) {
	var (
		pubKeyHash256 []byte
	)

	//计算公钥的Keccak256哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(uncompressPubKey)

	pubKeyHash256 = hash.Sum(nil)

	//取上一步结果取后20bytes即以太坊地址
	address := string("0x") + hex.EncodeToString(pubKeyHash256[len(pubKeyHash256)-20:])
	fmt.Print(address)
}
