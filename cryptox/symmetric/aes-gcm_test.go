// Package symmetric
// Create on 2025/7/13
// @author xuzhuoxi
package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/xuzhuoxi/infra-go/slicex"
)

type aesGcmCaseItem struct {
	Plaintext  []byte
	Key        []byte
	Nonce      []byte
	Ciphertext []byte
	Decrypted  []byte
}

func (o aesGcmCaseItem) ShortCiphertext() []byte {
	return o.Ciphertext[:len(o.Plaintext)]
}

func (o aesGcmCaseItem) Tag() []byte {
	return o.Ciphertext[len(o.Plaintext):]
}

func (o aesGcmCaseItem) String() string {
	pStr := fmt.Sprintf("Plaintext: %v | %s", o.Plaintext, string(o.Plaintext))
	kStr := fmt.Sprintf("Key: %v", o.Key)
	//kStr := fmt.Sprintf("Key: %X", o.Key)
	nStr := fmt.Sprintf("Nonce: %v", o.Nonce)
	//nStr := fmt.Sprintf("Nonce: %X", o.Nonce)
	cStr := fmt.Sprintf("Ciphertext: %v", o.Ciphertext)
	stStr := fmt.Sprintf("ShortCiphertext: %v | Tag: %v", o.ShortCiphertext(), o.Tag())
	dStr := fmt.Sprintf("Decrypted: %v | %s", o.Decrypted, string(o.Decrypted))
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n", pStr, kStr, nStr, cStr, stStr, dStr)
}

func TestAesGcmDemo(t *testing.T) {
	plaintext := []byte("Hello AES-GCM from Go!")
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	_, _ = rand.Read(key)
	_, _ = rand.Read(nonce)

	block, _ := aes.NewCipher(key)
	aesGcm, _ := cipher.NewGCM(block)
	ciphertext := aesGcm.Seal(nil, nonce, plaintext, nil)

	item := &aesGcmCaseItem{
		Plaintext:  plaintext,
		Key:        key,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}
	fmt.Println(item)
}

