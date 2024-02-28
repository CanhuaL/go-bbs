package logic

import (
	"go_bbs/dao/mysql"
	"go_bbs/models"
)

func CreateComment(p *models.ParamComment) error {
	if err := mysql.CreateComment(p); err != nil {
		return err
	}
	return nil
}
