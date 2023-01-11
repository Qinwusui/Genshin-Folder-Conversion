# Android端 国际服与国服原神数据文件夹名称互转程序

## 缘起

由于本人：

- 手机内存不足，没地方放两个原神的数据文件
- 手机有Root权限
- 网络状况不佳，不能流畅游玩云·原神
- 发现国际服的数据文件与国服的数据文件可通用
- 会写简单的Golang程序和简单的Shell

## 故

使用Go+Tasker简单写了一个Yuanshen与Genshin两个数据文件夹之间相互转换文件名的工具</br>
即com.miHoYo.GenshinImpact ->com.miHoYo.Yuanshen</br>
与com.miHoYo.Yuanshen -> com.miHoYo.GenshinImpact

## How To Use

## 如何使用

- **有Go编译环境?** clone该仓库，并编译为Android可执行程序

  - ```shell
    #使用Windows10 amd64编译 ，交叉编译需要更改Go编译参数，编译完成后再改回来
    go env -w GOOS=linux ; go env -w GOARCH=arm64 ; go build -o run .\main.go ; go env -w GOOS=windows ;go env -w GOARCH=amd64
    #使用Android arm64编译，在Android上不需要交叉编译
    go build -o run ./main.go
    ```

- **没有Go编译环境?** 直接下载Release中的可执行文件

- **不论如何** 该程序需要放在Android的`/data/`文件夹下以确保拥有较高的权限，**没有root**权限无法移动文件到`/data/`下

- **如果使用Tasker** Tasker只需要新建一个任务，并且在任务中添加一个调用shell的命令即可
  - 代码在`tasker.xml`中，下载后通过Tasker导入即可
  - 运行该任务必须以Root权限运行
  - 程序有两种输出，两个输出可以通过Tasker捕获并存储在变量中
    - `Genshin To Yuanshen`
    - `Yuanshen To Genshin`
- **如果使用Termux 等Android终端工具**
  - 启动终端工具，并cd到存储该程序的目录下，如：`cd /data/tmp`
  - 获取su权限，并为权限授予执行权限，代码：`chmod 700 run`
  - 输入运行程序代码，如：`./run`

## 结论

- 该程序可以省去通过进入文件管理器中手动更改文件名的繁杂
- 该程序可以降低两个原神App导致的数据占用过大
- 该程序需要Root，对用户不是很友好，不过，可以通过Shizuku进行优化
- 搭配Tasker后只需要执行任务就可以完成国际服数据与国服数据文件互转的操作，且还能提供Toast提示用户当前存在的目录属于国服App还是国际服App

## 缺点

- 不能通过监听App启动来执行该程序，因为执行优先级不正确
- 不能同时运行国服App和国际服App，否则可能会有一个程序无法找到数据文件
- 需要Root权限
