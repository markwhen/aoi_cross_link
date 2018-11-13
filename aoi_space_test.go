package crosslink

import (
	"fmt"
	"testing"
)

type gameEntity struct {
	x float32
	z float32

	id uint32
}

func (thisEnt *gameEntity) AoiCLX() CLPosValType {
	return CLPosValType(thisEnt.x)
}

func (thisEnt *gameEntity) AoiCLZ() CLPosValType {
	return CLPosValType(thisEnt.z)
}

func (thisEnt *gameEntity) AoiCLEntityID() EntityIDValType {
	return EntityIDValType(thisEnt.id)
}

func TestALL(t *testing.T) {
	entities := []*gameEntity{}
	for i := -100; i < 100; i++ {
		for j := -100; j < 100; j++ {
			newEnt := new(gameEntity)
			newEnt.x = float32(i)
			newEnt.z = float32(j)
			newEnt.id = uint32(1000000 + i*1000 + j)
			entities = append(entities, newEnt)
		}
	}
	fmt.Println("step: entities created", len(entities))

	aoiSpace := NewAOISpaceCL(3, 6)

	for _, ent := range entities {
		aoiSpace.AddEntity(ent)
	}

	fmt.Println("step: entities added", len(entities))

	for _, ent := range entities {
		aoiSpace.AddRangeOfEntity(ent, 5, 5, EVENT_ALL)
	}

	fmt.Println("step: range added", len(entities))

	testEnt := entities[20000]
	ids, err := aoiSpace.EntitiesInRange(testEnt, 3, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("EntitiesInRange", testEnt.id, ids)

	aoiSpace.MoveEntity(testEnt, testEnt.AoiCLX()+5, testEnt.AoiCLZ()+5)
	ids, err = aoiSpace.EntitiesInRange(testEnt, 3, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("EntitiesInRange after move", testEnt.id, ids)
}
