package facade

import "fmt"

type Api interface {
	Test() string
}

type ApiImpl struct {
	a AModuleApi
	b BModuleApi
}

func (a *ApiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewApi() Api {
	return &ApiImpl{
		a: NewAModuleApi(),
		b: NewBModuleAPi(),
	}
}

type AModuleApi interface {
	TestA() string
}

type aModuleImpl struct {
}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

func NewAModuleApi() AModuleApi {
	return &aModuleImpl{}
}

type BModuleApi interface {
	TestB() string
}

type bModuleImpl struct {
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

func NewBModuleAPi() BModuleApi {
	return &bModuleImpl{}
}
