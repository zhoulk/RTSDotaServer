package entry

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-simplejson"
	"github.com/name5566/leaf/log"
)

type Zone struct {
	Id         string
	TCPAddr    string
	MaxConnNum int32
	Name       string
	IsNew      bool
}

var ZoneList []*Zone = make([]*Zone, 0)

func init() {
	data, err := ioutil.ReadFile("data/json/zone.json")
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
		var z *Zone
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Debug("err was %v", err)
		}
		err = json.Unmarshal(bytes, &z)
		if err != nil {
			log.Debug("err was %v", err)
		}
		ZoneList = append(ZoneList, z)
	}

	// Server.MaxConnNum = 20000
	// Server.TCPAddr = ServerMap["Gate"][0].host + ":" + ServerMap["Gate"][0].port
}
