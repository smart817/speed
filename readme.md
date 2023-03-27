零配置，
speed.MiddlewareJwt
speed.Traslator  参数校验
speed.CreateToken
speed.ParseToken
speed.Log 调试日志

var db1 = speed.MySQL(root，123456，db,127.0.0.1，)
var db2 = speed.Redis(root，123456，db,127.0.0.1，)
var db3 = speed.mongodb(root，123456，db,127.0.0.1，)

speed.SetJwt(Secret,Issuer,Time)

speed/tools
tools.Wechat
tools.AliyunSmS
tools.IsEmail
tools.IsPhoneNumber
tools.Uuid
tools.Encode
tools.Decode

speed


git tag v0.0.5 打标签
git push --tags 提交标签
git push origin -d 0.0.2 删除远程标签