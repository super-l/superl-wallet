# 简介
    // 项目地址:https://github.com/super-l/superl-wallet
    // 作者:superl
    // 邮箱:86717375@qq.com
    // 博客:www.superl.org
    // QQ交流群:235586685
使用GO语言原创研发的一款基于区块链技术的数字代币钱包工具(比特币相关)，并且带有丰富的注释与算法的步骤说明。

目前已经实现比特币钱包相关基础功能。
-  可生成原始公钥
-  可生成压缩公钥
-  可生成未压缩公钥

-  可生成16进制格式私钥
-  可生成WIF未压缩格式私钥
-  可生成WIF压缩格式私钥

-  可通过私钥计算出压缩公钥和未压缩公钥
-  可通过指定公钥，计算出公钥对应的地址

### 可用性说明

   目前网上的大部分GO语言的实例代码中的椭圆曲线算法，都使用的是GO语言官方包中的比特币当中的曲线是secp256r1，而比特币等数字代币使用的是比特币当中的曲线是secp256k1。
    
   其他项目代码中的椭圆曲线选得不对，两条曲线的参数不同！踩了坑！
   
   这将导致相同的私钥在不同的曲线上会计算出不同的公钥，进而导致计算出不同的地址。 所以secp256r1的私钥在比特币区块链上是无法控制代码里的比特币地址的。
    
   本项目中使用了github.com/decred/dcrd/dcrec/secp256k1 中的secp256k1算法。规避了此问题。
   
   同时，本程序生成的私钥，经过本人实际测试，是可以导入目前主流的钱包中的。导入钱包后的地址等也与程序生成的地址相吻合！

### 运行效果

    ================正在创建钱包信息================
    原始公钥XY:1f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f9180e1b828f69dfa6a4046305924703a373bab892209b8b7e38da02aee44d90ab1
    公钥(未压缩):041f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f9180e1b828f69dfa6a4046305924703a373bab892209b8b7e38da02aee44d90ab1
    公钥(压缩):031f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f91
    
    根据压缩公钥获取地址:19EnjPYjPSLovhjZURyCQ6wws3ix9LhiMB
    根据未压缩公钥获取地址:1ETKBpawDfu1GZAoVbzVnjdDRabuxMoTCc
    
    私钥(原始16进制格式):a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14
    私钥(未压缩WIF格式):5K4BwzPw1vehTmAgLrTUeTpu9xzn7eRsDzNpmK5GhUJG9wGbYwr
    私钥(压缩WIF格式):L2h5KiLDR5oPb7wiedHs2MQRCaA28LCkgLFKpamrHdVBtCjxkqum
    
    逆向解密测试(未压缩WIF格式私钥转原始16进制):a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14
    逆向解密测试(压缩WIF格式私钥转原始16进制):a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14
    
    
    ================通过公钥计算出地址的算法测试==================
    当前录入的公钥:031f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f91
    比特币地址:19EnjPYjPSLovhjZURyCQ6wws3ix9LhiMB
    比特币地址是否有效:true
    
    ================通过私钥计算出公钥的算法测试==================
    当前录入的原始16进制格式私钥:a343a91fa1f377ee4f92b402e17cda8b88d00584de555d5194742fea00f18a14
    计算出的未压缩公钥:041f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f9180e1b828f69dfa6a4046305924703a373bab892209b8b7e38da02aee44d90ab1
    计算出的压缩公钥:031f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f91
    计算出的椭圆曲线X坐标:1f84dbf77dd950d661b223717b7163237a9836c45cd4d8ba87ca9b5ceb761f91
    计算出的椭圆曲线Y坐标:80e1b828f69dfa6a4046305924703a373bab892209b8b7e38da02aee44d90ab1
