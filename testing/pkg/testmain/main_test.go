package testmain

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// TestMainがフラグを使用する場合は、ここでflag.Parse（）を呼び出します
	// セットアップコード
	log.Print("setup")
	// パッケージ内のテストの実行
	code := m.Run()
	// ティアダウンコード
	log.Print("tear-down")
	// テストの終了コードで exit
	os.Exit(code)
}
