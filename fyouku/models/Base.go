package models

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

type Base struct {
	Id   int
	Name string
}

type BaseApiData struct {
	Code  int
	Msg   string
	Items []Base
	Count string
}

// 获取频道地区
func GetChannelRegion(channelId int) []Base {
	var info []Base
	req := httplib.Get(getApiURL() + "/channel/region?channelId=" + strconv.Itoa(channelId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &BaseApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

// 获取频道类型
func GetChannelType(channelId int) []Base {
	var info []Base
	req := httplib.Get(getApiURL() + "/channel/type?channelId=" + strconv.Itoa(channelId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &BaseApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}
