package entry

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-simplejson"
	"github.com/name5566/leaf/log"
)

var (
	playerExpList = make([]int32, 0)
	heroExpList   = make([]int32, 0)
)

type Exp struct {
	Player []int32
	Hero   []int32
}

func init() {
	InitExpList()
}

func InitExpList() {
	m, _ := ReadFile("data/json/exp.json")

	for _, v := range m {
		var exp Exp
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &exp)
		if err != nil {
			log.Debug("err was %v", err)
		}

		playerExpList = exp.Player
		heroExpList = exp.Hero
	}

	log.Debug("[playerExpList ] %v", playerExpList)
	log.Debug("[heroExpList ] %v", heroExpList)
}

func ReadFile(path string) ([]interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("%v", err)
	}

	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		log.Fatal("%v", err)
	}

	m, err := js.Array()
	if err != nil {
		log.Fatal("%v", err)
	}

	return m, err
}
