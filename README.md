# AutoFan

适配泰安S8030主板的多显卡服务器自动转速调节组件。

## 说明

由于S8030主板的系统风扇不能自动调节转速，该组件可以根据服务器上显卡在不同的温度下自动调整风扇转速度。一般情况下，使用该主板的用户可以根据下面的ipmitool命令进行调整风扇的转速，该命令参考自[Chiphell上面的帖子](https://www.chiphell.com/forum.php?mod=viewthread&tid=2604921&extra=page%3D1&mobile=no)：

```shell
ipmitool raw 0x2e 0x44 0xfd 0x19 0x00 <风扇ID> 0x01 <占空比>
```

理论上该组件也适用于其他的机器，用户自行修改程序中IPMI命令格式即可。


## 安装

可以直接从源码编译，使用 redhat 类系统的用户可以自行安装 Go 编译环境编译生成 RPM 文件。

```bash
make dev && make rpm
```

## 使用

安装完成后修改配置文件：

```javascript
mode: max             # max/mean可选，温度阈值是取决于所有显卡中的最高温度，还是所有显卡的平均温度
interval: 10          # 程序检查显卡温度的时间间隔，单位秒
thresholds:           # 设置的温度曲线，温度小于设置的最小的温度(摄氏度)时默认风扇转速为50%
  - temperature: 50   # 50<=temp<60, duty cycle=50%
    duty-cycle: 50
  - temperature: 60   # 60<=temp<70, duty cycle=60%
    duty-cycle: 60
  - temperature: 70   # 70<=temp<90, duty cycle=90%
    duty-cycle: 70
  - temperature: 90   # 90<=temp, duty cycle=100%
    duty-cycle: 100
pwd-ids: [2, 3, 4, 5] # 要自动调整转速的风扇id
gpu-debug: false      # 设置为true时，则用伪GPU数据，仅用于测试
ipmi-debug: false     # 设置为true时，则用伪IPMI自行，仅用于测试
```

然后使用如下的命令启动组件：

```shell
systemctl enable autofan --now
```

## 反馈

有任何问题可以在如下的邮箱进行反馈： sonmihpc@gmail.com


## 作者

- [@wytfy](https://www.github.com/wytfy)