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

func Get_api_v1_tunnel_environmentid_portinfoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		environmentIdVal, ok := args["environmentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: environmentId"), nil
		}
		environmentId, ok := environmentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: environmentId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["portNumber"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("portNumber=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/api/v1/Tunnel/%s/portInfo%s", cfg.BaseURL, environmentId, queryString)
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
		var result models.TunnelPortInfoResult
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

func CreateGet_api_v1_tunnel_environmentid_portinfoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_api_v1_Tunnel_environmentId_portInfo",
		mcp.WithDescription(""),
		mcp.WithString("environmentId", mcp.Required(), mcp.Description("")),
		mcp.WithNumber("portNumber", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_api_v1_tunnel_environmentid_portinfoHandler(cfg),
	}
}