var (
	demoAesGcmCase = aesGcmCaseItem{
		Plaintext:  []byte("Hello AES-CTR from Go!"),
		Key:        []byte{221, 114, 67, 111, 131, 146, 36, 124, 118, 142, 179, 76, 214, 202, 1, 58, 40, 227, 121, 123, 131, 175, 210, 141, 26, 202, 201, 128, 239, 34, 252, 132},
		Nonce:      []byte{201, 61, 50, 23, 128, 203, 94, 229, 160, 193, 219, 108},
		Ciphertext: []byte{177, 245, 71, 30, 174, 147, 216, 112, 113, 47, 206, 149, 85, 239, 246, 157, 14, 221, 143, 66, 126, 245, 39, 38, 172, 93, 34, 15, 48, 157, 242, 52, 202, 236, 167, 231, 100, 219},
	}
	aesNonceCases = []aesGcmCaseItem{
		{
			Plaintext: []byte("Hello World!Hello World!Hello World!"),
			Key: []byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
			Nonce: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67},
		},
		{
			Plaintext: []byte("Hello World!Hello World!Hello World!666"),
			Key: []byte{
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
				0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef,
			},
			Nonce: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x01, 0x23, 0x45, 0x67},
		},
	}
	aesCiphertextCases = []aesGcmCaseItem{
		{Plaintext: []byte("Hello"), Key: []byte("0123456789abcdef"), Nonce: []byte("abcdef012345"), Ciphertext: []byte{20, 247, 177, 128, 136, 215, 102, 217, 203, 145, 217, 135, 37, 244, 116, 156, 127, 196, 65, 102, 246}},
		{Plaintext: []byte("AES-GCM Test"), Key: []byte("00112233445566778899aabbccddeeff"), Nonce: []byte("112233445566"), Ciphertext: []byte{133, 183, 139, 186, 149, 215, 119, 19, 127, 218, 133, 157, 22, 82, 34, 7, 166, 81, 133, 173, 193, 217, 48, 200, 126, 163, 136, 16}},
		{Plaintext: []byte("Short"), Key: []byte("abcdefabcdefabcdefabcdefabcdefab"), Nonce: []byte("123456789012"), Ciphertext: []byte{107, 226, 65, 100, 30, 112, 222, 199, 171, 108, 214, 224, 123, 7, 157, 135, 165, 117, 98, 34, 208}},
		{Plaintext: []byte(""), Key: []byte("000102030405060708090a0b0c0d0e0f"), Nonce: []byte("aabbccddeeff"), Ciphertext: []byte{235, 142, 20, 13, 64, 110, 194, 203, 102, 24, 66, 142, 254, 248, 90, 112}},
		{Plaintext: []byte("Max Key Length"), Key: []byte("0123456789abcdef0123456789abcdef"), Nonce: []byte("cafebabeface"), Ciphertext: []byte{73, 43, 238, 134, 249, 35, 77, 217, 212, 129, 97, 40, 158, 66, 90, 202, 10, 247, 85, 141, 215, 186, 13, 97, 238, 88, 5, 34, 136, 164}},
		{Plaintext: []byte("Nonce Variation 1"), Key: []byte("00112233445566778899aabbccddeeff"), Nonce: []byte("111111111111"), Ciphertext: []byte{174, 255, 118, 1, 231, 49, 47, 28, 11, 189, 149, 213, 225, 99, 77, 176, 105, 52, 9, 42, 80, 142, 19, 128, 45, 52, 204, 161, 37, 253, 138, 83, 245}},
		{Plaintext: []byte("Nonce Variation 2"), Key: []byte("00112233445566778899aabbccddeeff"), Nonce: []byte("222222222222"), Ciphertext: []byte{168, 181, 164, 216, 204, 48, 140, 185, 184, 201, 21, 106, 7, 22, 116, 219, 124, 100, 121, 9, 234, 106, 90, 119, 32, 69, 156, 82, 94, 145, 100, 57, 1}},
		{Plaintext: []byte("Longer message to test block size handling and padding"), Key: []byte("abcdefabcdefabcdefabcdefabcdefab"), Nonce: []byte("0f1e2d3c4b5a"), Ciphertext: []byte{223, 62, 71, 180, 201, 51, 61, 201, 20, 238, 63, 150, 207, 117, 44, 62, 153, 83, 71, 55, 31, 238, 165, 143, 138, 66, 159, 166, 84, 217, 137, 56, 89, 34, 172, 49, 200, 69, 216, 211, 185, 239, 155, 57, 72, 235, 53, 237, 71, 200, 154, 192, 66, 188, 152, 41, 27, 239, 79, 103, 121, 81, 87, 92, 108, 111, 72, 45, 249, 174}},
		{Plaintext: []byte("Special chars !@#$%^&*()"), Key: []byte("9876543210abcdef9876543210abcdef"), Nonce: []byte("abcdef123456"), Ciphertext: []byte{64, 96, 89, 232, 119, 105, 104, 27, 10, 217, 23, 206, 92, 164, 111, 160, 134, 21, 203, 81, 75, 195, 77, 197, 131, 198, 49, 126, 137, 3, 64, 21, 23, 225, 87, 120, 135, 220, 68, 66}},
		{Plaintext: []byte("中文测试"), Key: []byte("1234567890abcdef1234567890abcdef"), Nonce: []byte("facefeedbeef"), Ciphertext: []byte{31, 166, 71, 112, 79, 169, 43, 213, 92, 99, 228, 190, 215, 148, 92, 195, 171, 132, 211, 80, 120, 85, 252, 180, 123, 8, 58, 63}},
		{Plaintext: []byte("123"), Key: []byte("0123456789abcdef0123456789abcdef"), Nonce: []byte("777777777777"), Ciphertext: []byte{226, 160, 44, 202, 90, 113, 40, 91, 37, 180, 0, 159, 69, 93, 225, 160, 213, 50, 197}},
		{Plaintext: []byte("Long long long long long long long long text"), Key: []byte("99999999999999999999999999999999"), Nonce: []byte("888888888888"), Ciphertext: []byte{64, 73, 28, 44, 240, 184, 13, 225, 148, 39, 113, 196, 162, 204, 10, 114, 13, 245, 103, 30, 188, 119, 154, 131, 130, 192, 205, 74, 191, 80, 3, 251, 84, 201, 226, 161, 15, 197, 28, 104, 15, 194, 192, 64, 225, 213, 168, 163, 4, 41, 115, 53, 149, 161, 91, 0, 120, 79, 221, 86}},
		{Plaintext: []byte("Same nonce, different keys"), Key: []byte("01010101010101010101010101010101"), Nonce: []byte("999999999999"), Ciphertext: []byte{55, 225, 90, 79, 129, 172, 48, 35, 69, 169, 147, 186, 11, 66, 41, 175, 3, 171, 28, 195, 74, 5, 178, 78, 125, 151, 99, 112, 179, 25, 132, 235, 138, 50, 81, 2, 126, 197, 174, 237, 66, 155}},
		{Plaintext: []byte("Same key, different nonce"), Key: []byte("01010101010101010101010101010101"), Nonce: []byte("aaaaaaaaaaaa"), Ciphertext: []byte{181, 3, 37, 55, 32, 165, 67, 12, 199, 127, 193, 97, 191, 204, 234, 89, 6, 16, 34, 254, 225, 208, 1, 140, 123, 130, 244, 17, 78, 218, 184, 181, 72, 32, 176, 21, 142, 28, 243, 59, 14}},
		{Plaintext: []byte("Testing zero key"), Key: make([]byte, 16), Nonce: []byte("bbbbbbbbbbbb"), Ciphertext: []byte{184, 208, 207, 17, 234, 86, 131, 77, 190, 59, 216, 27, 250, 78, 168, 188, 9, 195, 69, 156, 142, 73, 82, 109, 168, 90, 237, 141, 54, 238, 129, 94}},
		{Plaintext: []byte("Testing zero nonce"), Key: []byte("cccccccccccccccccccccccccccccccc"), Nonce: make([]byte, 12), Ciphertext: []byte{96, 144, 134, 199, 14, 166, 87, 235, 95, 200, 190, 92, 28, 5, 78, 47, 144, 14, 164, 131, 215, 167, 64, 8, 128, 160, 10, 142, 192, 181, 84, 242, 214, 39}},
		{Plaintext: []byte("Edge case: single byte"), Key: []byte("aabbccddeeff00112233445566778899"), Nonce: []byte("555555555555"), Ciphertext: []byte{184, 176, 247, 15, 127, 38, 235, 235, 144, 31, 145, 135, 119, 183, 246, 18, 144, 83, 157, 51, 192, 136, 32, 3, 164, 50, 221, 200, 226, 29, 214, 126, 94, 175, 49, 150, 27, 247}},
	}
)

