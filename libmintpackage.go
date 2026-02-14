package main

import "C"
import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"strings"
	"unsafe"
)

func classifyFile(filename string) string {
	l := strings.ToLower(filename)
	if strings.HasSuffix(l, ".mp3") || strings.HasSuffix(l, ".wav") || strings.HasSuffix(l, ".ogg") {
		return "SOUND"
	}
	if strings.HasSuffix(l, ".txt") || strings.HasSuffix(l, ".pdf") || strings.HasSuffix(l, ".docx") {
		return "DOCUMENT"
	}
	if strings.HasSuffix(l, ".data") {
		return "DATA"
	}
	return "FILE"
}

func compress(data []byte) []byte {
	var buf bytes.Buffer
	w, _ := flate.NewWriter(&buf, flate.BestCompression)
	w.Write(data)
	w.Close()
	return buf.Bytes()
}

func decompress(data []byte) []byte {
	r := flate.NewReader(bytes.NewReader(data))
	var buf bytes.Buffer
	io.Copy(&buf, r)
	r.Close()
	return buf.Bytes()
}

func encryptGCM(key, plaintext []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key[:32])
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	ct := gcm.Seal(nil, nonce, plaintext, nil)
	return ct, nonce
}

func decryptGCM(key, ciphertext, nonce []byte) []byte {
	block, _ := aes.NewCipher(key[:32])
	gcm, _ := cipher.NewGCM(block)
	pt, _ := gcm.Open(nil, nonce, ciphertext, nil)
	return pt
}

func EncryptFile(filename *C.char, content *C.char, contentLen C.int, key *C.char) *C.char {
	fname := C.GoString(filename)
	data := C.GoBytes(unsafe.Pointer(content), contentLen)
	k := C.GoBytes(unsafe.Pointer(key), 32)
	_ = classifyFile(fname)
	comp := compress(data)
	enc, _ := encryptGCM(k, comp)
	return C.CString(string(enc))
}

func DecryptSection(data *C.char, dataLen C.int, key *C.char) *C.char {
	d := C.GoBytes(unsafe.Pointer(data), dataLen)
	k := C.GoBytes(unsafe.Pointer(key), 32)
	nonce := d[:12]
	dec := decryptGCM(k, d, nonce)
	decomp := decompress(dec)
	return C.CString(string(decomp))
}

func main() {}
