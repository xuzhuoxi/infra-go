// Package yamlx provides YAML encoding handlers and buffer codecs built on encodingx and gopkg.in/yaml.v2.
// Package yamlx 基于 encodingx 与 gopkg.in/yaml.v2 提供 YAML 编解码处理器及缓冲编解码器。
//
// Created by xuzhuoxi on 2019-02-24.
// @author xuzhuoxi
//
// Capabilities:
//   - NewYamlCodingHandlerAsync/Sync: ICodingHandler backed by yaml.Marshal/Unmarshal
//   - NewYamlBuffEncoder, NewYamlBuffDecoder, NewYamlBuffCodecs: bytex buffer integration
//   - Default pools and NewDefault* factories for handler and buff codec reuse
//
// 主要职责：
//   - YAML 编解码处理器（yaml.Marshal/Unmarshal 实现 ICodingHandler）
//   - 与 bytex 集成的 YAML 缓冲编码器、解码器与双向编解码器
//   - 默认对象池及 NewDefault* 工厂，便于复用处理器与缓冲实例
package yamlx
