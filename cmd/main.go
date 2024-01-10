package main

import (
	"fmt"
	"smartNotify/internal/structures"
	"time"
)

func main() {

	t0 := structures.NewTask(1, "Report 1", time.Now().Add(time.Hour*1))
	t1 := structures.NewTask(2, "Report 2", time.Now().Add(time.Hour*1))
	t2 := structures.NewTask(3, "Report 3", time.Now().Add(time.Hour*1))
	t3 := structures.NewTask(4, "Report 4", time.Now().Add(time.Hour*1))

	root, child1 := &structures.Container{}, &structures.Container{}

	root.AddTask(*t0)
	root.AddTask(*t1)

	child1.AddTask(*t2)
	child1.AddTask(*t3)

	root.AddChild(child1)

	agent := &structures.Agent{}
	agent.MoveToUndone(root)

	firstUnDoneContainer := agent.GetFirstUndoneContainer()

	fmt.Printf("%+v\n", firstUnDoneContainer)
}
