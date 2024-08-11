package api

import (
	"beatbump-server/backend/_youtube"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ArtistEndpointHandler(c echo.Context) error {
	browseId := c.Param("artistId")
	if browseId == "" {
		return c.JSON(http.StatusOK, struct{}{})
	}
	var responseBytes []byte
	var err error

	responseBytes, err = _youtube.Browse(nil, browseId, _youtube.PageType_MusicPageTypeArtist, "", nil, nil, nil, _youtube.WebMusic)

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
	}

	/*if category == "" {

		var exploreResponse _youtube.Explore
		err = json.Unmarshal(responseBytes, &exploreResponse)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
		}

		r := parseExplore(exploreResponse)
		return c.JSON(http.StatusOK, r)
	} else {*/
	var homeResponse _youtube.HomeResponse
	err = json.Unmarshal(responseBytes, &homeResponse)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error building API request: %s", err))
	}

	r := parseHome(homeResponse)

	return c.JSON(http.StatusOK, r)
	//}
}
