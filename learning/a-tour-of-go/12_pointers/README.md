# pointer

## ポインタ型
メモリー上のアドレスを記憶する変数の型

## ポインタ型変数
ポインタ型で宣言された変数のこと
メモリ上のアドレスを値として入れられる変数のこと

## ポインタの使い所

- 引数やレシーバを関数内で書き換える必要がある場合
- コピーを避けたいデータを引数、レシーバにする場合
- 大きな構造体や配列を扱う場合（プリミティブな型や大きくない構造体の場合はむしろ非効率になる）
- 大きな構造体をスライスに持たせる場合
