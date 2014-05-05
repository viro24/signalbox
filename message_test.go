/*
 * Copyright (c) Clinton Freeman 2014
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"testing"
)

func TestUtf8Encoding(t *testing.T) {
	message := string([]byte{0xff, 0xfe, 0xfd})
	_, _, err := ParseMessage(message)
	if err == nil {
		t.Errorf("Expected utf8 error")
	}

	_, _, err = ParseMessage("/announce|{\"id\":\"dc6ac0ae-6e15-409b-b211-228a8f4a43b9\"}|{\"browser\":\"node\",\"browserVersion\":\"?\",\"id\":\"dc6ac0ae-6e15-409b-b211-228a8f4a43b9\",\"agent\":\"signaller@0.18.3\",\"room\":\"test-room\"}")
	if err != nil {
		t.Errorf("Unexpected utf8 error")
	}
}

func TestPartSeperation(t *testing.T) {
	message := "/announce"
	_, _, err := ParseMessage(message)
	if err != nil {
		t.Errorf("")
	}
}
