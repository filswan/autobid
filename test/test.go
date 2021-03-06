package test

import (
	"fmt"
	"go-swan/common/constants"
	"go-swan/common/utils"
	"go-swan/logs"
	"go-swan/models"
	"go-swan/service"
	"math/rand"
)

func Test() {
	models.GetAutoBidMiners()
}

func testRandomInt2() {
	var stat [160]int
	for i := 0; i < 100; i++ {
		val := float64(rand.Intn(100)) * 1.5
		valInt := int(val)
		stat[valInt]++
		fmt.Println(val, valInt)
	}

	for x := range stat {
		if stat[x] != 0 {
			fmt.Println(x, stat[x])
		}
	}
	//fmt.Println(cnt5, cnt6, cnt7, cnt8, cnt9, cnt10, cntOther)
}

func testRandomInt1() {
	var stat [11]int
	for i := 0; i < 10; i++ {
		val := (rand.Float64()-rand.Float64())*(1+10-5) + 5
		valInt := int(val)
		stat[valInt]++
		fmt.Println(val, valInt)
	}

	for x := range stat {
		fmt.Println(x, stat[x])
	}
	//fmt.Println(cnt5, cnt6, cnt7, cnt8, cnt9, cnt10, cntOther)
}

func testRandomInt() {
	cnt5 := 0
	cnt6 := 0
	cnt7 := 0
	cnt8 := 0
	cnt9 := 0
	cnt10 := 0
	cntOther := 0
	for i := 0; i < 100; i++ {
		val := utils.GetRandInRange(5, 10)
		switch val {
		case 5:
			cnt5++
		case 6:
			cnt6++
		case 7:
			cnt7++
		case 8:
			cnt8++
		case 9:
			cnt9++
		case 10:
			cnt10++
		default:
			cntOther++
		}
		//fmt.Println(val)
	}
	fmt.Println(cnt5, cnt6, cnt7, cnt8, cnt9, cnt10, cntOther)
}

func TestTask_GetTasks() {
	tasks, err := models.GetTasks(0, 10, constants.TASK_STATUS_CREATED)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, task := range tasks {
		logs.GetLogger().Info(utils.ToJson(task))
	}
}

func TestTask_GetAutoBidTasks() {
	tasks, err := models.GetAutoBidTasks(0, 10, constants.TASK_STATUS_CREATED)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, task := range tasks {
		logs.GetLogger().Info(utils.ToJson(task))
	}
}

func TestTask_GetTaskById() {
	task, err := models.GetTaskById(1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(utils.ToJson(task))
}

func TestMiner_GetMiners() {
	miners, err := models.GetMiners(0, 10, constants.MINER_STATUS_ACTIVE)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, miner := range miners {
		logs.GetLogger().Info(utils.ToJson(miner))
	}
}

func TestMiner_GetAllMiners() {
	miners := service.GetMiners()

	for j := 0; j < 100; j++ {
		ratio := utils.GetRandInRange(0, 100)
		for i, miner := range miners {
			//score := miner.ScorePercent * multiple
			//logs.GetLogger().Info("Score:", ratio, randNum)
			if ratio < miner.Score {
				fmt.Println(i, " ScorePercent=", miner.Score, " MinerFid=", miner.MinerFid, " ratio=", ratio, " is selected")
				break
			}
		}
	}
}
