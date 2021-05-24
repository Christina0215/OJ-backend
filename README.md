# QKCODE 后端

# 环境准备
> golang : 1.13^
>
> mysql : 8.0^
>
> goose : go get bitbucket.org/liamstask/goose/cmd/goose

# 启动前
>
> `cp env.example env.toml`
>
> `cp db/dbconf.example.yml db/dbconf.yml`
>
> 按照要求完善上述两个配置文件
>
> `goose up` 该命令执行数据库迁移
>
> 启动项目 `go run main.go`