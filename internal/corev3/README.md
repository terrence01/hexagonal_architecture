# 第二版CA架構調整項目
* 移除單一usercase package的實作方式
* 
## Interface 有三種使用情境，分別放在domain、ports/in、ports/out 目錄下
* 簡化統一放在 internal/corev3/interfaces，讓所有Service共用
* 移除in/out使用概念

## Trigger Usecase的Command放在ports/in/command
* 移至service層下的command folder

## Adapter層的實作也有分類 in/out
* 扁平化，移除in/out使用概念

## Service層每個Service為單一Usecase
* 改為聚合型

## Controller層每個Controller對應Service層單一Usecase
* 改為聚合型