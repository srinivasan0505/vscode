package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/vsonline/mcp-server/config"
	"github.com/vsonline/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_api_v1_environmentsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["access"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("access=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateCloudEnvironmentBody
		
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
		url := fmt.Sprintf("%s/api/v1/Environments%s", cfg.BaseURL, queryString)
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
		var result models.CloudEnvironmentResult
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

func CreatePost_api_v1_environmentsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_v1_Environments",
		mcp.WithDescription(""),
		mcp.WithBoolean("access", mcp.Description("")),
		mcp.WithString("platform", mcp.Description("")),
		mcp.WithObject("features", mcp.Description("")),
		mcp.WithString("analyticsTrackingId", mcp.Description("")),
		mcp.WithNumber("autoShutdownDelayMinutes", mcp.Description("")),
		mcp.WithObject("personalization", mcp.Description("")),
		mcp.WithBoolean("testAccount", mcp.Description("")),
		mcp.WithBoolean("createAsPrebuild", mcp.Description("")),
		mcp.WithString("gitHubPfsAuthEndpoint", mcp.Description("")),
		mcp.WithString("location", mcp.Description("")),
		mcp.WithString("userTier", mcp.Description("")),
		mcp.WithString("friendlyName", mcp.Required(), mcp.Description("")),
		mcp.WithString("gitHubApiUrl", mcp.Description("")),
		mcp.WithObject("billableOwner", mcp.Description("")),
		mcp.WithBoolean("hasDevcontainerJson", mcp.Description("")),
		mcp.WithString("devContainerPath", mcp.Description("")),
		mcp.WithString("gitHubAppUrl", mcp.Description("")),
		mcp.WithObject("identity", mcp.Description("")),
		mcp.WithString("devContainerJson", mcp.Description("")),
		mcp.WithString("workingDirectory", mcp.Description("")),
		mcp.WithString("skuName", mcp.Description("")),
		mcp.WithObject("experimentalFeatures", mcp.Description("")),
		mcp.WithObject("runtimeConstraints", mcp.Description("")),
		mcp.WithString("githubEnvironmentEndpoint", mcp.Description("")),
		mcp.WithString("containerImage", mcp.Description("")),
		mcp.WithString("planId", mcp.Description("")),
		mcp.WithString("label", mcp.Description("")),
		mcp.WithObject("netmonCorrelationData", mcp.Description("")),
		mcp.WithArray("secrets", mcp.Description("")),
		mcp.WithObject("seed", mcp.Description("")),
		mcp.WithString("type", mcp.Required(), mcp.Description("")),
		mcp.WithObject("connection", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_api_v1_environmentsHandler(cfg),
	}
}
