package protox
const (
    // CodeSuc 成功
    CodeSuc int32 = 0
    // CodeProtoFail 协议错误-协议不存在
    CodeProtoFail int32 = 1
    // CodeArgs 参数错误
    CodeArgs int32 = 2
    // CodeInternal 服务器内部错误
    CodeInternal int32 = 3
    // CodeDbQuery 数据库执行错误
    CodeDbQuery int32 = 4
    // CodeTimeout 请求超时
    CodeTimeout int32 = 5
    // CodeRight 权限不足
    CodeRight int32 = 6
    // CodeStatus 状态不匹配
    CodeStatus int32 = 7
    // CodeRepeat 请求重复
    CodeRepeat int32 = 8
    // CodeFreq 请求过于频繁
    CodeFreq int32 = 9
    // CodeOther 其它错误
    CodeOther int32 = 10
)