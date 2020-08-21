// 项目地址:https://github.com/super-l/superl-wallet
// 作者:superl
// 邮箱:86717375@qq.com
// 博客:www.superl.org
// QQ交流群:235586685

package utils

import "fmt"

// 把字节数组转换为字符串
func ByteToString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02X", b[i])
	}
	return s
}

func PaddedAppend(size uint, dst, src []byte) []byte {
	/*
		把src数组转换成指定长度的数组，长度不够则添加0
			:param size: 要返回的数组长度
			:param dst: byte类型的切片，需要返回的切片
			:param src: 原byte数组
	*/
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
