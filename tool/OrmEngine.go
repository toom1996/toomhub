package tool

import "github.com/go-xorm/xorm"

type Orm struct {
	*xorm.Engine
}

func OrmEngine() (*Orm, error) {
	engine, err := xorm.NewEngine("mysql", "")

	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	return orm, nil
}
