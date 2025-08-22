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

function getPost_Api_V1_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Microsoft_Vs_Online_Plans_Resource_Name_Resource_Creation_CompletedHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
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
    if (args.resourceName !== undefined) {
        queryParams.push(`resourceName=${args.resourceName}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_subscriptions_subscription_id_resource_groups_resource_group_providers_microsoft_vs_online_plans_resource_name_resource_creation_completed${queryString}`;
            
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

function createPost_Api_V1_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Microsoft_Vs_Online_Plans_Resource_Name_Resource_Creation_CompletedTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_subscriptions_subscription_id_resource_groups_resource_group_providers_microsoft_vs_online_plans_resource_name_resource_creation_completed",
            description: "No description available",
            parameters: {
        subscriptionId: { type: "string", required: true, description: "" },
        resourceGroup: { type: "string", required: true, description: "" },
        resourceName: { type: "string", required: true, description: "" },
            }
        },
        handler: getPost_Api_V1_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Microsoft_Vs_Online_Plans_Resource_Name_Resource_Creation_CompletedHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Microsoft_Vs_Online_Plans_Resource_Name_Resource_Creation_CompletedHandler,
    createPost_Api_V1_Subscriptions_Subscription_Id_Resource_Groups_Resource_Group_Providers_Microsoft_Vs_Online_Plans_Resource_Name_Resource_Creation_CompletedTool
};