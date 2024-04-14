package services

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"shareen/src/internal/models"
	"strings"
)

type SitesService struct {
	youtubeTemplate string
	rutubeTemplate  string
	vkTemplate      string
	twitchTemplate  string
}

// Site's templates for inserting
const youtubeTemplate = "<iframe width=\"560\" " +
	"height=\"315\" src=\"https://www.youtube.com/embed/\" " +
	"title=\"YouTube video player\" frameborder=\"0\" " +
	"allow=\"accelerometer; autoplay; clipboard-write; " +
	"encrypted-media; gyroscope; picture-in-picture; " +
	"\"web-share\" allowfullscreen></iframe>"
const rutubeTemplate = "<iframe width=\"720\" height=\"405\" " +
	"src=\"https://rutube.ru/play/embed/\" frameBorder=\"0\" allow=\"clipboard-write; autoplay\" " +
	"webkitAllowFullScreen mozallowfullscreen allowFullScreen></iframe>"
const vkTemplate = "<iframe src=\"https://vk.com/video_ext.php?oid=-\" width=\"853\" height=\"480\" " +
	"allow=\"autoplay; encrypted-media; fullscreen; picture-in-picture;\" frameborder=\"0\" " +
	"allowfullscreen></iframe>"
const twitchTemplate = "<div id=\"twitch-embed\"></div>" +
	"<script src=\"https://player.twitch.tv/js/embed/v1.js\"></script>" +
	"<script type=\"text/javascript\">" +
	"new Twitch.Player(\"twitch-embed\", {channel: \"\"});</script>"

func NewSitesService() *SitesService {
	return &SitesService{
		youtubeTemplate: youtubeTemplate,
		rutubeTemplate:  rutubeTemplate,
		vkTemplate:      vkTemplate,
		twitchTemplate:  twitchTemplate,
	}
}

func (ss *SitesService) GetSiteURL(URL string) (string, *models.ResponseError) {
	var siteURL string
	if strings.Contains(URL, "youtube") {
		siteURL = getYoutubeURL(URL)
		log.Println(siteURL)
		return siteURL, nil
	}

	if strings.Contains(URL, "rutube") {
		siteURL = getRutubeURL(URL)
		log.Println(siteURL)
		return siteURL, nil
	}

	if strings.Contains(URL, "twitch") {
		siteURL = getTwitchURL(URL)
		log.Println(siteURL)
		return siteURL, nil
	}

	if strings.Contains(URL, "vk.com") {
		siteURL := getVkURL(URL)
		log.Println(siteURL)
		return siteURL, nil
	}

	return siteURL, &models.ResponseError{
		Message: fmt.Sprintf("Сайт: {%s}, не поддерживается", URL),
		Status:  http.StatusBadRequest,
	}
}

func getRutubeURL(URL string) string {
	rutubeVideoIndex := strings.Index(URL, "video") + 6
	templateEmbedIndex := strings.Index(youtubeTemplate, "embed") + 5
	readyURL := rutubeTemplate[:templateEmbedIndex] + URL[rutubeVideoIndex:len(URL)-1] + rutubeTemplate[templateEmbedIndex:]
	return readyURL
}

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
