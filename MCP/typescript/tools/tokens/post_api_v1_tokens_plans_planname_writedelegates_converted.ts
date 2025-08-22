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

function getPost_Api_V1_Tokens_Plans_Plan_Name_Write_DelegatesHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.planName !== undefined) {
        queryParams.push(`planName=${args.planName}`);
    }
    if (args.x-subscription-id !== undefined) {
        queryParams.push(`x-subscription-id=${args.x-subscription-id}`);
    }
    if (args.scope !== undefined) {
        queryParams.push(`scope=${args.scope}`);
    }
    if (args.expiration !== undefined) {
        queryParams.push(`expiration=${args.expiration}`);
    }
    if (args.identity !== undefined) {
        queryParams.push(`identity=${args.identity}`);
    }
    if (args.portNumbers !== undefined) {
        queryParams.push(`portNumbers=${args.portNumbers}`);
    }
    if (args.environmentIds !== undefined) {
        queryParams.push(`environmentIds=${args.environmentIds}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_tokens_plans_plan_name_write_delegates${queryString}`;
            
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

function createPost_Api_V1_Tokens_Plans_Plan_Name_Write_DelegatesTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_tokens_plans_plan_name_write_delegates",
            description: "No description available",
            parameters: {
        planName: { type: "string", required: true, description: "" },
        x-subscription-id: { type: "string", required: false, description: "" },
        scope: { type: "string", required: false, description: "" },
        expiration: { type: "string", required: false, description: "" },
        identity: { type: "string", required: false, description: "" },
        portNumbers: { type: "string", required: false, description: "" },
        environmentIds: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_Tokens_Plans_Plan_Name_Write_DelegatesHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_Tokens_Plans_Plan_Name_Write_DelegatesHandler,
    createPost_Api_V1_Tokens_Plans_Plan_Name_Write_DelegatesTool
};