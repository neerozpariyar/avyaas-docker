package utils

import (
	"avyaas/utils/file"
	"encoding/base64"

	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/spf13/viper"
)

/*
encrypt performs AES encryption on the input URL using a key retrieved from the configuration.
It pads the URL to match the block size, generates a random initialization vector (IV), and uses
Cipher Block Chaining (CBC) mode for encryption. The encrypted data includes the IV prepended to
the ciphertext. The encrypted URL is returned as a byte slice.

Parameters:
  - url: The input URL to be encrypted.

Returns:
  - []byte: The encrypted URL.
  - error: An error if encryption fails.
*/
func Encrypt(url []byte) ([]byte, error) {
	// Retrieve the encryption key from the configuration
	encryptKey := []byte(viper.GetString("encryptKey"))

	// Create a new AES cipher block using the encryption key
	block, err := aes.NewCipher(encryptKey)
	if err != nil {
		return nil, err
	}

	// Determine the block size for padding
	blockSize := aes.BlockSize

	// Calculate the length of padding required
	padLength := blockSize - (len(url) % blockSize)

	// Create a padding slice with each byte set to the pad length
	pad := bytes.Repeat([]byte{byte(padLength)}, padLength)

	// Append the padding to the original URL
	paddedURL := append(url, pad...)

	// Create a slice to store the IV and ciphertext
	ciphertext := make([]byte, blockSize+len(paddedURL))

	// Generate a random initialization vector (IV)
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create a new AES CBC encrypter using the block and IV
	mode := cipher.NewCBCEncrypter(block, iv)

	// Encrypt the padded URL and store the result in the ciphertext slice
	mode.CryptBlocks(ciphertext[blockSize:], paddedURL)

	// Return the encrypted URL
	return ciphertext, nil
}

/*
GetEncryptedSignedUrlString generates an encrypted and base64-encoded version of a signed URL.
It utilizes the GetSignedURL function from the "file" package to obtain a signed URL. The obtained
URL is then encrypted using the encrypt function, and the result is base64-encoded.

Parameters:
  - url: A content URL string to be encrypted.

Returns:
  - string: The base64-encoded, encrypted version of the signed URL.
*/
func GetEncryptedSignedUrlString(url string) (string, error) {
	// Obtain a signed URL from the "file" package.
	signedURL, err := file.GetSignedURL(url)
	if err != nil {
		return "", err
	}

	// Encrypt the obtained signed URL.
	signedURLEncrypted, err := Encrypt([]byte(signedURL))
	if err != nil {
		return "", err
	}

	// Encode the encrypted URL using base64.
	encryptedURLString := base64.URLEncoding.EncodeToString(signedURLEncrypted)

	// Return the base64-encoded, encrypted URL.
	return encryptedURLString, nil
}
