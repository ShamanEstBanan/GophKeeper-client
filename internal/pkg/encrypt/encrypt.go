package encrypt

import (
	"GophKeeper-client/internal/config"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

const nones = "super_nones"

func EncryptString(data []byte) ([]byte, error) {
	//src := []byte("Этюд в розовых тонах") // данные, которые хотим зашифровать
	fmt.Printf("original: %s\n", data)

	//будем использовать AES256, создав ключ длинной 32 байта
	key, err := getMasterKey() // crypto key
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	//запишем Key в файл

	//определяем блок для шифрования
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	//определяем шифрование GCM
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()

	nonce := key[:nonceSize]

	dst := aesgcm.Seal(nil, nonce, data, nil) //encrypting
	fmt.Printf("encrypted:%x\n", string(dst))

	return dst, nil

}

func DecryptString(data []byte) ([]byte, error) {

	//будем использовать AES256, создав ключ длинной 32 байта
	key, err := getMasterKey() // crypto key
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	//определяем блок для шифрования
	aesblock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	//определяем шифрование GCM
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	nonce := key[:aesgcm.NonceSize()]
	src2, err := aesgcm.Open(nil, nonce, data, nil)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}
	fmt.Printf("decrypted: %s\n", src2)

	return src2, nil
}

func getMasterKey() ([]byte, error) {
	cfg := config.GetConfig()
	keyFile := cfg.Directory + "/master_key"
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	return key, nil
}
