## 文件操作系统的API
#### 登陆
url:http://localhost:8080/v1/user/login
method:post
- 请求参数：
```bash
form
user_name:user
pass_word:123456
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "user_id":1
    }
}

```
#### 注册
url:http://localhost:8080/v1/user/register
method:post
- 请求参数：
```
user_name:user
pass_word:123456
```
- 响应
```json
{
    code:0,
    msg:"success",
    data:{
        "user_id":1
    }
}
```
#### 上传文件
url:http://localhost:8080/v1/auth/file/upload
method:post
- 请求参数：
```
file:a.txt
```
- 响应
```json
{
    "code": 0,
    "data": {
        "file_hash": "ec227ef30b07afab302cd668895c4dec7812e14f83cc6214d2326923699527d3",
        "file_id": 26,
        "file_name": "b.docx"
    },
    "msg": "SUCCESS"
}
```
#### 下载文件
url:http://localhost:8080/v1/auth/file/download?id=1
method:get
- 请求参数：
```
id:文件
```
- 响应
文件下载
#### 获取文件列表
url:http://localhost:8080/v1/auth/files/:page_num
method:get
- 请求参数：
```
page_num:分页数
```
- 响应
```json
{
    "code": 0,
    "data": [
        {
            "create_date": "2019-03-27T15:05:09.494465Z",
            "file_id": 19,
            "file_name": "a.docx"
        },
        {
            "create_date": "2019-03-27T15:05:26.187876Z",
            "file_id": 20,
            "file_name": "test.docx"
        },
        {
            "create_date": "2019-03-27T15:05:36.922464Z",
            "file_id": 21,
            "file_name": "a.docx"
        }
    ],
    "msg": "SUCCESS"
}
```