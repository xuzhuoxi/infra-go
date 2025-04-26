# HttpCode 状态码

包**net/http**中提供了部分状态码的常量定义。

HTTP 状态码是由 3 位数字组成的，每个状态码有不同的含义，表示 HTTP 请求的结果。HTTP 状态码分为五个类别：

### 1. **1xx - 信息性状态码**

这些状态码表示请求已被接收，继续处理。

- **100 Continue**：表示请求的前半部分已被接收，客户端可以继续发送剩余部分。
- **101 Switching Protocols**：服务器根据客户端的请求切换协议（例如，从 HTTP/1.1 到 HTTP/2）。
- **102 Processing**：服务器正在处理请求，但尚未完成。

### 2. **2xx - 成功状态码**

这些状态码表示请求成功，服务器已经成功处理了请求。

- **200 OK**：请求成功，服务器返回了请求的数据。
- **201 Created**：请求成功，且服务器创建了新的资源（通常用于 POST 请求）。
- **202 Accepted**：请求已接受，但尚未处理（例如，异步任务处理）。
- **203 Non-Authoritative Information**：服务器返回的信息不是来自原始服务器，而是来自缓存的副本。
- **204 No Content**：请求成功，但没有内容返回（例如，DELETE 请求后没有返回数据）。
- **205 Reset Content**：请求成功，且客户端应重置（清空）页面的内容。
- **206 Partial Content**：服务器返回部分数据（通常用于范围请求，如下载部分文件）。

### 3. **3xx - 重定向状态码**

这些状态码表示客户端需要进一步的操作才能完成请求。

- **300 Multiple Choices**：请求有多个可能的响应选项，客户端需要选择一个。
- **301 Moved Permanently**：请求的资源已永久移动到新的 URI，客户端应使用新 URI。
- **302 Found**：请求的资源临时移动，客户端应继续使用原 URI。
- **303 See Other**：服务器建议客户端使用另一个 URI 获取资源，通常是 POST 请求后返回。
- **304 Not Modified**：资源未修改，客户端可以使用缓存的版本。
- **305 Use Proxy**：客户端必须使用代理服务器来访问请求的资源。
- **306 (Unused)**：已不再使用。
- **307 Temporary Redirect**：请求的资源临时移动，客户端应使用新 URI，且请求方法不变。
- **308 Permanent Redirect**：请求的资源永久移动，客户端应使用新 URI，且请求方法不变。

### 4. **4xx - 客户端错误状态码**

这些状态码表示请求有错误，客户端的请求有问题。

- **400 Bad Request**：请求无效，服务器无法理解。
- **401 Unauthorized**：请求未授权，客户端需要提供身份验证信息。
- **402 Payment Required**：保留状态码，当前未使用，预留给将来的支付功能。
- **403 Forbidden**：服务器拒绝请求，客户端没有权限访问资源。
- **404 Not Found**：请求的资源不存在或无法找到。
- **405 Method Not Allowed**：请求方法不被允许（例如，使用 POST 请求访问只支持 GET 的资源）。
- **406 Not Acceptable**：请求的资源不符合客户端的需求（如 MIME 类型不匹配）。
- **407 Proxy Authentication Required**：客户端需要通过代理服务器进行身份验证。
- **408 Request Timeout**：请求超时，客户端没有在规定时间内完成请求。
- **409 Conflict**：请求与当前资源的状态冲突（例如，版本冲突）。
- **410 Gone**：请求的资源已被永久删除，且不再可用。
- **411 Length Required**：服务器要求请求包含 `Content-Length` 头部。
- **412 Precondition Failed**：请求头中的某个条件失败（例如 `If-None-Match`）。
- **413 Payload Too Large**：请求体太大，服务器无法处理。
- **414 URI Too Long**：请求的 URI 太长，服务器无法处理。
- **415 Unsupported Media Type**：请求体的媒体类型不被支持。
- **416 Range Not Satisfiable**：请求的部分内容无法提供（例如，文件范围超出文件大小）。
- **417 Expectation Failed**：服务器无法满足 `Expect` 请求头的要求。
- **418 I'm a teapot**：这个状态码是一个玩笑，源自 HTTP 418 草率协议。
- **421 Misdirected Request**：请求被发送到不适当的服务器。
- **422 Unprocessable Entity**：请求格式正确，但服务器无法处理。
- **423 Locked**：请求的资源已被锁定。
- **424 Failed Dependency**：由于前一个请求失败，导致当前请求失败。
- **426 Upgrade Required**：客户端需要升级协议（例如，HTTP 升级到 HTTPS）。
- **428 Precondition Required**：请求必须包含某些条件（例如，If-Match）。
- **429 Too Many Requests**：客户端在给定的时间内发送了太多请求（例如，速率限制）。
- **431 Request Header Fields Too Large**：请求头字段过大。
- **451 Unavailable For Legal Reasons**：由于法律原因，服务器无法提供该资源。

### 5. **5xx - 服务器错误状态码**

这些状态码表示服务器在处理请求时发生了错误。

- **500 Internal Server Error**：服务器遇到意外情况，无法完成请求。
- **501 Not Implemented**：服务器不支持请求的功能。
- **502 Bad Gateway**：服务器作为网关或代理时，从上游服务器收到无效响应。
- **503 Service Unavailable**：服务器当前不可用（通常是过载或停机维护）。
- **504 Gateway Timeout**：服务器作为网关或代理时，未能在规定时间内从上游服务器获取响应。
- **505 HTTP Version Not Supported**：服务器不支持请求使用的 HTTP 协议版本。
- **506 Variant Also Negotiates**：服务器存在配置错误，导致无法处理内容协商。
- **507 Insufficient Storage**：服务器无法存储完成请求所需的表示（例如，磁盘满）。
- **508 Loop Detected**：服务器在处理请求时检测到无限循环。
- **510 Not Extended**：请求需要扩展才能被处理。
- **511 Network Authentication Required**：需要网络身份验证才能访问请求的资源。
