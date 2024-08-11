package api

import (
	"beatbump-server/backend/_youtube"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PlayerAPIResponse struct {
	// Define your response structure based on the actual data fields
}

func PlayerEndpointHandler(c echo.Context) error {
	url := c.Request().URL
	query := url.Query()
	videoId := query.Get("videoId")
	playlistId := query.Get("playlistId")
	//playerParams := query.Get("playerParams")

	//authorization := headers.Get("Authorization")
	if videoId == "" {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Missing required param: videoId"))
	}

	tokenObj := extractToken(c)

	playerResponse, err := callPlayerAPI(tokenObj, _youtube.IOS, videoId, playlistId)

	//try IOS_MUSIC as fallback
	if err != nil {
		playerResponse, err = callPlayerAPI(tokenObj, _youtube.WebMusic, videoId, playlistId)
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, playerResponse)
}

func callPlayerAPI(token *oauth2.Token, clientInfo _youtube.ClientInfo, videoId string, playlistId string) (map[string]interface{}, error) {

	responseBytes, err := _youtube.Player(token, videoId, playlistId, clientInfo, nil)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error building API request: %s", err))
	}

	var playerResponse map[string]interface{}
	err = json.Unmarshal(responseBytes, &playerResponse)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error building API request: %s", err))
	}

	playabilityStatus := playerResponse["playabilityStatus"]
	if playabilityStatus != nil && playabilityStatus != "" {
		playabilityStatusMap := playabilityStatus.(map[string]interface{})

		if playabilityStatusMap["status"] != "OK" {
			return nil, errors.New(fmt.Sprintf("Error building API request: %s", playabilityStatusMap["status"]))
		}
	}
	return playerResponse, nil
}
