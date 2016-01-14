package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"github.com/aws/aws-sdk-go/aws"
	"github.com/codegangsta/cli"
)

type EC2Instance struct {
	InstanceType   string   `json:"instance_type"`
	ImageId        string   `json:"image_id"`
	SecurityGroups []string `json:"security_groups"`
	Id             string
	Tags           []map[string]string
}

func main() {
	app := cli.NewApp()
	app.Name = "uranus"
	app.Usage = "cloud builder"

	dat, err := ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	var ec2_instance EC2Instance

	err = json.Unmarshal(dat, &ec2_instance)
	if err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}

	fmt.Print(ec2_instance)
}
