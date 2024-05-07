package lib

import (
	"fmt"
	"regexp"
	"strings"
)

// Site's templates for inserting
const (
	youtubeTemplate = "<iframe width=\"560\" " +
		"height=\"315\" src=\"https://www.youtube.com/embed/\" " +
		"title=\"YouTube video player\" frameborder=\"0\" " +
		"allow=\"accelerometer; autoplay; clipboard-write; " +
		"encrypted-media; gyroscope; picture-in-picture; " +
		"\"web-share\" allowfullscreen></iframe>"
	vkTemplate = "<iframe src=\"https://vk.com/video_ext.php?oid=-\" width=\"853\" height=\"480\" " +
		"allow=\"autoplay; encrypted-media; fullscreen; picture-in-picture;\" frameborder=\"0\" " +
		"allowfullscreen></iframe>"
	twitchTemplate = "<div id=\"twitch-embed\"></div>" +
		"<script src=\"https://player.twitch.tv/js/embed/v1.js\"></script>" +
		"<script type=\"text/javascript\">" +
		"new Twitch.Player(\"twitch-embed\", {channel: \"\"});</script>"
)

// const rutubeTemplate = "<iframe width=\"720\" height=\"405\" " +
//
//	"src=\"https://rutube.ru/play/embed/\" frameBorder=\"0\" allow=\"clipboard-write; autoplay\" " +
//	"webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>"

func GetIframe(URL string) (string, error) {
	var siteURL string
	if strings.Contains(URL, "youtube") {
		siteURL = getYoutubeURL(URL)
	} else if strings.Contains(URL, "twitch") {
		siteURL = getTwitchURL(URL)
	} else if strings.Contains(URL, "vk.com") {
		siteURL = getVkURL(URL)
	}

	if siteURL != "" {
		fmt.Println(siteURL)
		return siteURL, nil
	}

	return "", fmt.Errorf("unknown site")
}

// func getRutubeURL(URL string) string {
// 	rutubeVideoIndex := strings.Index(URL, "video") + 6
// 	templateEmbedIndex := strings.Index(youtubeTemplate, "embed") + 5
// 	readyURL := rutubeTemplate[:templateEmbedIndex] + URL[rutubeVideoIndex:len(URL)-1] + rutubeTemplate[templateEmbedIndex:]
// 	return readyURL
// }

func getYoutubeURL(URL string) string {
	youtubeCodeIndex := strings.Index(URL, "=") + 1 //Получение кода из URL
	youtubeEmbedIndex := strings.Index(youtubeTemplate, "embed") + 6
	youtubeCode := URL[youtubeCodeIndex:]
	readyURL := youtubeTemplate[:youtubeEmbedIndex] + youtubeCode + youtubeTemplate[youtubeEmbedIndex:]
	return readyURL
}

func getTwitchURL(URL string) string {
	twitchIndexTemplate := strings.Index(twitchTemplate, "\"\"") + 1
	urlIndex := strings.LastIndex(URL, "/") + 1
	readyURL := twitchTemplate[:twitchIndexTemplate] + URL[urlIndex:] + twitchTemplate[twitchIndexTemplate:]
	return readyURL
}

func getVkURL(URL string) string {
	regex, _ := regexp.Compile(`\d*[0-9]_\d*[0-9]`)
	rawVideoVkID := regex.FindString(URL)
	indexToSplit := strings.Index(rawVideoVkID, "_")
	indexOfTemplate := strings.Index(vkTemplate, "=-") + 2
	readyVideoVkId := rawVideoVkID[:indexToSplit] + "&id=" + rawVideoVkID[indexToSplit+1:] + "&hd=2"
	readyURL := vkTemplate[:indexOfTemplate] + readyVideoVkId + vkTemplate[indexOfTemplate:]
	return readyURL
}
