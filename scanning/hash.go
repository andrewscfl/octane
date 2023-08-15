package scanning

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5HashFile(filePath string) (string, error) {
	var md5String string
	file, err := os.Open(filePath)
	if err != nil {
		return md5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return md5String, err
	}
	//Get the hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert to string
	md5String = hex.EncodeToString(hashInBytes)
	return md5String, nil
}
