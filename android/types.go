// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 25 May 2016 17:32:37 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package android

/*
#cgo LDFLAGS: -landroid -llog
#include <android/api-level.h>
#include <android/asset_manager.h>
#include <android/asset_manager_jni.h>
#include <android/configuration.h>
#include <android/input.h>
#include <android/keycodes.h>
#include <android/log.h>
#include <android/looper.h>
#include <android/native_activity.h>
#include <android/native_window.h>
#include <android/native_window_jni.h>
#include <android/obb.h>
#include <android/rect.h>
#include <android/storage_manager.h>
#include <android/tts.h>
#include <android/window.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// AssetManager as declared in android/asset_manager.h:28
type AssetManager C.AAssetManager

// AssetDir as declared in android/asset_manager.h:31
type AssetDir C.AAssetDir

// Asset as declared in android/asset_manager.h:34
type Asset C.AAsset

// Jobject type as declared in include/jni.h:98
type Jobject unsafe.Pointer

// JNINativeMethod as declared in include/jni.h:147
type JNINativeMethod C.JNINativeMethod

// JNIEnv as declared in include/jni.h:157
type JNIEnv C.JNIEnv

// JavaVM as declared in include/jni.h:158
type JavaVM C.JavaVM

// JavaVMAttachArgs as declared in include/jni.h:1106
type JavaVMAttachArgs C.JavaVMAttachArgs

// JavaVMOption as declared in include/jni.h:1115
type JavaVMOption C.JavaVMOption

// JavaVMInitArgs as declared in include/jni.h:1123
type JavaVMInitArgs C.JavaVMInitArgs

// Configuration as declared in android/configuration.h:27
type Configuration C.AConfiguration

// Looper as declared in android/looper.h:39
type Looper C.ALooper

// LooperCallbackFunc type as declared in android/looper.h:159
type LooperCallbackFunc func(fd int32, events int32, data unsafe.Pointer) int32

// NativeActivity as declared in android/native_activity.h:108
type NativeActivity struct {
	Callbacks        *NativeActivityCallbacks
	Vm               *JavaVM
	Env              *JNIEnv
	Clazz            Jobject
	InternalDataPath string
	ExternalDataPath string
	SdkVersion       int32
	Instance         unsafe.Pointer
	AssetManager     *AssetManager
	ObbPath          string
	ref2cc295bf      *C.ANativeActivity
	allocs2cc295bf   interface{}
}

// NativeActivityCallbacks as declared in android/native_activity.h:225
type NativeActivityCallbacks C.ANativeActivityCallbacks

// NativeActivityCreateFunc type as declared in android/native_activity.h:235
type NativeActivityCreateFunc func(activity *NativeActivity, savedState unsafe.Pointer, savedStateSize uint32)

// InputEvent as declared in android/input.h:136
type InputEvent C.AInputEvent

// InputQueue as declared in android/input.h:803
type InputQueue C.AInputQueue

// NativeWindow as declared in android/native_window.h:36
type NativeWindow C.ANativeWindow

// NativeWindowBuffer as declared in android/native_window.h:57
type NativeWindowBuffer struct {
	Width          int32
	Height         int32
	Stride         int32
	Format         int32
	Bits           unsafe.Pointer
	Reserved       [6]uint32
	ref3db2646c    *C.ANativeWindow_Buffer
	allocs3db2646c interface{}
}

// Rect as declared in android/rect.h:35
type Rect struct {
	Left           int32
	Top            int32
	Right          int32
	Bottom         int32
	ref9511c547    *C.ARect
	allocs9511c547 interface{}
}

// ObbInfo as declared in android/obb.h:28
type ObbInfo C.AObbInfo

// StorageManager as declared in android/storage_manager.h:28
type StorageManager C.AStorageManager

// StorageManagerObbCallbackFunc type as declared in android/storage_manager.h:98
type StorageManagerObbCallbackFunc func(filename string, state int32, data unsafe.Pointer)
