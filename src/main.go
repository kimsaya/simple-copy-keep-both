package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	dir := "C:/Users/AI/Desktop/211820/"
	des := dir + "temps/"
	mods, err := ioutil.ReadDir(dir)
	log.Println(err)
	for _, mod := range mods {
		if mod.IsDir() {
			if strings.Contains(mod.Name(), "temps") {
				continue
			}
			modPart := dir + mod.Name()
			items, err := ioutil.ReadDir(modPart)
			log.Println(err)
			for _, item := range items {
				itemPath := modPart + "/" + item.Name()
				savePath := des + mod.Name() + ".pak"
				if !item.IsDir() {
					source, err := os.Open(itemPath)
					if err != nil {
						log.Println(err)
					}
					defer source.Close()
					destination, err := os.Create(savePath)
					if err != nil {
						log.Println(err)
					}
					defer destination.Close()
					_, err = io.Copy(destination, source)
					if err != nil {
						log.Println(err)
					}
				}

			}
		}
	}
}
