package main

/*
Author: isar
date: 2019-01-10 16:10:49
Function: golang implement singleton pattern
Changed: No
*/

import (
	"fmt"
)

var m *Manager

func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("this is singleton pattern")
}

func main() {
	m := GetInstance()
	m.Manage()
}

/*
单例模式的优点：

	在内存中只有一个对象，节省内存空间。
	避免频繁的创建销毁对象，可以提高性能。
	避免对共享资源的多重占用。
	可以全局访问。
适用场景：由于单例模式的以上优点，所以是编程中用的比较多的一种设计模式。我总结了一下我所知道的适合使用单例模式的场景：

	需要频繁实例化然后销毁的对象。
	创建对象时耗时过多或者耗资源过多，但又经常用到的对象。
	有状态的工具类对象。
	频繁访问数据库或文件的对象。
	以及其他我没用过的所有要求只有一个对象的场景。
单例模式注意事项：

	只能使用单例类提供的方法得到单例对象，不要使用反射，否则将会实例化一个新对象。
	不要做断开单例类对象与类中静态引用的危险操作。
	多线程使用单例使用共享资源时，注意线程安全问题。
*/
