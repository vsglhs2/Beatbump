package api

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	clientID            = "861556708454-d6dlm3lh05idd8npek18k6be8ba3oc68.apps.googleusercontent.com"
	clientSecret        = "SboVhoG9s0rNafixCSGGKXAT"
	scopes              = "http://gdata.youtube.com https://www.googleapis.com/auth/youtube"
	TokenURL            = "https://www.youtube.com/o/oauth2/token"
	oauth2DeviceCodeURL = "https://www.youtube.com/o/oauth2/device/code"
)

var pendingTokens = map[string]context.Context{}

type YouTubeOAuth2Handler struct {
	token oauth2.Token
}

func isTokenValid(token oauth2.Token) bool {
	return token.AccessToken != "" && token.TokenType != "" && token.RefreshToken != ""
}

func RefreshToken(refreshToken string) (oauth2.Token, error) {
	data := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return oauth2.Token{}, err
	}

	resp, err := http.Post(TokenURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return oauth2.Token{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return oauth2.Token{}, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(body, &token); err != nil {
		return oauth2.Token{}, err
	}

	if token.AccessToken == "" {
		return oauth2.Token{}, fmt.Errorf("failed to refresh access token")
	}
	values := map[string]interface{}{}
	json.Unmarshal(body, &values)
	token.RefreshToken = refreshToken
	token.Expiry = time.Now().UTC().Add(time.Second * time.Duration(values["expires_in"].(float64)))
	return token, nil
}

func DeviceOauth(c echo.Context) error {
	data := map[string]string{
		"client_id":    clientID,
		"scope":        scopes,
		"device_id":    uuid.New().String(),
		"device_model": "ytlr::",
	}

	tokenObj := extractToken(c)

	if tokenObj != nil {
		url := "https://www.googleapis.com/youtube/v3/channels?part=id&mine=true"

		// Create a Bearer string by appending string access token
		var bearer = tokenObj.TokenType + " " + tokenObj.AccessToken

		// Create a new request using http
		req, err := http.NewRequest("GET", url, nil)

		// add authorization header to the req
		req.Header.Add("Authorization", bearer)

		// Send req using http Client
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			output := map[string]string{}
			output["loggedIn"] = "true"
			return c.JSON(http.StatusOK, output)
		}
	}

	deleteTokenInCookie(c)

	payload, err := json.Marshal(data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resp, err := http.Post(oauth2DeviceCodeURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var codeResponse struct {
		VerificationURL string `json:"verification_url"`
		UserCode        string `json:"user_code"`
		DeviceCode      string `json:"device_code"`
		Interval        int    `json:"interval"`
	}
	if err := json.Unmarshal(body, &codeResponse); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	//authorize(codeResponse.DeviceCode,codeResponse.Interval)
	output := map[string]string{}
	output["oauthStart"] = fmt.Sprintf("%s?user_code=%s", codeResponse.VerificationURL, codeResponse.UserCode)
	output["deviceCode"] = codeResponse.DeviceCode
	return c.JSON(http.StatusOK, output)
}

func AuthorizeOauth(c echo.Context) error {
	deviceCode := c.Param("deviceCode")
	token, err := authorize(deviceCode, 5)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	storeTokenInCookie(c, *token)

	return c.JSON(http.StatusOK, token)
}

func storeTokenInCookie(c echo.Context, token oauth2.Token) {
	tokenJson, _ := json.Marshal(token)
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = base64.StdEncoding.EncodeToString(tokenJson)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/api"
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)
	cookie.SameSite = http.SameSiteStrictMode
	//cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
}

func deleteTokenInCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/api"
	cookie.Expires = time.Unix(0, 0)
	cookie.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie)
}

func authorize(deviceCode string, interval int) (*oauth2.Token, error) {

	if currcontext, ok := pendingTokens[deviceCode]; ok {
		if currcontext != nil {
			currcontext.Done()
		}
	}
	data := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          deviceCode,
		"grant_type":    "http://oauth.net/grant_type/device/1.0",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if interval == 0 {
		interval = 5
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	pendingTokens[deviceCode] = ctx
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer func() {
		pendingTokens[deviceCode] = nil
		ticker.Stop()
	}()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			tok, err := retrieveToken(payload)
			if err == nil {
				return tok, nil
			}
			if err.Error() == "expired_token" {
				cancel()
			}
		}
	}

	return nil, errors.New("failed to get token")
}

func retrieveToken(payload []byte) (*oauth2.Token, error) {
	resp, err := http.Post(TokenURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}

	if token.AccessToken != "" {
		values := map[string]interface{}{}
		json.Unmarshal(body, &values)
		token.Expiry = time.Now().UTC().Add(time.Second * time.Duration(values["expires_in"].(float64)))
		fmt.Println("Authorization successful")
		return &token, nil
	}

	// Handle specific OAuth errors
	var errorResponse struct {
		Error string `json:"error"`
	}
	err = json.Unmarshal(body, &errorResponse)

	if err != nil {
		return nil, err
	}

	return nil, errors.New(errorResponse.Error)
}
