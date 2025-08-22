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

func GetprebuildreadinessskusrouteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		prebuildHashVal, ok := args["prebuildHash"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: prebuildHash"), nil
		}
		prebuildHash, ok := prebuildHashVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: prebuildHash"), nil
		}
		locationVal, ok := args["location"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: location"), nil
		}
		location, ok := locationVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: location"), nil
		}
		devContainerPathVal, ok := args["devContainerPath"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: devContainerPath"), nil
		}
		devContainerPath, ok := devContainerPathVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: devContainerPath"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["storageType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("storageType=%v", val))
		}
		if val, ok := args["fastPathEnabled"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fastPathEnabled=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/v2/prebuilds/templates/skus/repo/%s/branch/%s/hash/%s/location/%s/devcontainerpath/%s%s", cfg.BaseURL, repoId, branchName, prebuildHash, location, devContainerPath, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.PrebuildReadinessResult
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

func CreateGetprebuildreadinessskusrouteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_v2_prebuilds_templates_skus_repo_repoId_branch_branchName_hash_prebuildHash_location_location_devcontainerpath_devContainerPath",
		mcp.WithDescription(""),
		mcp.WithString("repoId", mcp.Required(), mcp.Description("")),
		mcp.WithString("branchName", mcp.Required(), mcp.Description("")),
		mcp.WithString("prebuildHash", mcp.Required(), mcp.Description("")),
		mcp.WithString("location", mcp.Required(), mcp.Description("")),
		mcp.WithString("devContainerPath", mcp.Required(), mcp.Description("")),
		mcp.WithNumber("storageType", mcp.Description("")),
		mcp.WithBoolean("fastPathEnabled", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetprebuildreadinessskusrouteHandler(cfg),
	}
}
