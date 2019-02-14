# JSON Formatterコマンド

## 実行方法

```
$ go run main.go
```

- `./testdata/test.json` にテスト用のjsonが含まれる

## 実装のヒント

### 標準パッケージ

以下のメソッドの公式ドキュメントを参考にするとよい

#### jsonパッケージ

- JSONのpretty print: `json.Indent()`
- JSONのminify: `json.Compact()`

#### stringsパッケージ

- 特定の文字列を繰り返す: `strings.Repeat()`

### 利用する型

- `bytes.Buffer` 構造体
    - 宣言: `buf := bytes.Buffer{}`
    - 文字列(`string`)の出力: `buf.String()`
