# Benchlist

[Container Lifecycle](https://github.com/opencontainers/runtime-spec/blob/master/runtime.md#lifecycle)

## lifecycle

### 目的
起動から終了までのlifecycleの時間が測れる

### 測定方法
n回hello-worldのプログラムを動作させる

### 測定点
1. Create
2. Start
3. Delete

## low-load

### 目的
低負荷のデーモンプロセスの時間を測る

### 測定方法
`sleep`を1度実行する

### 測定点
1. Create
2. Start
3. Kill
4. Delete

## high-load

### 目的
低負荷のデーモンプロセスの時間を測る

### 測定方法
`yes`を1度実行する

### 測定点
1. Create
2. Start
3. Kill
4. Delete