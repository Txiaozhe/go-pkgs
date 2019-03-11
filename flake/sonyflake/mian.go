package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sony/sonyflake"
)

func GetMachineId() (uint16, error) {
	var machineId uint16
	var err error

	machineId = readMachineIdFromLocal()
	if machineId == 0 {
		machineId, err = generateMachineId()
		if err != nil {
			return 0, err
		}
	}

	return machineId, nil
}

func readMachineIdFromLocal() uint16 {
	return 0
}

func generateMachineId() (uint16, error) {
	return 1, nil
}

func saddMachineIdToRedis() (int, error) {
	return 1, nil
}

func saveMachineIdToLocal() error {
	return nil
}

func checkMachineId(machineId uint16) bool {
	saddResult, err := saddMachineIdToRedis()
	if err != nil || saddResult == 0 {
		return true
	}

	err = saveMachineIdToLocal()
	if err != nil {
		return true
	}

	return false
}

func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      GetMachineId,
		CheckMachineID: checkMachineId,
	}

	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)
}
