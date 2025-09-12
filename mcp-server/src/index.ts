import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import {
  CallToolRequestSchema,
  ListToolsRequestSchema,
  ListResourcesRequestSchema,
  ReadResourceRequestSchema,
  ErrorCode,
  McpError,
} from '@modelcontextprotocol/sdk/types.js';
import { Octokit } from '@octokit/rest';
import { graphql } from '@octokit/graphql';
import dotenv from 'dotenv';
import { z } from 'zod';

// Load environment variables
dotenv.config();

// Configuration
const config = {
  githubToken: process.env.GITHUB_TOKEN || process.env.GITHUB_PERSONAL_ACCESS_TOKEN || '',
  githubOrg: process.env.GITHUB_ORG || 'vertikon',
  defaultRepo: process.env.GITHUB_DEFAULT_REPO || 'ecosystem',
  serverName: process.env.MCP_SERVER_NAME || 'vertikon-mcp-ultra',
  serverPort: parseInt(process.env.MCP_SERVER_PORT || '3100'),
  enableCache: process.env.ENABLE_CACHE === 'true',
  cacheTTL: parseInt(process.env.CACHE_TTL || '300'),
  githubHost: process.env.GITHUB_HOST || 'github.com',
  githubApiUrl: process.env.GITHUB_API_URL || 'https://api.github.com',
  sshKeyPath: process.env.SSH_KEY_PATH || '',
  readOnly: process.env.GITHUB_READ_ONLY === 'true',
  toolsets: (process.env.GITHUB_TOOLSETS || 'context,repos,issues,pull_requests,actions,code_security').split(','),
};

// Validate configuration
if (!config.githubToken) {
  console.error('ERROR: GITHUB_TOKEN is required');
  process.exit(1);
}

// Initialize GitHub clients
const octokit = new Octokit({
  auth: config.githubToken,
  baseUrl: config.githubApiUrl,
});

const graphqlWithAuth = graphql.defaults({
  headers: {
    authorization: `token ${config.githubToken}`,
  },
  baseUrl: `${config.githubApiUrl}/graphql`,
});

// Initialize MCP Server
const server = new Server(
  {
    name: config.serverName,
    version: '1.0.0',
  },
  {
    capabilities: {
      tools: {},
      resources: {},
    },
  }
);

// Cache implementation
const cache = new Map<string, { data: any; timestamp: number }>();

function getCached(key: string): any | null {
  if (!config.enableCache) return null;
  
  const cached = cache.get(key);
  if (!cached) return null;
  
  const now = Date.now();
  if (now - cached.timestamp > config.cacheTTL * 1000) {
    cache.delete(key);
    return null;
  }
  
  return cached.data;
}

function setCached(key: string, data: any): void {
  if (!config.enableCache) return;
  cache.set(key, { data, timestamp: Date.now() });
}

// Tool Schemas
const CreateIssueSchema = z.object({
  repo: z.string().describe('Repository name (e.g., "vertikon/ecosystem")'),
  title: z.string().describe('Issue title'),
  body: z.string().describe('Issue body (supports Markdown)'),
  labels: z.array(z.string()).optional().describe('Labels to add'),
  assignees: z.array(z.string()).optional().describe('Users to assign'),
});

const CreatePullRequestSchema = z.object({
  repo: z.string().describe('Repository name'),
  title: z.string().describe('PR title'),
  body: z.string().describe('PR description'),
  head: z.string().describe('Branch with changes'),
  base: z.string().default('main').describe('Target branch'),
  draft: z.boolean().optional().describe('Create as draft PR'),
});

const SearchCodeSchema = z.object({
  query: z.string().describe('Search query'),
  repo: z.string().optional().describe('Limit to specific repo'),
  language: z.string().optional().describe('Programming language'),
  path: z.string().optional().describe('File path pattern'),
});

const ListWorkflowRunsSchema = z.object({
  repo: z.string().describe('Repository name'),
  workflow: z.string().optional().describe('Workflow name or ID'),
  branch: z.string().optional().describe('Branch name'),
  status: z.enum(['completed', 'in_progress', 'queued']).optional(),
  limit: z.number().default(10).describe('Number of runs to fetch'),
});

const CreateRepositorySchema = z.object({
  name: z.string().describe('Repository name'),
  description: z.string().optional().describe('Repository description'),
  private: z.boolean().default(false).describe('Create as private repository'),
  auto_init: z.boolean().default(true).describe('Initialize with README'),
  gitignore_template: z.string().optional().describe('Gitignore template (e.g., Node, Go, Python)'),
  license_template: z.string().optional().describe('License template (e.g., mit, apache-2.0)'),
});

