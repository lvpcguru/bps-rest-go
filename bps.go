package bps-rest-go


import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)


const Version = "0.1.0"

// Supported HTTP verbs.
const (
	Get    Method = "GET"
	Post   Method = "POST"
	Put    Method = "PUT"
	Patch  Method = "PATCH"
	Delete Method = "DELETE"
)

// Request holds the request to an API Call.
type Request struct {
	Method      Method
	BaseURL     string 
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

// Response holds the response from an API call.
type loginResponse struct {
    apiKey string         `json:"apiKey"`
    sessionName string	  `json:"sessionName"`
    sessionId string      `json:"sessionId"`
    username string 	  `json:"username"`
    userAccountUrl string `json:"userAccountUrl"`
}



// RestError is a struct for an error handling.
type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// Error is the implementation of the error interface.
func (e *RestError) Error() string {
	return e.Response.Body
}

// Client allows modification of client headers, redirect policy
// and other settings
// See https://golang.org/pkg/net/http
type Client struct {
	HTTPClient *http.Client
}

type System struct {
	host	string
	user	string
	password string
	sessionId string
    apiKey string
    sessionName string
    sessionId string
    userAccountUrl string
}

//bps = BPS(bps_system, bpsuser, bpspass)
func BPS(bps_system,bpsuser,bpspass) (s *System, error){
    return System{
        host: bps_system,
        user: bpsuser,
        password: bpspass,
    }
}
/*
evasionProfile
reports 
capture 
network 
topology 
superflow 
testmodel 
administration 
results 
statistics 
appProfile 
strikes 
loadProfile 
strikeList 
*/
//connect to the system
func (s *System) connect() (status, error){
	//https://<System Controller IP>/bps/api/v1
	//post url='https://' + self.host + '/bps/api/v2/core/auth/logout', data=json.dumps({'username': self.user, 'password': self.password, 'sessionId': self.sessionId}), headers={'content-type': 'application/json'}, verify=False)
    req, err := http.NewRequest("POST", 'https://' + s.host + '/bps/api/v1/auth/session', nil)
    if err != nil {
        return nil, err
    }

    //add the header
    req.Header.Set("Content-Type", "application/json")
    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }

    defer res.Body.Close()

    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            return errors.New(errRes.Message)
        }

        return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    }

    //save the response example
    if err = json.NewDecoder(res.Body).Decode(&loginResponse); err != nil {
        return err
    }
    s.apiKey = loginResponse.apiKey
    s.sessionName = loginResponse.sessionName
    s.sessionId = loginResponse.sessionId
    s.userAccountUrl = loginResponse.userAccountUrl

    return nil
}

//connect to the system
func (s *System) login() (response, error){
	//https://<System Controller IP>/bps/api/v1
	//post url='https://' + self.host + '/bps/api/v2/core/auth/logout', data=json.dumps({'username': self.user, 'password': self.password, 'sessionId': self.sessionId}), headers={'content-type': 'application/json'}, verify=False)
	req, err := http.NewRequest("POST", 'https://' + s.host + '/bps/api/v2/core/auth/login', nil)

}





// AddQueryParameters adds query parameters to the URL.
func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}
