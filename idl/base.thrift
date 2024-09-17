enum Code {
    SUCCESS = 0,         // 操作成功
    INVALID_REQUEST = 1, // 无效请求
    NOT_FOUND = 2,       // 资源未找到
    SERVER_ERROR = 3,    // 服务器错误
    UNAUTHORIZED = 4,    // 未授权
    FAIL = 5
}

struct BaseResp {
    1: required Code code,
    2: required string msg,
}