func TestAesGcmDemoCase(t *testing.T) {
	cp := NewAESCipher(demoAesGcmCase.Key)
	c, err := cp.EncryptNonceGCM(demoAesGcmCase.Plaintext, demoAesGcmCase.Nonce)
	if err != nil {
		t.Error(err)
		return
	}
	p, err := cp.DecryptNonceGCM(c, demoAesGcmCase.Nonce)
	if err != nil {
		t.Error(err)
		return
	}
	if !slicex.EqualUint8(p, demoAesCtrCase.Plaintext) {
		t.Log(demoAesCtrCase.Plaintext)
		t.Log(p)
		t.Fatal("plaintext != ciphertext")
	}
}

func TestGcmNonceCases(t *testing.T) {
	title := "GCM"
	for _, value := range aesNonceCases {
		testNoCiphertext(t, title, value)
		fmt.Println("---------- ---------- ---------- ---------- ---------- ----------")
	}
}

func TestGcmCases(t *testing.T) {
	title := "GCM"
	for _, value := range aesCiphertextCases {
		testCiphertext(t, title, value)
		fmt.Println("---------- ---------- ---------- ---------- ---------- ----------")
	}
}

func TestComputeCiphertext(t *testing.T) {
	title := "GCM.ComputeCiphertext"
	for index, value := range aesCiphertextCases {
		ciphertext, err := encryptAESGCM(value.Plaintext, value.Key, value.Nonce)
		if err != nil {
			t.Error(err)
			return
		}
		if !slicex.EqualUint8(ciphertext, value.Ciphertext) {
			fmt.Println(title, index, ciphertext, value.Ciphertext)
			fmt.Println("---------- ---------- ---------- ---------- ---------- ----------")
		} else {
			value.Ciphertext = ciphertext
			t.Log(title, index, value)
		}
	}
}

func encryptAESGCM(plaintext, key, nonce []byte) ([]byte, error) {
	// 创建 AES 密钥块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("无法创建 AES 密钥: %v", err)
	}

	// 创建 GCM 模式（默认 Tag 长度 16 字节）
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("无法创建 GCM 模式: %v", err)
	}

	// 校验 nonce 长度是否符合规范（一般为 12 字节）
	if len(nonce) != gcm.NonceSize() {
		return nil, fmt.Errorf("nonce 长度必须为 %d 字节", gcm.NonceSize())
	}

	// 执行加密，结果包含密文 + Tag
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func testNoCiphertext(t *testing.T, title string, caseVale aesGcmCaseItem) {
	cp := NewAESCipher(caseVale.Key)
	ciphertext, err1 := cp.EncryptNonceGCM(caseVale.Plaintext, caseVale.Nonce)
	decrypted, err2 := cp.DecryptNonceGCM(ciphertext, caseVale.Nonce)
	caseVale.Ciphertext, caseVale.Decrypted = ciphertext, decrypted
	if err1 != nil || err2 != nil {
		t.Fatal(title, caseVale, err1, err2)
		return
	}
	if !slicex.EqualUint8(caseVale.Plaintext, caseVale.Decrypted) {
		t.Fatal(title, caseVale)
		return
	}
	t.Log(title, caseVale)
}
func testCiphertext(t *testing.T, title string, caseVale aesGcmCaseItem) {
	cp := NewAESCipher(caseVale.Key)
	decrypted, err := cp.DecryptNonceGCM(caseVale.Ciphertext, caseVale.Nonce)
	if err != nil {
		t.Fatal(title, caseVale, err)
		return
	}
	caseVale.Decrypted = decrypted
	if !slicex.EqualUint8(caseVale.Plaintext, caseVale.Decrypted) {
		t.Fatal(title, caseVale)
		return
	}
	t.Log(title, caseVale)
}
