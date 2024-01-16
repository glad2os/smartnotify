package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"smartNotify/internal/config"
	"smartNotify/internal/structures"
	"smartNotify/internal/utils"
	"sort"
	"time"
)

func main() {
	config := config.Config{}
	file, err := os.ReadFile("./configs/schedule.yaml")

	if err != nil {
		log.Fatalln("Error reading filename")
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln("Error unmarshalling YAML")
	}

	scheduleToContainer := make(map[string]*structures.Container)
	var i = 0

	for _, user := range config.Users {
		for _, notification := range user.Notifications {
			if _, exists := scheduleToContainer[notification.Trigger.Schedule]; !exists {
				scheduleToContainer[notification.Trigger.Schedule] = &structures.Container{}
			}
			task := structures.NewTask(i, notification.Message, structures.PENDING)
			i += 1
			scheduleToContainer[notification.Trigger.Schedule].AddTask(*task)
		}
	}

	var sortedContainers []*structures.Container
	for schedule, container := range scheduleToContainer {
		nextRun, err := utils.GetNextRunTime(schedule)
		if err != nil {
			log.Fatalf("Error parsing CRON expression: %v", err)
		}
		container.DueDate = nextRun.Format(time.RFC3339) // Store as a sortable string
		sortedContainers = append(sortedContainers, container)
	}

	// Sort by DueDate
	sort.Slice(sortedContainers, func(i, j int) bool {
		return sortedContainers[i].DueDate < sortedContainers[j].DueDate
	})

	// Linking containers
	for i, parentContainer := range sortedContainers {
		if i+1 < len(sortedContainers) {
			childContainer := sortedContainers[i+1]
			parentContainer.AddChild(childContainer)
		}
	}

	agent := &structures.Agent{}

	for _, container := range sortedContainers {
		agent.MoveToUndone(container)
	}

	fmt.Printf("%+v\n", agent.GetFirstUndoneContainer())
}
