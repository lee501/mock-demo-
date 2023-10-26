#### golang如何使用mock测试

1. 安装
```shell
$ go get -u github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen
```

2. 生成mock文件
   - source：设置需要模拟（mock）的接口文件
   - destination：设置 mock 文件输出的地方，若不设置则打印到标准输出中
   - package：设置 mock 文件的包名，若不设置则为 mock_ 前缀加上文件名（如本文的包名会为 mock_person）
```shell
$ mockgen -source=./person/male.go -destination=./mockd/male_mock.go -package=mockd
```

3. 编写测试用例

```go
 /*
 1. 创建mock controller gomock.NewController
 2. 创建一个新的 mock 实例 mock.NewMockMale 
 3. gomock.InOrder 声明给定的调用应按顺序进行
 4. mockMale.EXPECT().Get(id).Return(nil)：这里有三个步骤，EXPECT()返回一个允许调用者设置期望和返回值的对象。Get(id) 是设置入参并调用 mock 实例中的方法。Return(nil) 是设置先前调用的方法出参。简单来说，就是设置入参并调用，最后设置返回值
 5. NewUser(mockMale)：创建 User 实例，值得注意的是，在这里注入了 mock 对象，因此实际在随后的 user.GetUserInfo(id) 调用（入参：id 为 1）中。它调用的是我们事先模拟好的 mock 方法
 6. ctl.Finish()：进行 mock 用例的期望值断言，一般会使用 defer 延迟执行，以防止我们忘记这一操作
 */

```

4. 测试
```shell
go test ./user
```

5. 查看覆盖率
```shell
go test -cover ./user
```

6. 生成多个mock文件
   - 在male中配置 //go:generate mockgen
```shell
$ go generate ./...
```

7. 可视化界面
```shell
$ go test ./... -coverprofile=cover.out
$ go tool cover -html=cover.out
```
