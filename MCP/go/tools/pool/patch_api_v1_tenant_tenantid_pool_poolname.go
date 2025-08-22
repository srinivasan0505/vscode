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

func Patch_api_v1_tenant_tenantid_pool_poolnameHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		tenantIdVal, ok := args["tenantId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: tenantId"), nil
		}
		tenantId, ok := tenantIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: tenantId"), nil
		}
		poolNameVal, ok := args["poolName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: poolName"), nil
		}
		poolName, ok := poolNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: poolName"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateOrUpdatePoolBody
		
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
		url := fmt.Sprintf("%s/api/v1/tenant/%s/Pool/%s", cfg.BaseURL, tenantId, poolName)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.PoolResult
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

func CreatePatch_api_v1_tenant_tenantid_pool_poolnameTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_api_v1_tenant_tenantId_Pool_poolName",
		mcp.WithDescription(""),
		mcp.WithString("tenantId", mcp.Required(), mcp.Description("")),
		mcp.WithString("poolName", mcp.Required(), mcp.Description("")),
		mcp.WithObject("hotPoolSettings", mcp.Description("")),
		mcp.WithString("poolGroupName", mcp.Required(), mcp.Description("")),
		mcp.WithObject("tags", mcp.Description("")),
		mcp.WithString("userGroupName", mcp.Description("")),
		mcp.WithObject("vmSpecs", mcp.Required(), mcp.Description("")),
		mcp.WithObject("domainUserCredentials", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_api_v1_tenant_tenantid_pool_poolnameHandler(cfg),
	}
}
