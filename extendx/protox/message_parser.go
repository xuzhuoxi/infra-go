package protox

import "github.com/xuzhuoxi/infra-go/bytex"

type IProtoMessageParser interface {
}

type defaultProtoMessageParser struct {
}

// ParseMessage
// block0 : eName utf8
// block1 : pid	utf8
// block2 : uid	utf8
// [n]其它信息
func (m *defaultProtoMessageParser) ParseMessage(msgBytes []byte) (name string, pid string, uid string, data [][]byte) {
	index := 0
	buffToData := bytex.DefaultPoolBuffToData.GetInstance()
	defer bytex.DefaultPoolBuffToData.Recycle(buffToData)

	buffToData.WriteBytes(msgBytes)
	name = buffToData.ReadString()
	pid = buffToData.ReadString()
	uid = buffToData.ReadString()
	if buffToData.Len() > 0 {
		for buffToData.Len() > 0 {
			n, d := buffToData.ReadDataTo(msgBytes[index:]) //由于msgBytes前部分数据已经处理完成，可以利用这部分空间
			if nil == d {
				break
			}
			data = append(data, d)
			index += n
		}
	}
	return name, pid, uid, data
}
