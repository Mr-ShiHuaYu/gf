// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gmutex

import "sync"

// RWMutex is a high level RWMutex, which implements more rich features for mutex.
type RWMutex struct {
	mu          sync.Mutex
	readers     int
	writerActive bool
	writable    chan struct{}
	readable    chan struct{}
	once        sync.Once
}

func (m *RWMutex) init() {
	m.once.Do(func() {
		m.writable = make(chan struct{}, 1)
		m.readable = make(chan struct{}, 1)
		m.writable <- struct{}{}
		m.readable <- struct{}{}
	})
}

func (m *RWMutex) Lock() {
	m.init()
	<-m.writable
	m.mu.Lock()
	m.writerActive = true
	m.mu.Unlock()
	<-m.readable
}

func (m *RWMutex) Unlock() {
	m.mu.Lock()
	m.writerActive = false
	m.mu.Unlock()
	m.readable <- struct{}{}
	m.writable <- struct{}{}
}

func (m *RWMutex) RLock() {
	m.init()
	m.mu.Lock()
	if m.readers == 0 && !m.writerActive {
		<-m.readable
	}
	m.readers++
	m.mu.Unlock()
}

func (m *RWMutex) RUnlock() {
	m.mu.Lock()
	m.readers--
	if m.readers == 0 && !m.writerActive {
		m.readable <- struct{}{}
	}
	m.mu.Unlock()
}

func (m *RWMutex) TryLock() bool {
	m.init()
	select {
	case <-m.writable:
		m.mu.Lock()
		if m.readers > 0 {
			m.mu.Unlock()
			m.writable <- struct{}{}
			return false
		}
		m.writerActive = true
		m.mu.Unlock()
		select {
		case <-m.readable:
			return true
		default:
			m.mu.Lock()
			m.writerActive = false
			m.mu.Unlock()
			m.writable <- struct{}{}
			return false
		}
	default:
		return false
	}
}

func (m *RWMutex) TryRLock() bool {
	m.init()
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.writerActive {
		return false
	}
	if m.readers == 0 {
		select {
		case <-m.readable:
		default:
			return false
		}
	}
	m.readers++
	return true
}

// LockFunc locks the mutex for writing with given callback function `f`.
// If there's a write/reading lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
func (m *RWMutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()
	f()
}

// RLockFunc locks the mutex for reading with given callback function `f`.
// If there's a writing lock the mutex, it will block until the lock is released.
//
// It releases the lock after `f` is executed.
func (m *RWMutex) RLockFunc(f func()) {
	m.RLock()
	defer m.RUnlock()
	f()
}

// TryLockFunc tries locking the mutex for writing with given callback function `f`.
// it returns true immediately if success, or if there's a write/reading lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
func (m *RWMutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		result = true
		defer m.Unlock()
		f()
	}
	return
}

// TryRLockFunc tries locking the mutex for reading with given callback function `f`.
// It returns true immediately if success, or if there's a writing lock on the mutex,
// it returns false immediately.
//
// It releases the lock after `f` is executed.
func (m *RWMutex) TryRLockFunc(f func()) (result bool) {
	if m.TryRLock() {
		result = true
		defer m.RUnlock()
		f()
	}
	return
}
