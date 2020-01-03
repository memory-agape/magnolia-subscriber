package controller

import (
	"lupinus/config"
	"lupinus/servers/http"
	"lupinus/servers/http/web/behavior"
	"lupinus/share"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func RequestFavorite(clientMeta http.HttpClientMeta) (*http.HttpBody, *http.HttpHeader) {
	session := behavior.GetSignInInfo(clientMeta)

	if session == nil {
		// Not exists a session
		return &http.HttpBody{
				Payload: map[string]interface{}{
					"message": "Unauthorized",
				},
			},
			&http.HttpHeader{
				Status: 401,
			}
	}

	body := &http.HttpBody{}
	header := &http.HttpHeader{}
	switch clientMeta.Method {
	case "GET":
		body, header = requestFavoriteByGet(clientMeta)
		break
	case "POST":
		body, header = requestFavoriteByPost(clientMeta)
		break
	default:
		body = nil
		header = nil
	}
	return body, header
}

func requestFavoriteByGet(clientMeta http.HttpClientMeta) (*http.HttpBody, *http.HttpHeader) {
	session := behavior.GetSignInInfo(clientMeta)

	files, _ := filepath.Glob(
		config.GetRootDir() + "/storage/" + session.Data["id"].(string) + "/*/*.jpg",
	)


	dates := map[string]interface{}{}

	for _, file := range files {
		unixTime, _ := strconv.Atoi(
			strings.Replace(
				filepath.Base(file),
				".jpg",
				"",
				-1),
		)

		date := time.Unix(int64(unixTime), 0)
		id := date.Format("20060102")

		type dateObject = []map[string]interface{}
		if _, ok := dates[id]; !ok {
			dates[id] = dateObject{}
		}
		dates[id] = append(
			dates[id].(dateObject),
			map[string]interface{}{
				"src": "image?id=" + strconv.Itoa(unixTime),
			},
		)
	}

	return &http.HttpBody{
		Payload: map[string]interface{}{
			"status": 200,
			"dates": dates,
		},
	},
	&http.HttpHeader{
		Status: 200,
	}
}

func requestFavoriteByPost(clientMeta http.HttpClientMeta) (*http.HttpBody, *http.HttpHeader) {
	share.AddProcedure(share.Procedure{
		Callback: func(data string) {

		},
	})
	return &http.HttpBody{
		Payload: map[string]interface{}{
			"status": 200,
		},
	},
	&http.HttpHeader{
		Status: 200,
	}
}
