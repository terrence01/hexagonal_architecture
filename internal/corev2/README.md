# 第一版CA架構調整項目

## Interface 有三種使用情境，分別放在domain、ports/in、ports/out 目錄下
* 簡化為每一個usecase都有自己的interface，避免傳入不需要使用到的API。 internal/corev2/application/user/createuserservice/interfaces.go
* 移除in/out使用概念

## Trigger Usecase的Commnd(ports/in/command)
* 移至service層usecase下的command.go

## Adapter層的實作也有分類 in/out
* 扁平化，移除in/out使用概念