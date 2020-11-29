package abstract_factory

import "fmt"

//OrderMainDAO 为订单主记录
type OrderMainDao interface {
	SaveOrderMain()
}

//OrderDetailDAO 为订单详情纪录
type OrderDetailDao interface {
	SaveOrderDetail()
}

//DAOFactory DAO 抽象模式工厂接口
type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

//RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDao struct {

}

func (*RDBMainDao) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

//RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDao struct {

}

func (*RDBDetailDao) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBFactory struct {

}

func (*RDBFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (*RDBFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

//XMLMainDAO XML存储
type XMLMainDao struct {

}

func (*XMLMainDao) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

type XMLDetailDao struct {

}

func (*XMLDetailDao) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

type XMLDaoFactory struct {

}

func (*XMLDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (*XMLDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}