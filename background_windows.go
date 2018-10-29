package main

import (
	"syscall"
)

// https://docs.microsoft.com/en-us/windows/desktop/api/processthreadsapi/nf-processthreadsapi-setthreadpriority

var (
	dllKernel32 = syscall.MustLoadDLL("kernel32")

	thrdGet      = dllKernel32.MustFindProc("GetCurrentThread")
	thrdPriority = dllKernel32.MustFindProc("SetThreadPriority")
	thrdBoost    = dllKernel32.MustFindProc("SetThreadPriorityBoost")

	procGet      = dllKernel32.MustFindProc("GetCurrentProcess")
	procPriority = dllKernel32.MustFindProc("SetPriorityClass")
	procBoost    = dllKernel32.MustFindProc("SetProcessPriorityBoost")
)

const (
	thrdPriorityIdle        uintptr = 0xFFFFFFF1 // (-15) THREAD_PRIORITY_IDLE
	thrdModeBackgroundBegin uintptr = 0x00010000 // THREAD_MODE_BACKGROUND_BEGIN
	thrdModeBackgroundEnd   uintptr = 0x00020000 // THREAD_MODE_BACKGROUND_END

	procPriorityIdle        uintptr = 0x00000040 // IDLE_PRIORITY_CLASS
	procModeBackgroundBegin uintptr = 0x00100000 // PROCESS_MODE_BACKGROUND_BEGIN
	procModeBackgroundEnd   uintptr = 0x00200000 // PROCESS_MODE_BACKGROUND_END
)

func goroutineBackgroundStart() {
	mythread, _, _ := thrdGet.Call()
	thrdPriority.Call(mythread, thrdPriorityIdle)
	thrdPriority.Call(mythread, thrdModeBackgroundBegin)
	thrdBoost.Call(mythread, 1)
}

func goroutineBackgroundStop() {
	mythread, _, _ := thrdGet.Call()
	thrdPriority.Call(mythread, thrdModeBackgroundEnd)
	thrdBoost.Call(mythread, 0)
}

func processBackgroundStart() {
	myproc, _, _ := procGet.Call()
	procPriority.Call(myproc, procPriorityIdle)
}

func processBackgroundStop() {
}
