package abstract_factory

import "testing"

func getMainAndDetail(factory DaoFactory)  {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func TestExampleRdbFactory(t *testing.T) {
	var factory DaoFactory
	factory = &RDBFactory{}
	getMainAndDetail(factory)
}

func TestExampleXmlFactory(t *testing.T) {
	var factory DaoFactory
	factory = &XMLDaoFactory{}
	getMainAndDetail(factory)
}