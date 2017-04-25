package main

import (
	"fmt"
	"loyocloud-infrastructure/services/mongo"
	"loyocloud-infrastructure/tmodels"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type MyFLowEvent struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
}

type QueryParam struct {
	QType      int             //查询类型
	StartAt    int64           //开始时间
	EndAt      int64           //结束时间
	FlowId     bson.ObjectId   //流程id
	KeyWord    string          //关键字
	Status     int             //状态
	IsOverTime int             //是否逾期 1:已逾期 2:未逾期
	UserIds    []bson.ObjectId //人员
	Field      string          //排序字段
	OrderBy    int             //排序方式
	Source     int             //0:全部 1:规则流程 2:售后流程
	IsAll      bool            //是否全部(执行及分派)
	Role       int             //事件角色 1: 组织者 2: 执行者
	FlowType   int             //流程类型
	HideEmptye bool            //忽略没有事件的员工
	Page       *tmodels.Pagination
}

type FlowEventStatResult struct {
	EventId      bson.ObjectId `bson:"event_id" json:"eventId"`           //事件ID
	ExecutorName string        `bson:"executor_name" json:"executorName"` //执行者名字
	Total        int           `bson:"total" json:"total"`                //事件总数
	Processing   int           `bson:"processing" json:"processing"`      //处理中
	Finished     int           `bson:"finished" json:"finished"`          //已完成
	Pending      int           `bson:"pending" json:"pending"`            //待分派
	Terminated   int           `bson:"terminated" json:"terminated"`      //意外终止
}

type EventStatus int

const (
	_                 EventStatus = iota
	Event_Processing              //进行中
	Event_Finished                //已完成
	Event_NotStart                //未开始
	Event_Termination             //意外终止
	Event_Pending                 //待分派(没有执行者)
)

func (self *MyFLowEvent) GetMgoInfo() (string, string, string) {
	s := "Peter"
	d := "Ocean"
	c := "flow_event"
	return d, c, s
}

func (self *MyFLowEvent) GetId() bson.ObjectId {
	return self.Id
}

func getFlowDBKey(queryType int) string {
	switch queryType {
	case 1: // 开始时间
		return "start_at"
	case 2: // 完成时间
		return "finished_at"
	default:
		return ""
	}
}

func getEventDBKey(queryType int) string {
	return fmt.Sprintf("event_nodes.%s", getFlowDBKey(queryType))
}

func statusCond(eventStatus EventStatus) mongo.Map {
	cond := mongo.M("$cond",
		[]interface{}{
			mongo.M("$event_nodes.status", eventStatus, mongo.Equal),
			1,
			0,
		})
	return cond
}

func main() {
	//Initialize MongoDB pipeline
	o := mongo.NewOrm()
	pipe := o.Pipe(new(MyFLowEvent))

	//Initialize query conditions
	param := QueryParam{
		QType:  0,
		Source: 0,
	}
	companyId := bson.ObjectIdHex("57c7e26a608e4f0391d89eb9")
	dbKey := getEventDBKey(param.QType)

	//construct flow filter
	flowCond := mongo.M("company_id", companyId).
		MCond("flow_type", param.Source, param.Source > 0).
		MCond(dbKey, param.StartAt, param.QType != 0 && param.StartAt > 0, mongo.GTE).
		MCond(dbKey, param.EndAt, param.QType != 0 && param.EndAt > 0, mongo.LTE)

	//construt event filter
	//模糊查询
	eventCond := mongo.MCond("event_nodes.executor_name", param.KeyWord, len(strings.TrimSpace(param.KeyWord)) > 0, mongo.Like)
	if param.Role == 1 {
		eventCond.M("event_nodes.executorId", param.UserIds, mongo.IN)
	} else if param.Role == 2 {
		eventCond.M("organizer_id", param.UserIds, mongo.IN)
	}

	// Merge common cond to match cond
	for k, v := range eventCond {
		flowCond.M(k, v)
	}

	// Match by above condition and unwind event_nodes array
	pl := pipe.Do(mongo.Match, flowCond).Do(mongo.Unwind, "$event_nodes")

	//match flow event
	pl.Do(mongo.Match, flowCond)

	//construct group conditon
	groupCond := mongo.M("_id", "$event_nodes.executorId").
		M("executor_name", "$event_nodes.ExecutorName", mongo.Last).
		M("total", 1, mongo.Sum).
		M("processing", statusCond(Event_Processing), mongo.Sum).
		M("finished", statusCond(Event_Finished), mongo.Sum).
		M("pending", statusCond(Event_Pending), mongo.Sum).
		M("terminated", statusCond(Event_Termination), mongo.Sum)

	//take group here
	pl.Do(mongo.Group, groupCond)
	fmt.Println(pl.Pipelines())

	/*var err error

	//sort result in case of page parameter is not nil
	if param.Page != nil {
		pl.Do(mongo.Sort, mongo.M(param.Field, param.OrderBy), len(param.Field) > 0)
		//output data
		param.Page.Records = &[]*FlowEventStatResult{}
		err = pl.Pagination(param.Page)
	} else {
		//only need to calculate total data
		pl.Do(mongo.Group,
			mongo.M("_id", nil).
				M("total", "$total", mongo.Sum).
				M("processing", "processing", mongo.Sum).
				M("finished", "$finished", mongo.Sum).
				M("pending", "$pending", mongo.Sum).
				M("terminated", "$terminated", mongo.Sum))
		param.Page = &tmodels.Pagination{
			Records: &FlowEventStatResult{
				ExecutorName: "总计",
			},
		}
		err = pl.One(param.Page)
	}
	if err != nil {
		fmt.Println()
	}*/
}
