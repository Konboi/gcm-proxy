package gcm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Config struct {
	Port   int
	APIKey string
}

type Proxy struct {
	cfg *Config
}

type Payload struct {
	RegistrationIds []string `json:"registration_ids"`
	Data            *Data    `json:"data"`
}

type Data struct {
	Message string `json:"message"`
}

var Endpoint string = "https://android.googleapis.com/gcm/send"
var proxy *Proxy

func NewProxy(config *Config) (*Proxy, error) {
	if config.Port == 0 {
		return nil, errors.New("Please Set Port")
	}
	if config.APIKey == "" {
		return nil, errors.New("Please Set APIKey")
	}

	proxy = &Proxy{
		cfg: config,
	}

	return proxy, nil
}

func (p *Proxy) Run() {
	fmt.Printf("gcm-proxy start 0.0.0.0:%d \n", p.cfg.Port)
	http.HandleFunc("/", Reciver)
	http.ListenAndServe(":"+strconv.Itoa(p.cfg.Port), nil)
}

func Reciver(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Error(w, "Method Not Allowd", http.StatusMethodNotAllowed)
		return
	}

	if req.Body == nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()
	err := req.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	// TODO
	// when header json
	if req.Form.Get("token") == "" {
		log.Println("Parameter [token] is empty")
		http.Error(w, "Lack Parameter", http.StatusBadRequest)
		return
	}

	if req.Form.Get("alert") == "" {
		log.Println("Parameter [alert] is empty")
		http.Error(w, "Lack Parameter", http.StatusBadRequest)
		return
	}

	tokens := strings.Split(req.Form.Get("token"), ",")
	alert := req.Form.Get("alert")

	// TODO
	// use gorutine
	data := &Data{Message: alert}
	payload := &Payload{
		RegistrationIds: tokens,
		Data:            data,
	}
	Send(payload)

	w.WriteHeader(http.StatusOK)
}

func Send(payload *Payload) {
	p, err := json.Marshal(payload)
	body := strings.NewReader(string(p))

	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", Endpoint, body)
	if err != nil {
		log.Fatal(err)
	}

	apiKey := getAPIKey()

	req.Header.Set("Authorization", "key="+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Post GCM Error: %s", err.Error())
	} else if !strings.Contains(resp.Status, "200") {
		log.Printf("Post GCM Error: %s", resp.Status)
	} else {
		log.Print("Post GCM Success")
	}
}

// for testing method
func getAPIKey() string {
	var apiKey string

	if proxy == nil {
		apiKey = "test api key"
	} else {
		apiKey = proxy.cfg.APIKey
	}

	return apiKey
}