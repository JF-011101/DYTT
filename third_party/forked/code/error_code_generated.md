# 错误码
！！IAM 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。
## 功能说明
如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：
```json
{
  "code": 100101,
  "message": "Database error"
}
```
上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。
## 错误码列表
IAM 系统支持的错误码列表如下：
| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrUserNotFound | 110001 | 404 | User not found |
| ErrUserAlreadyExist | 110002 | 400 | User already exist |
| ErrReachMaxCount | 110101 | 400 | Secret reach the max count |
| ErrSecretNotFound | 110102 | 404 | Secret not found |
| ErrVideoNotFound | 120001 | 400 | Video not found |
| ErrCommentNotFound | 120002 | 400 | Comment not found |
| ErrRelationNotFound | 120003 | 400 | Relation not found |
| ErrFavoriteNotFound | 120004 | 400 | Favorite not found |
| ErrSuccess | 0 | 200 | OK |
| ErrUnknown | 100001 | 500 | Internal server error |
| ErrBind | 100002 | 400 | Error occurred while binding the request body to the struct |
| ErrValidation | 100003 | 400 | Validation failed |
| ErrTokenInvalid | 100004 | 401 | Token invalid |
| ErrDatabase | 100101 | 500 | Database base error |
| ErrRecordNotFound | 100102 | 500 | Internal server error |
| ErrInvalidTransaction | 100103 | 500 | Internal server error |
| ErrNotImplemented | 100104 | 500 | Internal server error |
| ErrMissingWhereClause | 100105 | 500 | Internal server error |
| ErrUnsupportedRelation | 100106 | 500 | Internal server error |
| ErrPrimaryKeyRequired | 100107 | 500 | Internal server error |
| ErrModelValueRequired | 100108 | 500 | Internal server error |
| ErrInvalidData | 100109 | 500 | Internal server error |
| ErrUnsupportedDriver | 100110 | 500 | Internal server error |
| ErrRegistered | 100111 | 500 | Internal server error |
| ErrInvalidField | 100112 | 500 | Internal server error |
| ErrEmptySlice | 100113 | 500 | Internal server error |
| ErrDryRunModeUnsupported | 100114 | 500 | Internal server error |
| ErrInvalidDB | 100115 | 500 | Internal server error |
| ErrInvalidValue | 100116 | 500 | Internal server error |
| ErrInvalidValueOfLength | 100117 | 500 | Internal server error |
| ErrPreloadNotAllowed | 100118 | 500 | Internal server error |
| ErrEncrypt | 100201 | 401 | Error occurred while encrypting the user password |
| ErrSignatureInvalid | 100202 | 401 | Signature is invalid |
| ErrExpired | 100203 | 401 | Token expired |
| ErrInvalidAuthHeader | 100204 | 401 | Invalid authorization header |
| ErrMissingHeader | 100205 | 401 | The `Authorization` header was empty |
| ErrorExpired | 100206 | 401 | Token expired |
| ErrPasswordIncorrect | 100207 | 401 | Password was incorrect |
| ErrPermissionDenied | 100208 | 403 | Permission denied |
| ErrEncodingFailed | 100301 | 500 | Encoding failed due to an error with the data |
| ErrDecodingFailed | 100302 | 500 | Decoding failed due to an error with the data |
| ErrInvalidJSON | 100303 | 500 | Data is not valid JSON |
| ErrEncodingJSON | 100304 | 500 | JSON data could not be encoded |
| ErrDecodingJSON | 100305 | 500 | JSON data could not be decoded |
| ErrInvalidYaml | 100306 | 500 | Data is not valid Yaml |
| ErrEncodingYaml | 100307 | 500 | Yaml data could not be encoded |
| ErrDecodingYaml | 100308 | 500 | Yaml data could not be decoded |
| ErrInvalidHash | 100309 | 500 | Encoded hash is not in the correct format |
| ErrIncompatibleVersion | 100310 | 500 | Incompatible version of encryption algorithm |

