/**
 * MCP Server function for No description available
 */

import axios, { AxiosResponse } from 'axios';

interface APIConfig {
    baseUrl: string;
    apiKey: string;
}

interface MCPRequest {
    params?: {
        arguments?: Record<string, any>;
    };
}

interface MCPToolResult {
    content: string;
    isError: boolean;
}

interface ToolDefinition {
    name: string;
    description: string;
    parameters: Record<string, {
        type: string;
        required: boolean;
        description: string;
    }>;
}

interface Tool {
    definition: ToolDefinition;
    handler: (ctx: any, request: MCPRequest) => Promise<MCPToolResult>;
}

class MCPToolResultImpl implements MCPToolResult {
    constructor(
        public content: string,
        public isError: boolean = false
    ) {}
}

function getPost_Api_V1_SecretsHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.planId !== undefined) {
        queryParams.push(`planId=${args.planId}`);
    }
    if (args.value !== undefined) {
        queryParams.push(`value=${args.value}`);
    }
    if (args.notes !== undefined) {
        queryParams.push(`notes=${args.notes}`);
    }
    if (args.secretName !== undefined) {
        queryParams.push(`secretName=${args.secretName}`);
    }
    if (args.type !== undefined) {
        queryParams.push(`type=${args.type}`);
    }
    if (args.scope !== undefined) {
        queryParams.push(`scope=${args.scope}`);
    }
    if (args.filters !== undefined) {
        queryParams.push(`filters=${args.filters}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_secrets${queryString}`;
            
            const headers = {
                'Authorization': `Bearer ${config.apiKey}`,
                'Accept': 'application/json'
            };
            
            const response: AxiosResponse = await axios.get(url, { headers });
            
            if (response.status >= 400) {
                return new MCPToolResultImpl(`API error: ${response.data}`, true);
            }
            
            const prettyJSON = JSON.stringify(response.data, null, 2);
            return new MCPToolResultImpl(prettyJSON);
            
        } catch (error: any) {
            if (error.response) {
                return new MCPToolResultImpl(`Request failed: ${error.response.data}`, true);
            }
            return new MCPToolResultImpl(`Unexpected error: ${error.message}`, true);
        }
    };
}

function createPost_Api_V1_SecretsTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_secrets",
            description: "No description available",
            parameters: {
        planId: { type: "string", required: false, description: "" },
        value: { type: "string", required: false, description: "" },
        notes: { type: "string", required: false, description: "" },
        secretName: { type: "string", required: false, description: "" },
        type: { type: "string", required: false, description: "" },
        scope: { type: "string", required: false, description: "" },
        filters: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_SecretsHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_SecretsHandler,
    createPost_Api_V1_SecretsTool
};