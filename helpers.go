package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/knadh/go-get-youtube/youtube"
)

func GetChannelVideoData(channelId string) (string, string) {
	requestURL := API_ENDPOINT_ID + channelId + "&" + MAX_RESULTS + "&" + ORDER_BY + "&" + TYPE + "&key=" + API_KEY

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var video ChannelBody

	json.Unmarshal([]byte(string(body)), &video)

	return video.Items[0].ID.VideoID, video.Items[0].Snippet.Title
}

func GetUserUploadsID(channelName string) string {
	requestURL := API_ENDPOINT_NAME + channelName + "&key=" + API_KEY

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user UserInformation

	json.Unmarshal([]byte(string(body)), &user)

	return user.Items[0].ContentDetails.RelatedPlaylists.Uploads
}

func GetUserVideoData(uploadsId string) (string, string) {
	requestURL := API_ENDPOINT_PLAYLIST + uploadsId + "&maxResults=2" + "&key=" + API_KEY

	resp, err := http.Get(requestURL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var video NameBody

	json.Unmarshal([]byte(string(body)), &video)

	return video.Items[1].Snippet.ResourceID.VideoID, video.Items[1].Snippet.Title
}

func DownloadVideoAndAudio(videoID, videoTitle string) {
	video, err := youtube.Get(videoID)
	if err != nil {
		log.Panic(err)
	}

	option := &youtube.Option{
		Rename: false,
		Resume: true,
		Mp3:    true,
	}
	video.Download(0, videoTitle+".mp4", option)
}