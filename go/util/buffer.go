/*******************************************************************************
 * Copyright (c) 2018 - Present VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: BSD-2
 ******************************************************************************/

package util

import "sync"

var bsPool = sync.Pool{New: func() interface{} { return make([]byte, 0, 512*1024) }}

func GetByteSlice() []byte {
	return bsPool.Get().([]byte)
}

func PutByteSlice(bs []byte) {
	bsPool.Put(bs)
}
