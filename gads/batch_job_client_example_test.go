// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gads_test

import (
	"context"

	gads "github.com/opteo/google-ads-go"
	"google.golang.org/api/iterator"
	servicespb "github.com/opteo/google-ads-go/services"
)

func ExampleNewBatchJobClient() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleBatchJobClient_MutateBatchJob() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.MutateBatchJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MutateBatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBatchJobClient_GetBatchJob() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.GetBatchJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetBatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleBatchJobClient_ListBatchJobResults() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.ListBatchJobResultsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListBatchJobResults(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleBatchJobClient_RunBatchJob() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.RunBatchJobRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RunBatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleBatchJobClient_AddBatchJobOperations() {
	ctx := context.Background()
	c, err := gads.NewBatchJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.AddBatchJobOperationsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AddBatchJobOperations(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