// Tool handlers
server.setRequestHandler(ListToolsRequestSchema, async () => {
  return {
    tools: [
      {
        name: 'create_repository',
        description: '🚀 Create a new GitHub repository in the Vertikon organization',
        inputSchema: {
          type: 'object',
          properties: {
            name: { type: 'string' },
            description: { type: 'string' },
            private: { type: 'boolean' },
            auto_init: { type: 'boolean' },
            gitignore_template: { type: 'string' },
            license_template: { type: 'string' },
          },
          required: ['name'],
        },
      },
      {
        name: 'create_issue',
        description: '📝 Create a new GitHub issue in a Vertikon repository',
        inputSchema: {
          type: 'object',
          properties: {
            repo: { type: 'string' },
            title: { type: 'string' },
            body: { type: 'string' },
            labels: { type: 'array', items: { type: 'string' } },
            assignees: { type: 'array', items: { type: 'string' } },
          },
          required: ['repo', 'title', 'body'],
        },
      },
      {
        name: 'create_pull_request',
        description: '🔀 Create a new pull request',
        inputSchema: {
          type: 'object',
          properties: {
            repo: { type: 'string' },
            title: { type: 'string' },
            body: { type: 'string' },
            head: { type: 'string' },
            base: { type: 'string' },
            draft: { type: 'boolean' },
          },
          required: ['repo', 'title', 'body', 'head'],
        },
      },
      {
        name: 'search_code',
        description: '🔍 Search for code in Vertikon repositories',
        inputSchema: {
          type: 'object',
          properties: {
            query: { type: 'string' },
            repo: { type: 'string' },
            language: { type: 'string' },
            path: { type: 'string' },
          },
          required: ['query'],
        },
      },
      {
        name: 'list_workflow_runs',
        description: '⚙️ List GitHub Actions workflow runs',
        inputSchema: {
          type: 'object',
          properties: {
            repo: { type: 'string' },
            workflow: { type: 'string' },
            branch: { type: 'string' },
            status: { type: 'string' },
            limit: { type: 'number' },
          },
          required: ['repo'],
        },
      },
      {
        name: 'get_repo_stats',
        description: '📊 Get repository statistics and insights',
        inputSchema: {
          type: 'object',
          properties: {
            repo: { type: 'string' },
          },
          required: ['repo'],
        },
      },
    ],
  };
});

// Resource handlers
server.setRequestHandler(ListResourcesRequestSchema, async () => {
  return {
    resources: [
      {
        uri: `github://${config.githubOrg}/repositories`,
        name: 'Vertikon Repositories',
        description: 'List of all Vertikon GitHub repositories',
        mimeType: 'application/json',
      },
      {
        uri: `github://${config.githubOrg}/${config.defaultRepo}/readme`,
        name: 'Main README',
        description: 'README of the main Vertikon ecosystem repository',
        mimeType: 'text/markdown',
      },
      {
        uri: `github://${config.githubOrg}/${config.defaultRepo}/issues`,
        name: 'Open Issues',
        description: 'List of open issues in the ecosystem',
        mimeType: 'application/json',
      },
      {
        uri: `github://${config.githubOrg}/${config.defaultRepo}/pulls`,
        name: 'Pull Requests',
        description: 'List of open pull requests',
        mimeType: 'application/json',
      },
    ],
  };
});

server.setRequestHandler(ReadResourceRequestSchema, async (request) => {
  const { uri } = request.params;
  const parts = uri.replace('github://', '').split('/');

  if (parts[1] === 'repositories') {
    const cacheKey = `repos:${config.githubOrg}`;
    let repos = getCached(cacheKey);
    
    if (!repos) {
      const response = await octokit.repos.listForOrg({
        org: config.githubOrg,
        sort: 'updated',
        per_page: 100,
      });
      repos = response.data;
      setCached(cacheKey, repos);
    }

    return {
      contents: [
        {
          uri,
          mimeType: 'application/json',
          text: JSON.stringify(repos, null, 2),
        },
      ],
    };
  }

  if (parts[2] === 'readme') {
    const [owner, repo] = parts.slice(0, 2);
    const readme = await octokit.repos.getReadme({
      owner,
      repo,
    });

    return {
      contents: [
        {
          uri,
          mimeType: 'text/markdown',
          text: Buffer.from(readme.data.content, 'base64').toString('utf-8'),
        },
      ],
    };
  }

  if (parts[2] === 'issues') {
    const [owner, repo] = parts.slice(0, 2);
    const issues = await octokit.issues.listForRepo({
      owner,
      repo,
      state: 'open',
      sort: 'updated',
      per_page: 30,
    });

    return {
      contents: [
        {
          uri,
          mimeType: 'application/json',
          text: JSON.stringify(issues.data, null, 2),
        },
      ],
    };
  }

  if (parts[2] === 'pulls') {
    const [owner, repo] = parts.slice(0, 2);
    const pulls = await octokit.pulls.list({
      owner,
      repo,
      state: 'open',
      sort: 'updated',
      per_page: 30,
    });

    return {
      contents: [
        {
          uri,
          mimeType: 'application/json',
          text: JSON.stringify(pulls.data, null, 2),
        },
      ],
    };
  }

  throw new McpError(ErrorCode.InvalidRequest, `Unknown resource: ${uri}`);
});

