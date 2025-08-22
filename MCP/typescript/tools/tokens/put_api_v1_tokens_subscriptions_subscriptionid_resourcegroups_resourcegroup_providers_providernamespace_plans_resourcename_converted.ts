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

function getPut_Api_V1_Tokens_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Provider_Namespace_Plans_Resource_NameHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.subscriptionId !== undefined) {
        queryParams.push(`subscriptionId=${args.subscriptionId}`);
    }
    if (args.resourceGroup !== undefined) {
        queryParams.push(`resourceGroup=${args.resourceGroup}`);
    }
    if (args.providerNamespace !== undefined) {
        queryParams.push(`providerNamespace=${args.providerNamespace}`);
    }
    if (args.resourceName !== undefined) {
        queryParams.push(`resourceName=${args.resourceName}`);
    }
    if (args.provisioningState !== undefined) {
        queryParams.push(`provisioningState=${args.provisioningState}`);
    }
    if (args.type !== undefined) {
        queryParams.push(`type=${args.type}`);
    }
    if (args.id !== undefined) {
        queryParams.push(`id=${args.id}`);
    }
    if (args.location !== undefined) {
        queryParams.push(`location=${args.location}`);
    }
    if (args.name !== undefined) {
        queryParams.push(`name=${args.name}`);
    }
    if (args.headers !== undefined) {
        queryParams.push(`headers=${args.headers}`);
    }
    if (args.properties !== undefined) {
        queryParams.push(`properties=${args.properties}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
    if (args.identity !== undefined) {
        queryParams.push(`identity=${args.identity}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/put_api_v1_tokens_subscriptions_subscription_id_resource_groups_resource_group_providers_provider_namespace_plans_resource_name${queryString}`;
            
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

function createPut_Api_V1_Tokens_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Provider_Namespace_Plans_Resource_NameTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "put_api_v1_tokens_subscriptions_subscription_id_resource_groups_resource_group_providers_provider_namespace_plans_resource_name",
            description: "No description available",
            parameters: {
        subscriptionId: { type: "string", required: true, description: "" },
        resourceGroup: { type: "string", required: true, description: "" },
        providerNamespace: { type: "string", required: true, description: "" },
        resourceName: { type: "string", required: true, description: "" },
        provisioningState: { type: "string", required: false, description: "" },
        type: { type: "string", required: false, description: "" },
        id: { type: "string", required: false, description: "" },
        location: { type: "string", required: false, description: "" },
        name: { type: "string", required: false, description: "" },
        headers: { type: "string", required: false, description: "" },
        properties: { type: "string", required: false, description: "" },
        tags: { type: "string", required: false, description: "" },
        identity: { type: "string", required: false, description: "" },
            }
        },
        handler: getPut_Api_V1_Tokens_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Provider_Namespace_Plans_Resource_NameHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPut_Api_V1_Tokens_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Provider_Namespace_Plans_Resource_NameHandler,
    createPut_Api_V1_Tokens_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Provider_Namespace_Plans_Resource_NameTool
};