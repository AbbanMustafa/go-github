// Copyright 2014 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRepositoriesService_Merge(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &RepositoryMergeRequest{
		Base:          String("b"),
		Head:          String("h"),
		CommitMessage: String("c"),
	}

	mux.HandleFunc("/repos/o/r/merges", func(w http.ResponseWriter, r *http.Request) {
		v := new(RepositoryMergeRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"sha":"s"}`)
	})

	ctx := context.Background()
	commit, _, err := client.Repositories.Merge(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Repositories.Merge returned error: %v", err)
	}

	want := &RepositoryCommit{SHA: String("s")}
	if !cmp.Equal(commit, want) {
		t.Errorf("Repositories.Merge returned %+v, want %+v", commit, want)
	}

	const methodName = "Merge"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Repositories.Merge(ctx, "\n", "\n", input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Repositories.Merge(ctx, "o", "r", input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestRepositoryMergeRequest_Marshal(t *testing.T) {
	testJSONMarshal(t, &RepositoryMergeRequest{}, "{}")

	u := &RepositoryMergeRequest{
		Base:          String("base"),
		Head:          String("head"),
		CommitMessage: String("cm"),
	}

	want := `{
		"base": "base",
		"head": "head",
		"commit_message": "cm"
	}`

	testJSONMarshal(t, u, want)
}

func TestRepositoriesService_MergeUpstream(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &RepoMergeUpstreamRequest{
		Branch: String("b"),
	}

	mux.HandleFunc("/repos/o/r/merge-upstream", func(w http.ResponseWriter, r *http.Request) {
		v := new(RepoMergeUpstreamRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !cmp.Equal(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"merge_type":"m"}`)
	})

	ctx := context.Background()
	result, _, err := client.Repositories.MergeUpstream(ctx, "o", "r", input)
	if err != nil {
		t.Errorf("Repositories.MergeUpstream returned error: %v", err)
	}

	want := &RepoMergeUpstreamResult{MergeType: String("m")}
	if !cmp.Equal(result, want) {
		t.Errorf("Repositories.MergeUpstream returned %+v, want %+v", result, want)
	}

	const methodName = "MergeUpstream"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Repositories.MergeUpstream(ctx, "\n", "\n", input)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Repositories.MergeUpstream(ctx, "o", "r", input)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}
