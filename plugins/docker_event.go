package plugins

import (
	"encoding/json"
	"io"
	"log"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/events"
	"github.com/docker/engine-api/types/filters"
	"golang.org/x/net/context"

	//"node/common"
)

const (
	ENDPOINT = "http://127.0.0.1:2375"
	VERSION  = "v1.24"
)

type EventHandler struct {
	//plugin common.BasePlugin
}

func (handler *EventHandler) Init() error {
	log.Println("EventHandler init")
	go handler.HandleStartContainer()
	return nil
}

func (handler *EventHandler) GetPluginName() string {
	return "Docker Event Handler"
}

func (handler *EventHandler) HandleStartContainer() {

	cli, err := client.NewClient(ENDPOINT, VERSION, nil, nil)
	if err != nil {
		panic(err)
	}
	args := filters.NewArgs()
	args.Add("event", "start")
	eventOpts := types.EventsOptions{Filters: args}
	body, err := cli.Events(context.Background(), eventOpts)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(body)
	for {
		var event events.Message
		err := dec.Decode(&event)
		if err != nil && err == io.EOF {
			log.Println("1")
			log.Println(err)
			break
		}
		log.Println(event)
		c_json, err := cli.ContainerInspect(context.Background(), event.ID)
		if err != nil {
			log.Println("2")
			log.Println(err)
			break
		}
		log.Println(c_json)
		log.Println(c_json.Mounts)
	}
}
