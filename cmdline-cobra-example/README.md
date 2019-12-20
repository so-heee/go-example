# cmdline-cobra-example

## 作成手順

- 1.Go module初期化
```
go mod init github.com/so-hee/cmdline-cobra-example
```

- 2.Cobraのインストール
```
go get -u github.com/spf13/cobra/cobra
```

- 3.Cobra初期化
```
cobra init --pkg-name github.com/so-heee/golang-example/cmdline-cobra-example

.
├── LICENSE
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
```

- 4.RootCmdにprint処理実装
```
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("HelloWorld!!")
  },
```

- 5.サブコマンドの追加
```
cobra add hoge

.
├── LICENSE
├── cmd
│   ├── hoge.go
│   └── root.go
├── go.mod
├── go.sum
└── main.go
```
