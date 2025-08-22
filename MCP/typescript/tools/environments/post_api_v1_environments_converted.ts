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

function getPost_Api_V1_EnvironmentsHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.platform !== undefined) {
        queryParams.push(`platform=${args.platform}`);
    }
    if (args.analyticsTrackingId !== undefined) {
        queryParams.push(`analyticsTrackingId=${args.analyticsTrackingId}`);
    }
    if (args.gitHubPfsAuthEndpoint !== undefined) {
        queryParams.push(`gitHubPfsAuthEndpoint=${args.gitHubPfsAuthEndpoint}`);
    }
    if (args.location !== undefined) {
        queryParams.push(`location=${args.location}`);
    }
    if (args.userTier !== undefined) {
        queryParams.push(`userTier=${args.userTier}`);
    }
    if (args.friendlyName !== undefined) {
        queryParams.push(`friendlyName=${args.friendlyName}`);
    }
    if (args.gitHubApiUrl !== undefined) {
        queryParams.push(`gitHubApiUrl=${args.gitHubApiUrl}`);
    }
    if (args.devContainerPath !== undefined) {
        queryParams.push(`devContainerPath=${args.devContainerPath}`);
    }
    if (args.gitHubAppUrl !== undefined) {
        queryParams.push(`gitHubAppUrl=${args.gitHubAppUrl}`);
    }
    if (args.devContainerJson !== undefined) {
        queryParams.push(`devContainerJson=${args.devContainerJson}`);
    }
    if (args.workingDirectory !== undefined) {
        queryParams.push(`workingDirectory=${args.workingDirectory}`);
    }
    if (args.skuName !== undefined) {
        queryParams.push(`skuName=${args.skuName}`);
    }
    if (args.githubEnvironmentEndpoint !== undefined) {
        queryParams.push(`githubEnvironmentEndpoint=${args.githubEnvironmentEndpoint}`);
    }
    if (args.containerImage !== undefined) {
        queryParams.push(`containerImage=${args.containerImage}`);
    }
    if (args.planId !== undefined) {
        queryParams.push(`planId=${args.planId}`);
    }
    if (args.label !== undefined) {
        queryParams.push(`label=${args.label}`);
    }
    if (args.type !== undefined) {
        queryParams.push(`type=${args.type}`);
    }
    if (args.autoShutdownDelayMinutes !== undefined) {
        queryParams.push(`autoShutdownDelayMinutes=${args.autoShutdownDelayMinutes}`);
    }
    if (args.access !== undefined) {
        queryParams.push(`access=${args.access}`);
    }
    if (args.testAccount !== undefined) {
        queryParams.push(`testAccount=${args.testAccount}`);
    }
    if (args.createAsPrebuild !== undefined) {
        queryParams.push(`createAsPrebuild=${args.createAsPrebuild}`);
    }
    if (args.hasDevcontainerJson !== undefined) {
        queryParams.push(`hasDevcontainerJson=${args.hasDevcontainerJson}`);
    }
    if (args.features !== undefined) {
        queryParams.push(`features=${args.features}`);
    }
    if (args.personalization !== undefined) {
        queryParams.push(`personalization=${args.personalization}`);
    }
    if (args.billableOwner !== undefined) {
        queryParams.push(`billableOwner=${args.billableOwner}`);
    }
    if (args.identity !== undefined) {
        queryParams.push(`identity=${args.identity}`);
    }
    if (args.experimentalFeatures !== undefined) {
        queryParams.push(`experimentalFeatures=${args.experimentalFeatures}`);
    }
    if (args.runtimeConstraints !== undefined) {
        queryParams.push(`runtimeConstraints=${args.runtimeConstraints}`);
    }
    if (args.netmonCorrelationData !== undefined) {
        queryParams.push(`netmonCorrelationData=${args.netmonCorrelationData}`);
    }
    if (args.seed !== undefined) {
        queryParams.push(`seed=${args.seed}`);
    }
    if (args.connection !== undefined) {
        queryParams.push(`connection=${args.connection}`);
    }
    if (args.secrets !== undefined) {
        queryParams.push(`secrets=${args.secrets}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_api_v1_environments${queryString}`;
            
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

function createPost_Api_V1_EnvironmentsTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_api_v1_environments",
            description: "No description available",
            parameters: {
        platform: { type: "string", required: false, description: "" },
        analyticsTrackingId: { type: "string", required: false, description: "" },
        gitHubPfsAuthEndpoint: { type: "string", required: false, description: "" },
        location: { type: "string", required: false, description: "" },
        userTier: { type: "string", required: false, description: "" },
        friendlyName: { type: "string", required: true, description: "" },
        gitHubApiUrl: { type: "string", required: false, description: "" },
        devContainerPath: { type: "string", required: false, description: "" },
        gitHubAppUrl: { type: "string", required: false, description: "" },
        devContainerJson: { type: "string", required: false, description: "" },
        workingDirectory: { type: "string", required: false, description: "" },
        skuName: { type: "string", required: false, description: "" },
        githubEnvironmentEndpoint: { type: "string", required: false, description: "" },
        containerImage: { type: "string", required: false, description: "" },
        planId: { type: "string", required: false, description: "" },
        label: { type: "string", required: false, description: "" },
        type: { type: "string", required: true, description: "" },
        autoShutdownDelayMinutes: { type: "string", required: false, description: "" },
        access: { type: "string", required: false, description: "" },
        testAccount: { type: "string", required: false, description: "" },
        createAsPrebuild: { type: "string", required: false, description: "" },
        hasDevcontainerJson: { type: "string", required: false, description: "" },
        features: { type: "string", required: false, description: "" },
        personalization: { type: "string", required: false, description: "" },
        billableOwner: { type: "string", required: false, description: "" },
        identity: { type: "string", required: false, description: "" },
        experimentalFeatures: { type: "string", required: false, description: "" },
        runtimeConstraints: { type: "string", required: false, description: "" },
        netmonCorrelationData: { type: "string", required: false, description: "" },
        seed: { type: "string", required: false, description: "" },
        connection: { type: "string", required: false, description: "" },
        secrets: { type: "string", required: false, description: "" },
            }
        },
        handler: getPost_Api_V1_EnvironmentsHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_Api_V1_EnvironmentsHandler,
    createPost_Api_V1_EnvironmentsTool
};