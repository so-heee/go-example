# まとめ

## 命名規約

### パッケージ名

- 全て小文字
- 1 単語で構成されるのが望ましい
- 複数単語の場合はそのままつなげる？（Kubernetesがそうなってた）

### ファイル名

- パッケージと同じ
- 複数単語の場合は_で区切る
- 「.」で始まるファイル名 または「_」はgoツールによって無視される
- 接尾辞「_test.go」が付いたファイルは、go testツールによってのみコンパイルおよび実行される

### 関数とメソッド

- キャメルケース（ローワーキャメルケース）の場合は、外部のパッケージから参照不可
- パスカルケース（アッパーキャメルケース）の場合は、外部のパッケージから参照可能
- レシーバは一文字で統一

### 定数

- 関数と同じ規約
- 全ての文字を大文字にするなども見かけるがKubernetesなどのコードを見ると違う

### 変数

- なるべく短くする（Config → c, conf, cfg）
- ソースコード全体で一貫した命名スタイルを使用する

### 型

- 関数と同じ

### インターフェース

- パスカルケース


## 内容

### variables