server.setRequestHandler(CallToolRequestSchema, async (request) => {
  const { name, arguments: args } = request.params;

  switch (name) {
    case 'create_repository': {
      const params = CreateRepositorySchema.parse(args);
      
      if (config.readOnly) {
        throw new McpError(ErrorCode.MethodNotFound, 'Repository creation disabled in read-only mode');
      }

      try {
        const repo = await octokit.repos.createInOrg({
          org: config.githubOrg,
          name: params.name,
          description: params.description,
          private: params.private,
          auto_init: params.auto_init,
          gitignore_template: params.gitignore_template,
          license_template: params.license_template,
        });

        return {
          content: [
            {
              type: 'text',
              text: `✅ Repository created successfully!\n🔗 URL: ${repo.data.html_url}\n📁 Clone URL: ${repo.data.clone_url}\n🔧 SSH URL: ${repo.data.ssh_url}`,
            },
          ],
        };
      } catch (error: any) {
        if (error.status === 422) {
          throw new McpError(ErrorCode.InvalidRequest, `Repository '${params.name}' already exists or name is invalid`);
        }
        throw new McpError(ErrorCode.InternalError, `Failed to create repository: ${error.message}`);
      }
    }

    case 'create_issue': {
      const params = CreateIssueSchema.parse(args);
      const [owner, repo] = params.repo.split('/');

      const issue = await octokit.issues.create({
        owner,
        repo,
        title: params.title,
        body: params.body,
        labels: params.labels,
        assignees: params.assignees,
      });

      return {
        content: [
          {
            type: 'text',
            text: `✅ Issue created: ${issue.data.html_url}`,
          },
        ],
      };
    }

    case 'create_pull_request': {
      const params = CreatePullRequestSchema.parse(args);
      const [owner, repo] = params.repo.split('/');

      const pr = await octokit.pulls.create({
        owner,
        repo,
        title: params.title,
        body: params.body,
        head: params.head,
        base: params.base || 'main',
        draft: params.draft,
      });

      return {
        content: [
          {
            type: 'text',
            text: `✅ Pull request created: ${pr.data.html_url}`,
          },
        ],
      };
    }

    case 'search_code': {
      const params = SearchCodeSchema.parse(args);
      let query = params.query;

      if (params.repo) {
        query += ` repo:${params.repo}`;
      } else {
        query += ` org:${config.githubOrg}`;
      }

      if (params.language) {
        query += ` language:${params.language}`;
      }

      if (params.path) {
        query += ` path:${params.path}`;
      }

      const results = await octokit.search.code({
        q: query,
        per_page: 20,
      });

      return {
        content: [
          {
            type: 'text',
            text: JSON.stringify(results.data, null, 2),
          },
        ],
      };
    }

    case 'list_workflow_runs': {
      const params = ListWorkflowRunsSchema.parse(args);
      const [owner, repo] = params.repo.split('/');

      const runs = await octokit.actions.listWorkflowRunsForRepo({
        owner,
        repo,
        workflow_id: params.workflow,
        branch: params.branch,
        status: params.status as any,
        per_page: params.limit,
      });

      return {
        content: [
          {
            type: 'text',
            text: JSON.stringify(runs.data, null, 2),
          },
        ],
      };
    }

    case 'get_repo_stats': {
      const { repo } = args as { repo: string };
      const [owner, repoName] = repo.split('/');

      // Use GraphQL for more detailed stats
      const query = `
        query RepoStats($owner: String!, $name: String!) {
          repository(owner: $owner, name: $name) {
            name
            description
            stargazerCount
            forkCount
            issues(states: OPEN) { totalCount }
            pullRequests(states: OPEN) { totalCount }
            releases { totalCount }
            diskUsage
            primaryLanguage { name }
            languages(first: 10) {
              edges {
                node { name }
                size
              }
            }
            defaultBranchRef {
              target {
                ... on Commit {
                  history { totalCount }
                }
              }
            }
          }
        }
      `;

      const stats = await graphqlWithAuth(query, {
        owner,
        name: repoName,
      });

      return {
        content: [
          {
            type: 'text',
            text: JSON.stringify(stats, null, 2),
          },
        ],
      };
    }

    default:
      throw new McpError(ErrorCode.MethodNotFound, `Unknown tool: ${name}`);
  }
});

// Start the server
async function main() {
  const transport = new StdioServerTransport();
  await server.connect(transport);
  
  console.error(`🚀 MCP Ultra Server started on stdio`);
  console.error(`🏢 Organization: ${config.githubOrg}`);
  console.error(`📁 Default repo: ${config.defaultRepo}`);
  console.error(`🔧 Server: ${config.serverName}`);
}

main().catch((error) => {
  console.error('💥 Fatal error:', error);
  process.exit(1);
});