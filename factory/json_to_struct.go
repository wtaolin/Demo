package factory

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	County string `json:"county"`
}

type link struct {
	name string `json:"name"`
	url  string `json:"url"`
}

type Context struct {
	Name        string  `json:"name"`
	url         string  `json:"url"`
	page        int     `json:"page"`
	isNonProfit bool    `json:"isNonProfit"`
	address     Address `json:"address"`
	links       []link
}

//Json2Struct
func Json2Struct() {
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		log.Fatalln("read log.txt failed", err)
	}
	logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime)
	logger.Printf("prepare to write log")
	file, err := os.OpenFile("json/content.json", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		logger.Printf("open file failed err:%+v", err)
	}
	defer func() {
		logFile.Close()
		file.Close()
	}()
	Bytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Printf("read file failed err:%#v", err)
	}
	logger.Printf("read context is %+v", string(Bytes))
	str := string(Bytes)
	context := new(Context)
	logger.Printf("read context is %+v", strings.Replace(str, "Google", "google", -1))
	Bytes = bytes.TrimPrefix(Bytes, []byte("\xef\xbb\xbf"))
	if err := json.Unmarshal(Bytes, context); err != nil {
		logger.Printf("err is %#v", err)
	}
	logger.Printf("Context is %#v", context)
}
