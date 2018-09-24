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

package ansicolor_test

import (
	"testing"

	"github.com/goombaio/ansicolor"
)

func TestColor_DisableColor(t *testing.T) {
	ansicolor.DisableColor()

	if ansicolor.Colorifier.IsColorEnabled() {
		t.Fatalf("Expected to be false but it does not")
	}

	if !ansicolor.Colorifier.IsColorDisabled() {
		t.Fatalf("Expected to be true but it does not")
	}
}

func TestColor_EnabledColor(t *testing.T) {
	ansicolor.EnableColor()

	if !ansicolor.Colorifier.IsColorEnabled() {
		t.Fatalf("Expected to be true but it does not")
	}

	if ansicolor.Colorifier.IsColorDisabled() {
		t.Fatalf("Expected to be false but it does not")
	}
}

func TestColor_Color(t *testing.T) {
	str := "This is an string"

	result := ansicolor.Color(str)

	if result != str {
		t.Fatalf("result expected %q but got %q", str, result)
	}
}
