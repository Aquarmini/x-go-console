package config

import (
	"app/controller"
	"github.com/limingxinleo/di"
)

func GetTaskMap(di di.Context) map[string]controller.ControllerInterface {
	TaskMap := make(map[string]controller.ControllerInterface)

	//TaskMap["Name"] = &controller.NameController{DI:di}

	return TaskMap
}

