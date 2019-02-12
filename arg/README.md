# コマンドライン引数

## 実行方法

### arg.go

```
$ go run arg.go -num 100 -name hoge arg1 arg2 arge3
```

ヘルプを表示する場合
```
$ go run arg.go -h
```

### flag_ex1.go

```
$ go run flag_ex1.go [ファイル名]
```

*.goのファイルを指定するとコンパイラがコンパイル対象と間違えてしまうので -- を付けて実行する
```
$ go run flag_ex1.go -- flag_ex1.go
```

### flag_ex2.go

```
$ go run flag_ex2.go 10
```

flagを利用する場合は
```
$ go run flag_ex2.go -n 10
```
