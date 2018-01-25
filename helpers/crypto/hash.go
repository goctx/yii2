package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func Md5(reader io.Reader) string {
	m := md5.New()
	io.Copy(m, reader)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func Sha1(reader io.Reader) string {
	m := sha1.New()
	io.Copy(m, reader)
	return fmt.Sprintf("%x", m.Sum(nil))
}
