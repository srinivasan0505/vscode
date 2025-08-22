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

function getPost_Api_V1_Geneva_Actions_Vnet_Pool_DefinitionsHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.subtype !== undefined) {
        queryParams.push(`subtype=${args.subtype}`);
    }
    if (args.targetCount !== undefined) {
        queryParams.push(`targetCount=${args.targetCount}`);
    }
    if (args.type !== undefined) {
        queryParams.push(`type=${args.type}`);
    }
    if (args.location !== undefined) {
        queryParams.push(`location=${args.location}`);
    }
    if (args.isEnabled !== undefined) {
        queryParams.push(`isEnabled=${args.isEnabled}`);
    }
    if (args.dimensions !== undefined) {
        queryParams.push(`dimensions=${args.dimensions}`);
    }
    if (args.logicalSkus !== undefined) {
        queryParams.push(`logicalSkus=${args.logicalSkus}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_geneva_actions_vnet_pool_definitions${queryString}`;
            
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

function createPost_Api_V1_Geneva_Actions_Vnet_Pool_DefinitionsTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_geneva_actions_vnet_pool_definitions",
            description: "No description available",
            parameters: {
        subtype: { type: "string", required: true, description: "" },
        targetCount: { type: "string", required: true, description: "" },
        type: { type: "string", required: true, description: "" },
        location: { type: "string", required: true, description: "" },
        isEnabled: { type: "string", required: true, description: "" },
        dimensions: { type: "string", required: true, description: "" },
        logicalSkus: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_Geneva_Actions_Vnet_Pool_DefinitionsHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_Geneva_Actions_Vnet_Pool_DefinitionsHandler,
    createPost_Api_V1_Geneva_Actions_Vnet_Pool_DefinitionsTool
};