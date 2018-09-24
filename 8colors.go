// Copyright 2018, Goomba project Authors. All rights reserved.
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

import (
	"fmt"
	"strconv"
)

// Color8Colors or Select Graphic Rendition) ...
//
// Also check: https://en.wikipedia.org/wiki/ANSI_escape_code#Colorss
func Color8Colors(str string, foreground int, background int) string {
	colorSequence := strconv.Itoa(foreground) + ";" + strconv.Itoa(background)

	result := fmt.Sprintf("%s[%sm", Colorifier.Escape, colorSequence)
	result += str
	result += fmt.Sprintf("%s[%dm", Colorifier.Escape, Colorifier.ResetAll)

	return result
}
