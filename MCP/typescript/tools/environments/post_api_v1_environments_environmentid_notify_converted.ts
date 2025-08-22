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

function getPost_Api_V1_Environments_Environment_Id_NotifyHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.environmentId !== undefined) {
        queryParams.push(`environmentId=${args.environmentId}`);
    }
    if (args.displayMode !== undefined) {
        queryParams.push(`displayMode=${args.displayMode}`);
    }
    if (args.message !== undefined) {
        queryParams.push(`message=${args.message}`);
    }
    if (args.details !== undefined) {
        queryParams.push(`details=${args.details}`);
    }
    if (args.modal !== undefined) {
        queryParams.push(`modal=${args.modal}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_environments_environment_id_notify${queryString}`;
            
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

function createPost_Api_V1_Environments_Environment_Id_NotifyTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_environments_environment_id_notify",
            description: "No description available",
            parameters: {
        environmentId: { type: "string", required: true, description: "" },
        displayMode: { type: "string", required: false, description: "" },
        message: { type: "string", required: false, description: "" },
        details: { type: "string", required: false, description: "" },
        modal: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_Environments_Environment_Id_NotifyHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_Environments_Environment_Id_NotifyHandler,
    createPost_Api_V1_Environments_Environment_Id_NotifyTool
};