package youtube

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"errors"
	"encoding/json"
)

// URK-encoded query string, 需要 decode
func GetVideoInfo(videoID string) (string, error) {
	url := "https://youtube.com/get_video_info?video_id=" + videoID
	fmt.Println("url:", url)

	// return Response and error
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// decode 用的
func ParseVideoInfo(videoInfo string) (url.Values, error) {
	answer, err := url.ParseQuery(videoInfo)
	if err != nil {
		return nil, err
	}
	
	status, ok := answer["status"]
	if !ok {
		err = fmt.Errorf("no response status found in the server's answer")
		return nil, err
	}

	if status[0] == "fail" {
		reason, ok := answer["reason"]
		if ok {
			err = fmt.Errorf("'fail' response status found in the server's answer, reason: '%s'", reason[0])
		} else {
			err = errors.New(fmt.Sprint("'fail' response status found in the server's answer, no reason given"))
		}
		return nil, err
	}

	if status[0] != "ok" {
		err = fmt.Errorf("non-success response status found in the server's answer (status: '%s')", status)
		return nil, err
	}

	return answer, err
}

func GetVideoTitleAuthor(in url.Values) (string, string) {
	playResponse, ok := in["player_response"]
	if !ok {
		return "", ""
	}
	personMap := make(map[string]interface{})

	if err := json.Unmarshal([]byte(playResponse[0]), &personMap); err != nil {
		panic(err)
	}

	s := personMap["videoDetails"]
	myMap := s.(map[string]interface{})
	// 這邊註記一下用法
	// s.() 稱為型別判定 (type assertion)
	// 因為 s 是 interface, 需要作判定才可以操作

	if title, ok := myMap["title"]; ok {
		if author, ok := myMap["author"]; ok {
			return title.(string), author.(string)
		}
	}
	return "", ""
}