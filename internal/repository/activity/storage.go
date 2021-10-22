package activity

import (
	"github.com/ozonmp/omp-bot/internal/model/activity"
	"time"
)

var Corrections = []activity.Correction {
	{ID: 1, Timestamp: time.Now(), UserID: 1, Object: "order1 description", Action: "update",
		Data: &activity.Data { OriginalData: "test11", RevisedData: "test12" }},
	{ID: 2, Timestamp: time.Now(), UserID: 2, Object: "order2 description", Action: "update",
		Data: &activity.Data { OriginalData: "test21", RevisedData: "test22" }},
	{ID: 3, Timestamp: time.Now(), UserID: 3, Object: "order3 description", Action: "update",
		Data: &activity.Data { OriginalData: "test31", RevisedData: "test32" }},
	{ID: 4, Timestamp: time.Now(), UserID: 4, Object: "order4 description", Action: "update",
		Data: &activity.Data { OriginalData: "test41", RevisedData: "test42" }},
	{ID: 5, Timestamp: time.Now(), UserID: 5, Object: "order5 description", Action: "update",
		Data: &activity.Data { OriginalData: "test51", RevisedData: "test52" }},
	{ID: 6, Timestamp: time.Now(), UserID: 6, Object: "order6 description", Action: "update",
		Data: &activity.Data { OriginalData: "test61", RevisedData: "test62" }},
	{ID: 7, Timestamp: time.Now(), UserID: 7, Object: "order7 description", Action: "update",
		Data: &activity.Data { OriginalData: "test71", RevisedData: "test72" }},
}