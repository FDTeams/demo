package main

import "fmt"

// 创建插件接口
type plugins interface {
	run(target string)
}

// 创建插件hash,用于存放所有的插件
type pluginsMap struct {
	pluginMap map[string]plugins // 键为plugins接口类型
}

// 初始化插件结构体
func (p *pluginsMap) init() {
	p.pluginMap = make(map[string]plugins)
}

// 注册插件
func (p *pluginsMap) registerPlugin(name string, plugin plugins) {
	p.pluginMap[name] = plugin
}

type demo struct{}

// 实现插件接口
func (d *demo) run(target string) {
	fmt.Println(target, " from demo")
}

type demo1 struct{}

func (d *demo1) run(target string) {
	fmt.Println(target, " from demo1")
}

func main() {
	// 初始化插件hash数据结构
	pluginMap := new(pluginsMap)
	// 实例化具体插件实体
	de := new(demo)
	de1 := new(demo1)

	// 调用插件初始化方法
	pluginMap.init()

	// 调用注册插件方法
	pluginMap.registerPlugin("demo", de)
	pluginMap.registerPlugin("demo1", de1)

	// 循环遍历插件并运行
	for _, plug := range pluginMap.pluginMap {
		plug.run("https://wwww.baidu.com")
	}
}
