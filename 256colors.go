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

// Color256Colors or Select Graphic Rendition) ...
//
// ESC[ … 38;5;<n> … m Select foreground color
// ESC[ … 48;5;<n> … m Select background color
//   0-  7:  standard colors (as in ESC [ 30–37 m)
//   8- 15:  high intensity colors (as in ESC [ 90–97 m)
//  16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)
// 232-255:  grayscale from black to white in 24 steps
//
// Also check: https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
func Color256Colors(str string, foreground int, background int) string {
	fgColorSequence := "38;5;" + strconv.Itoa(foreground)
	bgColorSequence := "48;5;" + strconv.Itoa(background)
	result := ""
	result += fmt.Sprintf("%s[%sm", Colorifier.Escape, fgColorSequence)
	result += fmt.Sprintf("%s[%sm", Colorifier.Escape, bgColorSequence)
	result += str
	result += fmt.Sprintf("%s[%dm", Colorifier.Escape, Colorifier.ResetAll)

	return result
}
