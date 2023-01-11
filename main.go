package main

import (
	"fmt"
	"os"
)

var (
	genshinPath  = "/storage/emulated/0/Android/data/com.miHoYo.GenshinImpact/"
	yuanshenPath = "/storage/emulated/0/Android/data/com.miHoYo.Yuanshen/"
)


func main() {
	s1, s2 := check()
	if s1 {
		moveG2Y()
	} else if s2 {
		moveY2G()
	} else {
		fmt.Println("似乎两个文件夹都是空")
	}
}

func check() (bool, bool) {
	s, er1 := os.Stat(genshinPath)
	s2, er2 := os.Stat(yuanshenPath)
	if er1 != nil {
		if er2 != nil {
			return false, false
		} else {
			return false, s2.IsDir()
		}
	} else {
		if er2 != nil {
			return s.IsDir(), false
		}
	}
	return s.IsDir(), s2.IsDir()
}
func moveG2Y() {
	os.Remove(yuanshenPath)
	err := os.Rename(genshinPath, yuanshenPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Genshin To Yuanshen")
}
func moveY2G() {
	os.Remove(genshinPath)
	err := os.Rename(yuanshenPath, genshinPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Yuanshen To Genshin")
}
