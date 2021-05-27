package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/dstotijn/go-notion"
	"gopkg.in/yaml.v2"
)

var configPath = flag.String("config", "config.yml", "Path to config file in yaml format")

type Config struct {
	PageID string `yaml:"page_id"`
	Token  string `yaml:"token"`
}

func parseConfig(path string) (*Config, error) {
	var config Config

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	flag.Parse()

	config, err := parseConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	today := time.Now()

	client := notion.NewClient(config.Token)

	db, err := client.FindDatabaseByID(context.Background(), config.PageID)
	if err != nil {
		log.Fatal(err)
	}

	dbPageProperties := notion.DatabasePageProperties{
		"Title": notion.DatabasePageProperty{Title: []notion.RichText{{Text: &notion.Text{Content: today.Weekday().String()}}}},
	}

	if _, ok := db.Properties["Date"]; ok {
		dbPageProperties["Date"] = notion.DatabasePageProperty{Date: &notion.Date{Start: notion.NewDateTime(today, false)}}
	}

	_, err = client.CreatePage(context.Background(),
		notion.CreatePageParams{
			ParentType:             notion.ParentTypeDatabase,
			ParentID:               config.PageID,
			DatabasePageProperties: &dbPageProperties})
	if err != nil {
		log.Fatal(err)
	}
}
