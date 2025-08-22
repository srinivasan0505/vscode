package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/vsonline/mcp-server/config"
	"github.com/vsonline/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_api_v1_tokens_plans_planname_deleteallcodespacesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		planNameVal, ok := args["planName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: planName"), nil
		}
		planName, ok := planNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: planName"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["expiration"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("expiration=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/v1/Tokens/plans/%s/deleteAllCodespaces%s", cfg.BaseURL, planName, queryString)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-subscription-id"]; ok {
			req.Header.Set("x-subscription-id", fmt.Sprintf("%v", val))
		}

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

func CreatePost_api_v1_tokens_plans_planname_deleteallcodespacesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_api_v1_Tokens_plans_planName_deleteAllCodespaces",
		mcp.WithDescription(""),
		mcp.WithString("planName", mcp.Required(), mcp.Description("")),
		mcp.WithString("x-subscription-id", mcp.Description("")),
		mcp.WithString("expiration", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_api_v1_tokens_plans_planname_deleteallcodespacesHandler(cfg),
	}
}
