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

function getPatch_Api_V1_Tenant_Tenant_Id_Pool_Group_Pool_Group_NameHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.tenantId !== undefined) {
        queryParams.push(`tenantId=${args.tenantId}`);
    }
    if (args.poolGroupName !== undefined) {
        queryParams.push(`poolGroupName=${args.poolGroupName}`);
    }
    if (args.displayName !== undefined) {
        queryParams.push(`displayName=${args.displayName}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/patch_api_v1_tenant_tenant_id_pool_group_pool_group_name${queryString}`;
            
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

function createPatch_Api_V1_Tenant_Tenant_Id_Pool_Group_Pool_Group_NameTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "patch_api_v1_tenant_tenant_id_pool_group_pool_group_name",
            description: "No description available",
            parameters: {
        tenantId: { type: "string", required: true, description: "" },
        poolGroupName: { type: "string", required: true, description: "" },
        displayName: { type: "string", required: true, description: "" },
        tags: { type: "string", required: false, description: "" },
            }
        },
        handler: getPatch_Api_V1_Tenant_Tenant_Id_Pool_Group_Pool_Group_NameHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPatch_Api_V1_Tenant_Tenant_Id_Pool_Group_Pool_Group_NameHandler,
    createPatch_Api_V1_Tenant_Tenant_Id_Pool_Group_Pool_Group_NameTool
};