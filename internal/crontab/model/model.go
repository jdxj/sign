package model

import (
	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/proto/crontab"
)

func GetTasks(req *crontab.GetTasksReq) (*crontab.GetTasksRsp, error) {
	query := db.Gorm.Raw(`
SELECT
	task.task_id,
	task.user_id,
	task.describe,
	task.kind,
	sp.spec,
	task.secret_id
FROM
	task LEFT JOIN specification AS sp ON task.spec_id = sp.spec_id
WHERE
	task.user_id = ?;
`, req.UserID)

	rsp := &crontab.GetTasksRsp{}
	return rsp, query.Find(&rsp.List).Error
}
