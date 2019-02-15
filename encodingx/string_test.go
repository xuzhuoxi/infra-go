//
//Created by xuzhuoxi
//on 2019-02-16.
//@author xuzhuoxi
//
package encodingx

//func TestStringHandler(t *testing.T) {
//	str := "hello world"
//	//str := "ab顶你个肺cc"
//	handler := &base64StringHandler{}
//	bs := handler.HandleEncode(str)
//	fmt.Println(111, bs, []byte(str))
//	var str2 string
//	handler.HandleDecode(bs, &str2)
//	fmt.Println(222, str2, len(str2))
//}
//
//func TestBase64StringHandler(t *testing.T) {
//	input := []byte("hello world")
//
//	// 演示base64编码
//	encodeString := base64.StdEncoding.EncodeToString(input)
//	fmt.Println(encodeString)
//
//	// 对上面的编码结果进行base64解码
//	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(string(decodeBytes))
//
//	fmt.Println(111, input, decodeBytes)
//
//	//fmt.Println("///////////////////////////////")
//	//
//	//// 如果要用在url中，需要使用URLEncoding
//	//uEnc := base64.URLEncoding.EncodeToString([]byte(input))
//	//fmt.Println(uEnc)
//	//
//	//uDec, err := base64.URLEncoding.DecodeString(uEnc)
//	//if err != nil {
//	//	log.Fatalln(err)
//	//}
//	//fmt.Println(string(uDec))
//}
