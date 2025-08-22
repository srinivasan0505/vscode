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

function getPut_Api_V1_Tenant_Tenant_Id_Pool_Pool_NameHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
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
    if (args.poolName !== undefined) {
        queryParams.push(`poolName=${args.poolName}`);
    }
    if (args.userGroupName !== undefined) {
        queryParams.push(`userGroupName=${args.userGroupName}`);
    }
    if (args.poolGroupName !== undefined) {
        queryParams.push(`poolGroupName=${args.poolGroupName}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
    if (args.vmSpecs !== undefined) {
        queryParams.push(`vmSpecs=${args.vmSpecs}`);
    }
    if (args.domainUserCredentials !== undefined) {
        queryParams.push(`domainUserCredentials=${args.domainUserCredentials}`);
    }
    if (args.hotPoolSettings !== undefined) {
        queryParams.push(`hotPoolSettings=${args.hotPoolSettings}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/put_api_v1_tenant_tenant_id_pool_pool_name${queryString}`;
            
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

function createPut_Api_V1_Tenant_Tenant_Id_Pool_Pool_NameTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "put_api_v1_tenant_tenant_id_pool_pool_name",
            description: "No description available",
            parameters: {
        tenantId: { type: "string", required: true, description: "" },
        poolName: { type: "string", required: true, description: "" },
        userGroupName: { type: "string", required: false, description: "" },
        poolGroupName: { type: "string", required: true, description: "" },
        tags: { type: "string", required: false, description: "" },
        vmSpecs: { type: "string", required: true, description: "" },
        domainUserCredentials: { type: "string", required: false, description: "" },
        hotPoolSettings: { type: "string", required: false, description: "" },
            }
        },
        handler: getPut_Api_V1_Tenant_Tenant_Id_Pool_Pool_NameHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPut_Api_V1_Tenant_Tenant_Id_Pool_Pool_NameHandler,
    createPut_Api_V1_Tenant_Tenant_Id_Pool_Pool_NameTool
};