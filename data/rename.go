package data

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"im_test/pkg"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const fileBytesRead int64 = 120 // only read first X bytes

//ReadFileBytes read first X bytes each file, return hex and add nano time tags(SHA key)
func ReadFileBytes(path string) (fileHex string, key string, err error) {
	rand.Seed(time.Now().UnixNano())
	stream, err := os.Open(path)
	if err != nil {
		return "nil", "nil", errors.Wrapf(err, "open file error, path (%v)", path)
	}
	defer stream.Close() // run after function exit

	info, err := stream.Stat()
	if err != nil {
		return "nil", "nil", errors.Wrapf(err, "get file attribute error, path (%v)", path)
	}

	buffer := make([]byte, pkg.MinInt64(info.Size(), fileBytesRead))
	bytesRead, err := stream.Read(buffer)
	if err != nil {
		return "nil", "nil", errors.Wrapf(err, "read file error, path (%v)", path)
	}

	key = strconv.Itoa(time.Now().Nanosecond() + rand.Int()) // nano second tag for divide same file
	fileHex = hex.EncodeToString(buffer[:bytesRead])

	return
}

//GetHash256 fileHex + key => SHA256
func GetHash256(fileHex string, key string) (encode string, err error) {
	h := hmac.New(sha256.New, []byte(key))
	_, err = h.Write([]byte(fileHex))
	if err != nil {
		return "nil", errors.Wrap(err, "sha256 error")
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

//rename rename file
func rename() {

}
