package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/vsonline/mcp-server/config"
	"github.com/vsonline/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_github_network_resourcetype_resourcename_resourcereadbeginHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		resourceTypeVal, ok := args["resourceType"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceType"), nil
		}
		resourceType, ok := resourceTypeVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceType"), nil
		}
		resourceNameVal, ok := args["resourceName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceName"), nil
		}
		resourceName, ok := resourceNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceName"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.NetworkSettingsResource
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/api/v1/subscriptions/%s/resourceGroups/%s/providers/GitHub.Network/%s/%s/resourceReadBegin", cfg.BaseURL, subscriptionId, resourceGroup, resourceType, resourceName)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreatePost_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_github_network_resourcetype_resourcename_resourcereadbeginTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_v1_subscriptions_subscriptionId_resourceGroups_resourceGroup_providers_GitHub_Network_resourceType_resourceName_resourceReadBegin",
		mcp.WithDescription(""),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("")),
		mcp.WithString("resourceGroup", mcp.Required(), mcp.Description("")),
		mcp.WithString("resourceType", mcp.Required(), mcp.Description("")),
		mcp.WithString("resourceName", mcp.Required(), mcp.Description("")),
		mcp.WithString("provisioningState", mcp.Description("")),
		mcp.WithObject("tags", mcp.Description("")),
		mcp.WithString("type", mcp.Description("")),
		mcp.WithString("id", mcp.Description("")),
		mcp.WithString("location", mcp.Description("")),
		mcp.WithString("name", mcp.Description("")),
		mcp.WithObject("properties", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_api_v1_subscriptions_subscriptionid_resourcegroups_resourcegroup_providers_github_network_resourcetype_resourcename_resourcereadbeginHandler(cfg),
	}
}
