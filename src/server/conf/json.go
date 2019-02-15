package conf

import (
	"io/ioutil"

	simplejson "github.com/go-simplejson"
	"github.com/name5566/leaf/log"
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
	id   string
	host string
	port string
}

var ServerMap map[string][]ServerItem = make(map[string][]ServerItem)

func init() {
	log.Debug("json init")
	data, err := ioutil.ReadFile("conf/serverConfig.json")
	if err != nil {
		log.Fatal("%v", err)
	}

	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		log.Fatal("%v", err)
	}

	m, err := js.Map()
	if err != nil {
		log.Fatal("%v", err)
	}

	for k, v := range m {
		log.Debug("%v = %v", k, v)
		var sArr = []ServerItem{}
		for _, s := range js.Get(k).MustArray() {
			eachS := s.(map[string]interface{})
			id := eachS["id"]
			log.Debug("%v", id)

			sItem := ServerItem{
				id: eachS["id"].(string),
			}
			append(sArr, sItem)
		}
		ServerMap[k] = sArr
	}
}
