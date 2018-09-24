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

package ansicolor_test

import (
	"fmt"
	"testing"

	"github.com/goombaio/ansicolor"
)

func TestColor_Color_truecolors(t *testing.T) {
	str := "A green foreground string"
	expected := "\x1b[38;2;76;168;67m\x1b[48;2;55;55;55mA green foreground string\x1b[0m"

	result := ansicolor.ColorTrueColors(str, 76, 168, 67, 55, 55, 55)

	if result != expected {
		t.Fatalf("result expected %q but got %q", expected, result)
	}

	fmt.Println(result)
}
