
### 1. "获取用户信息"

1. 路由定义

- Url: /userapi/v1/user/info
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. 请求定义


```golang
type UserInfoReq struct {
	UserId int64 `json:"userId"`
}
```


3. 返回定义


```golang
type UserInfoResp struct {
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
}
```
  


### 2. "修改用户信息"

1. 路由定义

- Url: /userapi/v1/user/update
- Method: POST
- Request: `UserUpdateReq`
- Response: `UserUpdateResp`

2. 请求定义


```golang
type UserUpdateReq struct {
	UserId int64 `json:"userId"`
	Nickname string `json:"nickname"`
}
```


3. 返回定义


```golang
type UserUpdateResp struct {
	Flag bool `json:"flag"`
}
```
  

