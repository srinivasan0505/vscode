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

function getGet_Api_V2_Prebuilds_Templates_Skus_Repo_Repo_Id_Branch_Branch_Name_Hash_Prebuild_Hash_Location_Location_Devcontainerpath_Dev_Container_PathHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.repoId !== undefined) {
        queryParams.push(`repoId=${args.repoId}`);
    }
    if (args.branchName !== undefined) {
        queryParams.push(`branchName=${args.branchName}`);
    }
    if (args.prebuildHash !== undefined) {
        queryParams.push(`prebuildHash=${args.prebuildHash}`);
    }
    if (args.location !== undefined) {
        queryParams.push(`location=${args.location}`);
    }
    if (args.devContainerPath !== undefined) {
        queryParams.push(`devContainerPath=${args.devContainerPath}`);
    }
    if (args.storageType !== undefined) {
        queryParams.push(`storageType=${args.storageType}`);
    }
    if (args.fastPathEnabled !== undefined) {
        queryParams.push(`fastPathEnabled=${args.fastPathEnabled}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/get_api_v2_prebuilds_templates_skus_repo_repo_id_branch_branch_name_hash_prebuild_hash_location_location_devcontainerpath_dev_container_path${queryString}`;
            
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

function createGet_Api_V2_Prebuilds_Templates_Skus_Repo_Repo_Id_Branch_Branch_Name_Hash_Prebuild_Hash_Location_Location_Devcontainerpath_Dev_Container_PathTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "get_api_v2_prebuilds_templates_skus_repo_repo_id_branch_branch_name_hash_prebuild_hash_location_location_devcontainerpath_dev_container_path",
            description: "No description available",
            parameters: {
        repoId: { type: "string", required: true, description: "" },
        branchName: { type: "string", required: true, description: "" },
        prebuildHash: { type: "string", required: true, description: "" },
        location: { type: "string", required: true, description: "" },
        devContainerPath: { type: "string", required: true, description: "" },
        storageType: { type: "string", required: false, description: "" },
        fastPathEnabled: { type: "string", required: false, description: "" },
            }
        },
        handler: getGet_Api_V2_Prebuilds_Templates_Skus_Repo_Repo_Id_Branch_Branch_Name_Hash_Prebuild_Hash_Location_Location_Devcontainerpath_Dev_Container_PathHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getGet_Api_V2_Prebuilds_Templates_Skus_Repo_Repo_Id_Branch_Branch_Name_Hash_Prebuild_Hash_Location_Location_Devcontainerpath_Dev_Container_PathHandler,
    createGet_Api_V2_Prebuilds_Templates_Skus_Repo_Repo_Id_Branch_Branch_Name_Hash_Prebuild_Hash_Location_Location_Devcontainerpath_Dev_Container_PathTool
};