package upload

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
)

func UploadAvatarToLocalStatic(file multipart.File, uid uint, account string) (string, error) {
	uid_str := strconv.Itoa(int(uid))
	path := "." + conf.AvatarPath + "user_" + uid_str + "/"
	if !dirExistOrNot(path) {
		createDir(path)
	}
	avatarPath := path + account + ".JPG"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user_" + uid_str + "/" + account + ".JPG", nil
}

func UploadProductToLocalStatic(file multipart.File, uid uint, name string) (string, error) {
	uid_str := strconv.Itoa(int(uid))
	path := "." + conf.ProductPath + "boss_" + uid_str + "/"
	if !dirExistOrNot(path) {
		createDir(path)
	}
	productPath := path + name + ".JPG"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "boss_" + uid_str + "/" + name + ".JPG", nil
}

// 判断文件夹路径是否存在
func dirExistOrNot(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 创建文件夹
func createDir(dirname string) bool {
	return os.MkdirAll(dirname, 755) == nil
}
