package main

import (
	"fmt"
	"path"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)

/*
	TODOを殺す、コンパイル時発動のTimeAssert。

	Goのコードの中からTimeAssertの記述を見つけ出して、爆発させる。
	コンパイル時に発火する方法が無いので、モチベーションが尽きたら作るのやめる。
	
	TODOを、コンパイルが必要なコードを書かせることで、コメントで書くのとは異なるメリットがある。

	コード化とコンパイル時爆発のメリット：
		・コンパイル時にTODOに気づける
		・日付を再設定し繰り越しさせることで、延期できる
		・importさせることで、使用を明示的にできる
		
	
	デメリット：
		・リリース時に害のある内容にしてると、詰む
		・他人の仕掛けたTODOを踏んでしまうことがあり、詰む(targetを指定すれば回避可能)(だけど無差別Assertは回避できない)

*/

func main() {
	// 引数を受取って、0番に入ってるパスをもとに、其処から先に入ってる全.go拡張子のファイルに対して処理を行う。
	// のだけど、引数受取るのまだ面倒なので、(コンパイル -> 引数化)　今は現在のパスを使用する。
	currentFolder, _ := os.Getwd()

	filepath.Walk(currentFolder, Visit)
}

func Visit(filePath string, info os.FileInfo, err error) error {
	if path.Ext(filePath) == ".go" {
		b, err := ioutil.ReadFile(filePath)
		lines := strings.Split(string(b), "\n")

		for _, line := range lines {
			if strings.Contains(line, "CompiletimeAssert") {
				// そのラインを候補としてチェックする
				// "2014/07/19 5:03:14" とかの内容を保持してて、コンパイルされる予定の筈。
				// となると、コンパイル後に走るんでもいいわけだ。あ、モチベーションが消えそう。
			}
		}

		return err
	}
	
	return nil
}

