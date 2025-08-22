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

function getPost_Api_V2_Prebuilds_TemplatesHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.devContainerPath !== undefined) {
        queryParams.push(`devContainerPath=${args.devContainerPath}`);
    }
    if (args.planId !== undefined) {
        queryParams.push(`planId=${args.planId}`);
    }
    if (args.gitHubPrebuildTemplateEndpoint !== undefined) {
        queryParams.push(`gitHubPrebuildTemplateEndpoint=${args.gitHubPrebuildTemplateEndpoint}`);
    }
    if (args.friendlyName !== undefined) {
        queryParams.push(`friendlyName=${args.friendlyName}`);
    }
    if (args.gitHubPrebuildInstanceEndpoint !== undefined) {
        queryParams.push(`gitHubPrebuildInstanceEndpoint=${args.gitHubPrebuildInstanceEndpoint}`);
    }
    if (args.storageType !== undefined) {
        queryParams.push(`storageType=${args.storageType}`);
    }
    if (args.seed !== undefined) {
        queryParams.push(`seed=${args.seed}`);
    }
    if (args.templateInfo !== undefined) {
        queryParams.push(`templateInfo=${args.templateInfo}`);
    }
    if (args.features !== undefined) {
        queryParams.push(`features=${args.features}`);
    }
    if (args.experimentalFeatures !== undefined) {
        queryParams.push(`experimentalFeatures=${args.experimentalFeatures}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v2_prebuilds_templates${queryString}`;
            
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

function createPost_Api_V2_Prebuilds_TemplatesTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v2_prebuilds_templates",
            description: "No description available",
            parameters: {
        devContainerPath: { type: "string", required: false, description: "" },
        planId: { type: "string", required: false, description: "" },
        gitHubPrebuildTemplateEndpoint: { type: "string", required: false, description: "" },
        friendlyName: { type: "string", required: true, description: "" },
        gitHubPrebuildInstanceEndpoint: { type: "string", required: false, description: "" },
        storageType: { type: "string", required: false, description: "" },
        seed: { type: "string", required: false, description: "" },
        templateInfo: { type: "string", required: false, description: "" },
        features: { type: "string", required: false, description: "" },
        experimentalFeatures: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V2_Prebuilds_TemplatesHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V2_Prebuilds_TemplatesHandler,
    createPost_Api_V2_Prebuilds_TemplatesTool
};