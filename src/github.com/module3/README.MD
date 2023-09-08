# Docker 底层原理
## NameSpace
ipc net pid mnt uts user

- 查询当前系统的 namaspace
``` shell
lsns -t <type>
```
- 查询某个进程的 namespace
```shell
    ls -la /proc/<pid>/ns/
```
- 进入某namespace运行命令
```shell
    nsenter -t <pid> -n ip addr
```
