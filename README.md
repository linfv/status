# API公共返回码说明

- 当接口正确的执行返回时, 客户端应该直接忽略返回参数 `code` 和 `message`
- 当接口返回参数 有 `code` 并且 `code` 不为 `0` 时, 接口执行不成功, 请留意返回参数 `message`

<details>
<summary>示例</summary>

|参数名 | 变量 |类型 | 必填 | 描述|示例|
|--------|--------|--------|--------|--------|--------|
|返回码| code | int| 否|  参考公共错误码|0|
|描述信息| message | string[1,128]| 否|  描述信息|SUCCESS|

无返回值接口执行成功, 客户端请忽略错误码
``` json 
{"code":0, "message": "SUCCESS"}
```

有返回值接口执行成功
``` json
{"token":"xxxx", "expires_in": "7200"}
```

接口执行失败
``` json
{"code":-1, "message": "未知异常"}
```

</details>

### 通用返回
| `code` | `message` |描述|
|--------|--------|--------|
|-1| 未知异常 | 系统错误|
|0 | SUCCESS | 成功    |
