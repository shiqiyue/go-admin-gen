package core

func (c *GenContext) editReqDtoFullName() string {
	return c.Cfg.GetDtoPackage() + "/" + c.editReqDtoName()
}

func (c *GenContext) editReqDtoName() string {
	return c.modelName() + "EditDto"
}

func (c *GenContext) genReqDto() {

}
