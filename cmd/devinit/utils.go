package devinit

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	files "github.com/nrc-no/notcore/internal/utils"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func getOrCreateRandomSecretStr(length int, filePaths ...string) (string, error) {
	secretBytes, err := getOrCreateRandomSecret(length, filePaths...)
	if err != nil {
		return "", err
	}
	secretStr := string(secretBytes)
	secretStr = strings.Trim(secretStr, "\n")
	return secretStr, nil
}

func getOrCreateRandomSecret(length int, filePaths ...string) ([]byte, error) {
	filePath := path.Join(filePaths...)
	exists, err := files.FileExists(filePath)
	if err != nil {
		return nil, err
	}
	if exists {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		if len(fileContent) != 0 {
			return fileContent, nil
		}
	}
	if _, err := files.CreateDirectoryIfNotExists(filepath.Dir(filePath)); err != nil {
		return nil, err
	}
	value := []byte(randomStringBase64(length))
	if err := os.WriteFile(filePath, value, os.ModePerm); err != nil {
		return nil, err
	}
	return value, nil
}
func randomStringBase64(length int) string {
	return base64.StdEncoding.EncodeToString([]byte(randomString(length)))
}

func randomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
