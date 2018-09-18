//メソッド名、パッケージ名は任意で良い
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//戻り値が２つ
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	//_はrange演算子からの戻り値（index）を使用しないという意味。使わないと宣言しないとエラーとなる。
	for _, file := range files {
		fmt.Println(file.Name())
	}
}