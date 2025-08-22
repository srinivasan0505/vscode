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

function getPut_Api_V1_Prebuilds_Pools_Pool_Id_InstancesHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.poolId !== undefined) {
        queryParams.push(`poolId=${args.poolId}`);
    }
    if (args.environmentOptions !== undefined) {
        queryParams.push(`environmentOptions=${args.environmentOptions}`);
    }
    if (args.secrets !== undefined) {
        queryParams.push(`secrets=${args.secrets}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/put_api_v1_prebuilds_pools_pool_id_instances${queryString}`;
            
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

function createPut_Api_V1_Prebuilds_Pools_Pool_Id_InstancesTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "put_api_v1_prebuilds_pools_pool_id_instances",
            description: "No description available",
            parameters: {
        poolId: { type: "string", required: true, description: "" },
        environmentOptions: { type: "string", required: false, description: "" },
        secrets: { type: "string", required: false, description: "" },
            }
        },
        handler: getPut_Api_V1_Prebuilds_Pools_Pool_Id_InstancesHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPut_Api_V1_Prebuilds_Pools_Pool_Id_InstancesHandler,
    createPut_Api_V1_Prebuilds_Pools_Pool_Id_InstancesTool
};