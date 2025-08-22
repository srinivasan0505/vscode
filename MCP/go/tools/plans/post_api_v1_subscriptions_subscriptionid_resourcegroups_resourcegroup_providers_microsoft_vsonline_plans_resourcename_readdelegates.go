package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vsonline/mcp-server/config"
	"github.com/vsonline/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_microsoft_vsonline_plans_resourcename_readdelegatesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		resourceGroupVal, ok := args["resourceGroup"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroup"), nil
		}
		resourceGroup, ok := resourceGroupVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroup"), nil
		}
		resourceNameVal, ok := args["resourceName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceName"), nil
		}
		resourceName, ok := resourceNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceName"), nil
		}
		url := fmt.Sprintf("%s/api/v1/subscriptions/%s/resourceGroups/%s/providers/Microsoft.VSOnline/plans/%s/readDelegates", cfg.BaseURL, subscriptionId, resourceGroup, resourceName)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreatePost_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_microsoft_vsonline_plans_resourcename_readdelegatesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_v1_subscriptions_subscriptionId_resourceGroups_resourceGroup_providers_Microsoft_VSOnline_plans_resourceName_readDelegates",
		mcp.WithDescription(""),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("")),
		mcp.WithString("resourceGroup", mcp.Required(), mcp.Description("")),
		mcp.WithString("resourceName", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_microsoft_vsonline_plans_resourcename_readdelegatesHandler(cfg),
	}
}
