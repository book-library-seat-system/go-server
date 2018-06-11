/***************************************************************************
Copyright(C) 2018
Author: huziang
Description: 用于控制读写锁
Date: 2018年6月11日 星期一 上午11:00
****************************************************************************/

package mutexmanager

import "sync"

// RWMutexManager 读写锁管理员
type RWMutexManager struct {
	// 使用map管理锁
	Locks map[string]*sync.RWMutex
}

// New 生成一个新的管理者
func New() *RWMutexManager {
	return &RWMutexManager{Locks: make(map[string]*sync.RWMutex)}
}

// AddLock 添加锁，如果重复，则不做操作
func (manager *RWMutexManager) AddLock(name string) *RWMutexManager {
	if _, ok := manager.Locks[name]; !ok {
		manager.Locks[name] = new(sync.RWMutex)
	}
	return manager
}

// GetLock 得到锁
func (manager *RWMutexManager) GetLock(name string) *sync.RWMutex {
	return manager.Locks[name]
}

// RLock reader上锁
func (manager *RWMutexManager) RLock(name string) *RWMutexManager {
	manager.Locks[name].RLock()
	return manager
}

// RUnlock reader解锁
func (manager *RWMutexManager) RUnlock(name string) *RWMutexManager {
	manager.Locks[name].RUnlock()
	return manager
}

// WLock writer上锁
func (manager *RWMutexManager) WLock(name string) *RWMutexManager {
	manager.Locks[name].Lock()
	return manager
}

// WUnlock writer解锁
func (manager *RWMutexManager) WUnlock(name string) *RWMutexManager {
	manager.Locks[name].Unlock()
	return manager
}
