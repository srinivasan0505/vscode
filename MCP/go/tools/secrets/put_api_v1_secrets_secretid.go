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

func Put_api_v1_secrets_secretidHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		secretIdVal, ok := args["secretId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: secretId"), nil
		}
		secretId, ok := secretIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: secretId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["planId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("planId=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ScopedUpdateSecretBody
		
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
		url := fmt.Sprintf("%s/api/v1/Secrets/%s%s", cfg.BaseURL, secretId, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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
		var result models.ScopedSecretResultBody
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

func CreatePut_api_v1_secrets_secretidTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_api_v1_Secrets_secretId",
		mcp.WithDescription(""),
		mcp.WithString("planId", mcp.Description("")),
		mcp.WithString("secretId", mcp.Required(), mcp.Description("")),
		mcp.WithString("secretName", mcp.Description("")),
		mcp.WithString("value", mcp.Description("")),
		mcp.WithArray("filters", mcp.Description("")),
		mcp.WithString("notes", mcp.Description("")),
		mcp.WithNumber("scope", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Put_api_v1_secrets_secretidHandler(cfg),
	}
}
