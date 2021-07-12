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

func ExampleNewCampaignExperimentClient() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCampaignExperimentClient_GetCampaignExperiment() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.GetCampaignExperimentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCampaignExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCampaignExperimentClient_CreateCampaignExperiment() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.CreateCampaignExperimentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateCampaignExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCampaignExperimentClient_MutateCampaignExperiments() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.MutateCampaignExperimentsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MutateCampaignExperiments(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCampaignExperimentClient_GraduateCampaignExperiment() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.GraduateCampaignExperimentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GraduateCampaignExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCampaignExperimentClient_PromoteCampaignExperiment() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.PromoteCampaignExperimentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.PromoteCampaignExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCampaignExperimentClient_EndCampaignExperiment() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.EndCampaignExperimentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.EndCampaignExperiment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCampaignExperimentClient_ListCampaignExperimentAsyncErrors() {
	ctx := context.Background()
	c, err := gads.NewCampaignExperimentClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.ListCampaignExperimentAsyncErrorsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCampaignExperimentAsyncErrors(ctx, req)
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
