package procapgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type Task struct {
	Id      string `json:"ID"`
	Time    int    `json:"Time"`
	Message string `json:"Message"`
	Success bool   `json:"Success"`
	Results struct {
		Pass         string `json:"Pass"`
		ChallengeKey string `json:"ChallengeKey"`
	} `json:"Results"`
}

type Options struct {
	RawUrl    string
	Sitekey   string
	Proxy     string
	UserAgent string
	Rqdata    string
	Apikey    string
}

type User struct {
	DailyLimit     int     `json:"daily_limit"`
	DailyReset     int64   `json:"next_reset"`
	DailyUsed      int     `json:"daily_used"`
	DailyRemaining int     `json:"daily_remaining"`
	Funds          float64 `json:"balance"`
	PlanExpire     int64   `json:"plan_expire"`
}

func CreateTask(o Options) (Task, error) {
	if o.Apikey == "" && o.Sitekey == "" && o.RawUrl == "" {
		return Task{}, errors.New("missing required parameters (apikey, sitekey, raw url)")
	}
	if !strings.Contains(o.RawUrl, "://") {
		o.RawUrl = "https://" + o.RawUrl
	}
	if !strings.Contains(o.Proxy, "://") && o.Proxy != "" {
		o.Proxy = "http://" + o.Proxy
	}
	payload := map[string]string{
		"url":       o.RawUrl,
		"sitekey":   o.Sitekey,
		"proxy":     o.Proxy,
		"userAgent": o.UserAgent,
		"rqdata":    o.Rqdata,
	}
	p, err := json.Marshal(payload)
	if err != nil {
		return Task{}, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.procap.wtf/createTask", bytes.NewReader(p))
	if err != nil {
		return Task{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", o.Apikey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Task{}, err
	}
	var t Task
	err = json.Unmarshal(b, &t)
	if err != nil {
		return Task{}, err
	}
	if !t.Success {
		return Task{}, errors.New(t.Message)
	}
	return t, nil
}

func CheckTask(t Task) (Task, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.procap.wtf/checkTask/"+t.Id, nil)
	if err != nil {
		return Task{}, err
	}
	req.Header.Set("apikey", t.Results.ChallengeKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Task{}, err
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return Task{}, err
	}

	return t, nil
}

func Solve(o Options) (string, string, error) {
	t, err := CreateTask(o)
	if err != nil {
		return "", "", err
	}
	start := time.Now()
	for {
		t, err = CheckTask(t)
		if t.Results.Pass != "" {
			return t.Results.Pass, t.Results.ChallengeKey, nil
		}
		if err != nil {
			return "", "", err
		}
		if !t.Success {
			return "", "", errors.New(t.Message)
		}
		if time.Since(start) > time.Minute*2 {
			return "", "", errors.New("solving timed out")
		}
		time.Sleep(time.Second)
	}
}

func GetUser(apikey string) (User, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.procap.wtf/user", nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Set("apikey", apikey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
