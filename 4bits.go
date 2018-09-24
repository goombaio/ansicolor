// Copyright 2018, gossiper project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package ansicolor

/*
import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// escape character
	escape = "\x1b"

	// reset all attributes sequence
	resetAllAttributes = 0

	// reset onnly colors sequeance
	resetColors = "39;49"
)

// SGR or Select Graphic Rendition) ...
func SGR(str string, codes ...int) string {
	strcodes := make([]string, len(codes))
	for i, code := range codes {
		strcodes[i] = strconv.Itoa(code)
	}
	colorSequence := strings.Join(strcodes, ";")

	result := fmt.Sprintf("%s[%sm", escape, colorSequence)
	result += str
	result += fmt.Sprintf("%s[%dm", escape, resetAllAttributes)

	return result
}
*/
