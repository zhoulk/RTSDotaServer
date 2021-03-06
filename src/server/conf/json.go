package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"

	simplejson "github.com/go-simplejson"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

type ServerItem struct {
	Id         string
	TCPAddr    string
	MaxConnNum int
}

var ServerMap map[string]*ServerItem = make(map[string]*ServerItem)

func init() {
	data, err := ioutil.ReadFile("conf/serverList.json")
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

	for _, v := range m {
		// log.Debug("%v = %v", k, v)
		var item *ServerItem
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &item)
		if err != nil {
			log.Debug("err was %v", err)
		}
		ServerMap[item.Id] = item
	}

	// Server.MaxConnNum = 20000
	// Server.TCPAddr = ServerMap["Gate"][0].host + ":" + ServerMap["Gate"][0].port
}
