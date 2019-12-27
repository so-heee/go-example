# firebase-admin-sdk-cmd

# Description
firebase/firebase-admin-goを使用したコマンドラインツールです。
FirebaseAuthenticationに関する操作が可能です。

# Setup

## Setting config
$HOME/.firebase-admin-go-cmd.yamlを設置するか、yamlを作成して--configフラグで指定してください。

```
firebase:
    credential-file: "./config/XXXXXXXXX.json"

debug: false
```

## Setting firebase credential file
firebase credentials jsonを.firebase-admin-go-cmd.yamlで指定したディレクトリに設置してください


# Usage

## Global Flags
```
--config string   config file (default is $HOME/.firebase-admin-go-cmd.yaml)
```

## Get
```
firebase-admin-go-cmd get [flags]

Flags:
  -h, --help        help for get
  -i, --id string   Get user UID
  -l, --list        Get user list
```

## Create
```
firebase-admin-go-cmd create [flags]

Flags:
  -d, --disabled          Create user Disabled
  -e, --email string      Create user Email
  -h, --help              help for create
  -i, --id string         Create user UID
  -n, --name string       Create user Name
  -w, --password string   Create user Password
  -p, --phone string      Create user Phone Number
  -u, --url string        Create user Photo Url
```

## Update
```
firebase-admin-go-cmd update [flags]

Flags:
  -d, --disabled          Create user Disabled
  -e, --email string      Create user Email
  -h, --help              help for update
  -i, --id string         Create user UID
  -n, --name string       Create user Name
  -w, --password string   Create user Password
  -p, --phone string      Create user Phone Number
  -u, --url string        Create user Photo Url
```

## Delete
```
firebase-admin-go-cmd delete [flags]

Flags:
  -h, --help        help for delete
  -i, --id string   Delete user UID
```
