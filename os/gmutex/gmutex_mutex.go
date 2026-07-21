// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gmutex

import "sync"

// Mutex is a high level Mutex, which implements more rich features for mutex.
type Mutex struct {
	lockChan chan struct{}
	once     sync.Once
}

func (m *Mutex) init() {
	m.once.Do(func() {
		m.lockChan = make(chan struct{}, 1)
		m.lockChan <- struct{}{}
	})
}

func (m *Mutex) Lock() {
	m.init()
	<-m.lockChan
}

func (m *Mutex) Unlock() {
	m.lockChan <- struct{}{}
}

func (m *Mutex) TryLock() bool {
	m.init()
	select {
	case <-m.lockChan:
		return true
	default:
		return false
	}
}

// LockFunc locks the mutex for writing with given callback function `f`.
// If there's a write/reading lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
func (m *Mutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// TryLockFunc tries locking the mutex for writing with given callback function `f`.
// it returns true immediately if success, or if there's a write/reading lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
func (m *Mutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		result = true
		defer m.Unlock()
		f()
	}
	return
}
