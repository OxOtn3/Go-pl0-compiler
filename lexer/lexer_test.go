package lexer

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestExp1(t *testing.T) {
	filepath := "../assets/a.txt"
	lexer := NewLexer(filepath)
	lexer.GetSym()

	hash := make(map[string]int)
	for _, sym := range lexer.symbols {
		// 记录各标识符出现的次数
		if sym.Id.IsIdent() {
			hash[string(sym.Value)]++
		}
	}
	for k, v := range hash {
		res := "(" + k + "," + strconv.Itoa(v) + ")"
		fmt.Println(res)
	}
}

func showLexResult(t *testing.T) {
	filepath := "../assets/a.txt"
	lexer := NewLexer(filepath)
	lexer.GetSym()
	for _, sym := range lexer.symbols {
		fmt.Println(sym.String())
	}
}

func TestExp2(t *testing.T) {
	showLexResult(t)
}

func TestCreateFile(t *testing.T) {
	opath := "../output"
	_, e := os.Stat(opath)
	if e != nil && os.IsNotExist(e) {
		// 不存在output目录，创建
		err := os.Mkdir("../output", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	path := opath + "/tmp.txt"
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("this is content ..\nhuanhang"))
	if err != nil {
		panic(err)
	}
}
