package controllers

import (
	"encoding/json"
	"pkg/models"
	"pkg/utils/convert"
	"pkg/utils/queue"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

type ProblemController struct{}

var AddrQueue = queue.NewQueue(6)

func (ProblemController) CreateProblem(c *gin.Context) {
	problem := &models.Problem{}
	_ = c.BindJSON(&problem)
	problem.ID = uuid.New().String()[:6]
	err := models.CreateProblem(problem)
	if err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", problem, 1)

}

func (ProblemController) GetProblemInfo(c *gin.Context) {
	id := c.Param("id")
	problem, err := models.QueryProblem(id)
	if err == nil {
		ReturnSuccess(c, 0, "suceess", problem, 1)
		return
	}
	ReturnError(c, 1, err)
}
func (ProblemController) GetList(c *gin.Context) {
	result, err := models.GetAllProblem()
	if err == nil {
		ReturnSuccess(c, 0, "success", result, int64(len(result)))
		return
	}
	ReturnError(c, 1, err)
}

func (ProblemController) SubmitProblem(c *gin.Context) {
	data := JudgeStruct{}
	if Err := c.BindJSON(&data); Err != nil {
		ReturnError(c, 1, Err)
		return
	}
	problem, err := models.QueryProblem(data.ID)
	var judgecase []map[string]string
	_ = json.Unmarshal([]byte(problem.JudgeCase), &judgecase)
	var results []models.Result
	channel := make(chan RetuenStruct)
	hash := make(map[string]string)
	for _, v := range judgecase {
		data.Input = v["input"]
		hash[v["input"]] = v["output"]
		ip := AddrQueue.GetOne()
		go JudgeController{}.EvaluateCode(data, ip, channel)
	}
	for range judgecase {
		respond := <-channel
		r, _ := convert.Encode2String(respond.Info)
		var result models.Result
		_ = convert.Decode2Object(r, &result)
		if result.Error != 0 {
			result.Display = "Runtime error"
			results = append(results, result)
			continue
		}
		if cmp.Equal(result.Output, hash[result.Input]) == true {
			result.Display = "Accepted"
			results = append(results, result)
			continue
		}
		result.Display = "Wrong answer"
		results = append(results, result)
	}
	if err == nil {
		ReturnSuccess(c, 0, "success", results, int64(len(results)))
		judgeResultJSON, _ := json.Marshal(results)
		record := &models.Record{
			ID:           uuid.New().String(),
			Code:         data.Code,
			Language:     data.Language,
			ProblemId:    data.ID,
			JudgeResult:  string(judgeResultJSON),
			UserId:       data.UserID,
			UserName:     data.Account,
			ProblemTitle: problem.Title,
		}
		_ = models.CreateRecord(record)
		return
	}
	ReturnError(c, 1, err)
}
