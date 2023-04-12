package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"util-send-gmail/internal/utils"

	"gopkg.in/yaml.v2"
)

type FormatConfig int8

const (
	JSON FormatConfig = iota
	XML
	YAML
)

type Config struct {
	Gmail ConfigGmail `json:"gmail,omitempty" xml:"gmail,omitempty" yaml:"gmail,omitempty"`
}

func (c *Config) Init() (bool, error) {
	switch {
	case utils.IsStatFile("send-gmail.json"):
		content, err := ioutil.ReadFile("send-gmail.json")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(JSON, content); err != nil {
			return true, err
		}
	case utils.IsStatFile("send-gmail.xml"):
		content, err := ioutil.ReadFile("send-gmail.xml")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(XML, content); err != nil {
			return true, err
		}
	case utils.IsStatFile("send-gmail.yaml"):
		content, err := ioutil.ReadFile("send-gmail.yaml")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(YAML, content); err != nil {
			return true, err
		}
	}
	return false, errors.New("config file not found")
}

func (c *Config) Marshal(format FormatConfig) ([]byte, error) {
	switch format {
	case JSON:
		return json.MarshalIndent(c, "", "\t")
	case XML:
		return xml.MarshalIndent(c, "", "\t")
	case YAML:
		return yaml.Marshal(c)
	}
	return nil, errors.New("unknown configuration format")
}

func (c *Config) Unmarshal(format FormatConfig, data []byte) error {
	switch format {
	case JSON:
		return json.Unmarshal(data, c)
	case XML:
		return xml.Unmarshal(data, c)
	case YAML:
		return yaml.Unmarshal(data, c)
	}
	return errors.New("unknown configuration format")
}

func (c *Config) Example() {
	c.Gmail.Login = "you_mail@gmail.com"
	c.Gmail.Pass = "you_password"
}

func (c *Config) Check() error {
	if err := c.Gmail.Check(); err != nil {
		return fmt.Errorf("gmail: %v", err)
	}
	return nil
}

type ConfigGmail struct {
	Login string `json:"login,omitempty" xml:"login,omitempty" yaml:"login,omitempty"`
	Pass  string `json:"password,omitempty" xml:"password,omitempty" yaml:"password,omitempty"`
}

func (g *ConfigGmail) Set(login, pass string) {
	if len(login) > 1 {
		g.Login = login
	}
	if len(pass) > 1 {
		g.Pass = pass
	}
}

func (g *ConfigGmail) GetFrom() string { return g.Login }

func (g *ConfigGmail) GetPass() string { return g.Pass }

func (g *ConfigGmail) Check() error {
	if len(g.Login) < 1 {
		return errors.New("login can't be empty")
	}
	if len(g.Pass) < 1 {
		return errors.New("bot token can't be empty")
	}
	return nil
}
