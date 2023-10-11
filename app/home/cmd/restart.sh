#/bin/bash
echo '准备编译...'
go build -o ct_home main.go
echo '编译成功'
echo '原始进程号:'
ps aux | grep "ct_home" | awk '{print $2}' | head -1
ps aux | grep "ct_home" | awk '{print $2}' | head -1 | xargs kill
echo '杀死原始进程成功, 准备拉起新进程...'
nohup ./ct_home > ../../../log/ct_home.log 2>&1 & disown
sleep 2
echo '新进程启动成功, 进程号:'
ps aux | grep "ct_home" | awk '{print $2}' | head -1
