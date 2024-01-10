package structures

type IAgent interface {
	goNext(container Container)
	goPrev(container Container)
	MoveToDone(doneContainer *Container)
	MoveToUndone(unDoneTasks *Container)
	GetFirstUndoneContainer()
}

type Agent struct {
	UnDoneTasks []*Container
	DoneTasks   []*Container
}

func (Agent *Agent) MoveToDone(container *Container) {
	Agent.DoneTasks = append(Agent.DoneTasks, container)
}

func (Agent *Agent) MoveToUndone(container *Container) {
	Agent.UnDoneTasks = append(Agent.UnDoneTasks, container)
}

func (agent *Agent) GetFirstUndoneContainer() []Task {
	if len(agent.UnDoneTasks) > 0 {
		return agent.UnDoneTasks[0].Tasks
	}

	return nil
}
