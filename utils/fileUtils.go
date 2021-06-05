package utils

import "os"


// PathIsExist 判断文件/文件夹是否存在
// golang 判断文件/文件夹是否存在的方法为 使用 os.Stat() 函数的错误
// 值来进行判断：
//		（1）如果返回的错误为 nil，说明文件/文件夹存在
//		（2）如果返回的错误使用 os.IsNotExist() 判断为 true，说明文件/文件夹不存在
//		（3）如果返回的错误为其他类型，则不确定文件/文件夹是否存在
func PathIsExist(path string) (bool, error) {
	_, e := os.Stat(path)
	if e != nil {
		if os.IsNotExist(e) {
			return false, nil
		}
		return false, e
	}
	return true, nil
}
