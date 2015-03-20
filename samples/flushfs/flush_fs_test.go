// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flushfs_test

import (
	"sync"
	"testing"

	"github.com/jacobsa/fuse/samples"
	"github.com/jacobsa/fuse/samples/flushfs"
	. "github.com/jacobsa/ogletest"
)

func TestFlushFS(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type FlushFSTest struct {
	samples.SampleTest

	mu sync.Mutex

	// GUARDED_BY(mu)
	flushes []string

	// GUARDED_BY(mu)
	fsyncs []string
}

func init() { RegisterTestSuite(&FlushFSTest{}) }

func (t *FlushFSTest) SetUp(ti *TestInfo) {
	var err error

	// Set up a file system.
	reportTo := func(slice *[]string) func(string) {
		return func(s string) {
			t.mu.Lock()
			defer t.mu.Unlock()
			*slice = append(*slice, s)
		}
	}

	t.FileSystem, err = flushfs.NewFileSystem(
		reportTo(&t.flushes),
		reportTo(&t.fsyncs))

	if err != nil {
		panic(err)
	}

	// Mount it.
	t.SampleTest.SetUp(ti)
}

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

// Return a copy of the current contents of t.flushes.
//
// LOCKS_EXCLUDED(t.mu)
func (t *FlushFSTest) getFlushes() (p []string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	p = make([]string, len(t.flushes))
	copy(p, t.flushes)
	return
}

// Return a copy of the current contents of t.fsyncs.
//
// LOCKS_EXCLUDED(t.mu)
func (t *FlushFSTest) getFsyncs() (p []string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	p = make([]string, len(t.fsyncs))
	copy(p, t.fsyncs)
	return
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *FlushFSTest) Close_NonOverlappingFileHandles() {
	AssertTrue(false, "TODO")
}

func (t *FlushFSTest) Close_OverlappingFileHandles() {
	AssertTrue(false, "TODO")
}

func (t *FlushFSTest) Close_ReadOnly() {
	AssertTrue(false, "TODO")
}

func (t *FlushFSTest) Close_WriteOnly() {
	AssertTrue(false, "TODO")
}

func (t *FlushFSTest) Fsync() {
	AssertTrue(false, "TODO")
}
