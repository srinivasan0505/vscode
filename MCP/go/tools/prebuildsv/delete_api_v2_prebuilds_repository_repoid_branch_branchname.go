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

func Delete_api_v2_prebuilds_repository_repoid_branch_branchnameHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		repoIdVal, ok := args["repoId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: repoId"), nil
		}
		repoId, ok := repoIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: repoId"), nil
		}
		branchNameVal, ok := args["branchName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: branchName"), nil
		}
		branchName, ok := branchNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: branchName"), nil
		}
		url := fmt.Sprintf("%s/api/v2/prebuilds/repository/%s/branch/%s", cfg.BaseURL, repoId, branchName)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateDelete_api_v2_prebuilds_repository_repoid_branch_branchnameTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_api_v2_prebuilds_repository_repoId_branch_branchName",
		mcp.WithDescription(""),
		mcp.WithNumber("repoId", mcp.Required(), mcp.Description("")),
		mcp.WithString("branchName", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_api_v2_prebuilds_repository_repoid_branch_branchnameHandler(cfg),
	}
}
