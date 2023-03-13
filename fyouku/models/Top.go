package models

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
	"strconv"
)

// GetChannelTop 获取动漫排行榜
func GetChannelTop(channelId int) []Video {
	var info []Video

	req := httplib.Get(getApiURL() + "/channel/top?channelId=" + strconv.Itoa(channelId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

// GetTypeTop 获取少女排行榜
func GetTypeTop(typeId int) []Video {
	var info []Video
	req := httplib.Get(getApiURL() + "/type/top?typeId=" + strconv.Itoa(typeId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}
