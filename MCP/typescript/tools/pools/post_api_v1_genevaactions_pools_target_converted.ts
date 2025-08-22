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

function getPost_Api_V1_Geneva_Actions_Pools_TargetHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.target !== undefined) {
        queryParams.push(`target=${args.target}`);
    }
    if (args.targetCount !== undefined) {
        queryParams.push(`targetCount=${args.targetCount}`);
    }
    if (args.comment !== undefined) {
        queryParams.push(`comment=${args.comment}`);
    }
    if (args.maxTargetCount !== undefined) {
        queryParams.push(`maxTargetCount=${args.maxTargetCount}`);
    }
    if (args.minTargetCount !== undefined) {
        queryParams.push(`minTargetCount=${args.minTargetCount}`);
    }
    if (args.poolCode !== undefined) {
        queryParams.push(`poolCode=${args.poolCode}`);
    }
    if (args.poolType !== undefined) {
        queryParams.push(`poolType=${args.poolType}`);
    }
    if (args.region !== undefined) {
        queryParams.push(`region=${args.region}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_geneva_actions_pools_target${queryString}`;
            
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

function createPost_Api_V1_Geneva_Actions_Pools_TargetTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_geneva_actions_pools_target",
            description: "No description available",
            parameters: {
        target: { type: "string", required: true, description: "" },
        targetCount: { type: "string", required: false, description: "" },
        comment: { type: "string", required: false, description: "" },
        maxTargetCount: { type: "string", required: false, description: "" },
        minTargetCount: { type: "string", required: false, description: "" },
        poolCode: { type: "string", required: false, description: "" },
        poolType: { type: "string", required: false, description: "" },
        region: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_Geneva_Actions_Pools_TargetHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_Geneva_Actions_Pools_TargetHandler,
    createPost_Api_V1_Geneva_Actions_Pools_TargetTool
};