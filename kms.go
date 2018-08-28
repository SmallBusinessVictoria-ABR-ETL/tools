package tools

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"io/ioutil"
	"strings"
)

var keys *kms.KMS

func init() {
	keys = kms.New(session.Must(session.NewSession()))
}

func ListKeys() {
	k, _ := keys.ListKeys(&kms.ListKeysInput{})
	for _, key := range k.Keys {
		fmt.Println(*key.KeyId)
	}
}

func Decrypt(encrypted string) (string, error) {
	reader := strings.NewReader(encrypted)
	b64 := base64.NewDecoder(base64.StdEncoding, reader)
	CiphertextBlob, err := ioutil.ReadAll(b64)
	if err != nil {
		return "", err
	}
	decrypted, err := keys.Decrypt(&kms.DecryptInput{
		CiphertextBlob: CiphertextBlob,
	})
	if err != nil {
		return "", err
	}
	return string(decrypted.Plaintext), nil
}
