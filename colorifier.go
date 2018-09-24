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

// ColorDepth ...
type ColorDepth int

const (
	// Color4Bits ...
	Color4Bits ColorDepth = 4

	// Color8Bits ...
	Color8Bits ColorDepth = 8

	// Color24Bits ...
	Color24Bits ColorDepth = 24
)

// Colorifier ...
type colorifier struct {
	// Escape character
	Escape string

	// Reset all attributes
	ResetAll int

	// Reset color attributes
	ResetColors string

	// Color Depth and Bits to use
	ColorDepth ColorDepth

	// If color is disabled or not
	noColor bool
}

// EnableColor ...
func (c *colorifier) EnableColor() {
	c.noColor = false
}

// IsColorEnabled ...
func (c *colorifier) IsColorEnabled() bool {
	return !c.noColor
}

// DisableColor ...
func (c *colorifier) DisableColor() {
	c.noColor = true
}

// IsColorDisabled ...
func (c *colorifier) IsColorDisabled() bool {
	return c.noColor
}

func (c *colorifier) Color(depth ColorDepth, str string, codes ...int) string {
	if c.IsColorDisabled() {
		return str
	}

	var result string

	switch depth {
	case Color4Bits:
		result = Color4bits(str, codes...)
	case Color8Bits:
		result = str
	case Color24Bits:
		result = str
	default:
		result = Color4bits(str, codes...)
	}

	return result
}
