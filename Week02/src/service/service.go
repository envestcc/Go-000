package service

import (
	"GO-000/Week02/dao"

	"github.com/pkg/errors"
)

func xxx(id uint64) error {
	d := dao.New()

	if err := d.Get(id); err != nil {
		if dao.IsErrNoRows(err) {
			// 执行数据库中没有记录时的处理逻辑
			// fmt.Printf("NoRows: %+v", err)
		} else {
			// 执行其他错误的处理逻辑，通常是返回
			return errors.WithMessage(err, "xxx")
		}
	}

	return nil
}
