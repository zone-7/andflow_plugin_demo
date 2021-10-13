package main

import (
	"github.com/zone-7/andflow_plugin"
	"context"
	"fmt"
)

type Demo struct {

}
func (d *Demo) GetName()string{
	return "demo"
}
func (d *Demo) Init(callback interface{}){
	//回调对象
	initCallbacker:=andflow_plugin.ParseInitCallbacker(callback)
	pluginpath := initCallbacker.GetFlowPluginPath()

	fmt.Println(pluginpath)
}
func (d *Demo) PrepareMetadata(flowCode string, metadata string) string{
	return metadata
}
func (d *Demo) Filter(ctx context.Context,runtimeId string,preActionId string, actionId string,callback interface{})(bool,error){
	//回调对象
	actionCallbacker:=andflow_plugin.ParseActionCallbacker(callback)

	//根据上个节点的state值判断是否执行
	state := actionCallbacker.GetRuntimeActionData(preActionId,"state")
	if state=="no"{
		return false,nil
	}

	return true,nil
}
func (d *Demo) Exec(ctx context.Context,runtimeId string,preActionId string, actionId string,callback interface{})(interface{},error){
	//回调对象
	actionCallbacker:=andflow_plugin.ParseActionCallbacker(callback)

	//获取流程运行参数
	value:=actionCallbacker.GetRuntimeParam("demo")
	fmt.Println(value)


	//获取上个节点设置的所有数据
	valuesMapPreAction := actionCallbacker.GetRuntimeActionDatas(preActionId)
	fmt.Println(valuesMapPreAction)

	//获取当前节点所有数据
	valuesMap := actionCallbacker.GetRuntimeActionDatas(actionId)
	fmt.Println(valuesMap)

	//获取当前节点数据
	demoValue := actionCallbacker.GetRuntimeActionData(actionId,"demo")
	fmt.Println(demoValue)


	//设置当前节点数据
	actionCallbacker.SetRuntimeActionData(actionId,"demo","hello this is action data")

	//设置参数数据
	actionCallbacker.SetRuntimeParam( "demo","hello this is param")

	//设置运行结果数据
	actionCallbacker.SetRuntimeData("demo","hello this is result  你好")


	return valuesMap,nil
}


func main(){
	andflow_plugin.InitPlugin(&Demo{})
}
