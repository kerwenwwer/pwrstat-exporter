package gopwrstat

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

type Pwrstat struct {
	Content string            `json:"-"`
	Status  map[string]string `json:"status"`
}

type Entries map[string]string

func construct(content string) *Pwrstat {
	s := &Pwrstat{}
	s.Content = content
	s.Status = map[string]string{}
	lines := strings.Split(s.Content, "\n")
	//var status string
	var statusArr []string
	for _, line := range lines {
		if len(line) > 0 {
			line = strings.Trim(line, "	")
			line = strings.Replace(line, ". ", ";", -1)
			line = strings.Replace(line, ".", "", -1)
			newline := strings.Split(line, ";")
			if len(newline) > 1 {
				statusArr = append(statusArr, newline...)
			}
		}
	}
	for i := 0; i < len(statusArr); i += 2 {
		s.Status[statusArr[i]] = statusArr[i+1]
	}
	return s
}

// A successful call returns err == nil.
func NewFromSystem() (*Pwrstat, error) {
	out, err := exec.Command("sudo", "pwrstat", "-status").Output()
	if err != nil {
		return &Pwrstat{}, errors.New("pwrstat missing")
	}

	s := construct(string(out))

	return s, nil
}

func NewFromFile(path string) (*Pwrstat, error) {
	out, err := ioutil.ReadFile(path)
	if err != nil {
		return &Pwrstat{}, err
	}

	s := construct(string(out))
	return s, nil
}

func (s *Pwrstat) JSON() string {
	out, _ := json.Marshal(s)

	return string(out)
}

func (s *Pwrstat) String() string {
	return s.JSON()
}
