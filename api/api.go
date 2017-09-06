package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type point struct {
	ID    string  `json:"id"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
	Total int     `json:"total"`
	Link  string  `json:"link"`
	Text  string  `json:"text"`
}

// WebsocketHandler -- Connect with client
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		getPhotos(string(msg), conn, msgType)
	}
}

func getPhotos(username string, conn *websocket.Conn, msgType int) {
	var photos []interface{}
	param := ""
	for {
		var data map[string]interface{}
		url := fmt.Sprintf("https://www.instagram.com/%s/media%s", username, param)
		resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal([]byte(bytes), &data)
		items := data["items"].([]interface{})
		for i := 0; i < len(items); i++ {
			_, err := items[i].(map[string]interface{})["location"].(map[string]interface{})
			if err != false {
				photos = append(photos, items[i])
			}
		}

		findMore := data["more_available"].(bool)
		if findMore != true {
			break
		} else {
			id := items[len(items)-1].(map[string]interface{})["id"].(string)
			param = fmt.Sprintf("?max_id=%v", id)
		}
	}
	points := make(chan point)
	for i := 0; i < len(photos); i++ {
		go generatePoint(photos[i].(map[string]interface{}), points, len(photos))
	}
	var found []point
	for p := range points {
		found = append(found, p)
		point, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = conn.WriteMessage(msgType, []byte(string(point)))
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(found) == len(photos) {
			err = conn.WriteMessage(msgType, []byte("end"))
			if err != nil {
				fmt.Println(err)
				return
			}
			close(points)
		}
	}
}

func generatePoint(photo map[string]interface{}, points chan point, total int) {
	url := fmt.Sprintf("https://www.instagram.com/p/%s/", photo["code"].(string))
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	txt := string(bytes)
	r := regexp.MustCompile(`"id": "(\d*)", "has_public_page": true`)
	locationID := r.FindAllStringSubmatch(txt, -1)[0][1]
	url = fmt.Sprintf("https://www.instagram.com/explore/locations/%s/", locationID)
	resp, _ = http.Get(url)
	bytes, _ = ioutil.ReadAll(resp.Body)
	txt = string(bytes)
	rLong := regexp.MustCompile(`longitude" content="(.*)"`)
	rLat := regexp.MustCompile(`latitude" content="(.*)"`)
	lat, _ := strconv.ParseFloat(rLat.FindAllStringSubmatch(txt, -1)[0][1], 64)
	long, _ := strconv.ParseFloat(rLong.FindAllStringSubmatch(txt, -1)[0][1], 64)
	link := photo["images"].(map[string]interface{})["standard_resolution"].(map[string]interface{})["url"].(string)
	text := photo["caption"].(map[string]interface{})["text"].(string)
	points <- point{photo["code"].(string), lat, long, total, link, text}
}
