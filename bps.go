//Package bpsrestgo provides an API wrapper for calling BreakingPoints
//restful API
package bpsrestgo

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

//Version documents the packages current version
const Version = "0.1.0"

// LoginResponse holds the response result of an API call.
type loginResponse struct {
	ApiKey         string
	SessionName    string
	SessionID      string
	Username       string
	UserAccountURL string
}

// System defines and tracks the system connection information
type System struct {
	host           string
	User           string `json:"username"`
	Password       string `json:"password"`
	sessionID      string
	apiKey         string
	sessionName    string
	userAccountURL string
	HTTPClient     *http.Client `json:"-"`
}

//BPS initializes the system structure that is used to interact with this lib
//   it is mostly formated like this to keep usage stylse with restPy
func BPS(bpsSystem string, bpsUser string, bpsPass string) *System {
	sys := &System{
		host:       bpsSystem,
		User:       bpsUser,
		Password:   bpsPass,
		HTTPClient: http.DefaultClient,
	}

	//need to ignore invalid certs
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return sys
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
func (s *System) connect() error {
	//https://<System Controller IP>/bps/api/v1
	//post url='https://' + self.host + '/bps/api/v2/core/auth/logout', data=json.dumps({'username': self.user, 'password': self.password, 'sessionId': self.sessionId}), headers={'content-type': 'application/json'}, verify=False)

	jsonCreds, _ := json.Marshal(s)
	fmt.Println(string(jsonCreds))
	req, err := http.NewRequest("POST", "https://"+s.host+"/bps/api/v1/auth/session", bytes.NewBuffer(jsonCreds))
	if err != nil {
		return err
	}

	//add the header
	req.Header.Set("Content-Type", "application/json")
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	//save the response example
	resp := loginResponse{}
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	//if all worked out store the response
	fmt.Println(resp.ApiKey)
	s.apiKey = resp.ApiKey
	s.sessionName = resp.SessionName
	s.sessionID = resp.SessionID
	s.userAccountURL = resp.UserAccountURL

	return nil
}

func (s *System) disconnect() error {
	req, err := http.NewRequest(http.MethodDelete, "https://"+s.host+"/bps/api/v1/auth/session", nil)
	if err != nil {
		return err
	}

	//add the header
	req.Header.Set("Content-Type", "application/json")
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return nil
}

func (s *System) Login() error {

}

func (s *System) Logout() error {
	
}

func (s *System) Get() error {
	
}

func (s *System) Patch() error {
	
}

func (s *System) Set() error {
	return patch()
}

func (s *System) Put() error {
	
}

func (s *System) Delete() error {
	
}

//Options uses the provided path and returns the server response as a string
func (s *System) Options(path string) string, error {
	req, err := http.NewRequest(http.MethodOptions, "https://"+s.host+"/bps/api/v2/core/"+path, nil)
	if err != nil {
		return err
	}

	//add the header
	req.Header.Set("Content-Type", "application/json")
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return nil
}



//connect to the system
func (s *System) Login() error {
	//https://<System Controller IP>/bps/api/v1
	//post url='https://' + self.host + '/bps/api/v2/core/auth/logout', data=json.dumps({'username': self.user, 'password': self.password, 'sessionId': self.sessionId}), headers={'content-type': 'application/json'}, verify=False)
	//req, err := http.NewRequest("POST", "https://"+s.host+"/bps/api/v2/core/auth/login", nil)

	return nil
}
