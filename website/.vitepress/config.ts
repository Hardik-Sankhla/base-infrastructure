import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Base Infrastructure',
  description: 'Cross-platform capability-aware bootstrap framework',
  base: '/base-infrastructure/',

  head: [
    ['meta', { name: 'theme-color', content: '#3b82f6' }],
    ['meta', { property: 'og:type', content: 'website' }],
    ['meta', { property: 'og:title', content: 'Base Infrastructure' }],
    ['meta', { property: 'og:description', content: 'Cross-platform capability-aware bootstrap framework' }],
  ],

  themeConfig: {
    logo: '/logo.svg',
    siteTitle: 'Base Infrastructure',

    nav: [
      { text: 'Getting Started', link: '/getting-started/install' },
      { text: 'Architecture', link: '/architecture/overview' },
      { text: 'CLI Reference', link: '/cli/reference' },
      { text: 'AI Context', link: '/ai/AI_CONTEXT' },
      {
        text: 'v0.1.0',
        items: [
          { text: 'Changelog', link: '/changelog' },
          { text: 'Roadmap', link: '/roadmap' },
          { text: 'GitHub', link: 'https://github.com/Hardik-Sankhla/base-infrastructure' },
        ]
      }
    ],

    sidebar: {
      '/getting-started/': [
        {
          text: 'Getting Started',
          items: [
            { text: 'Installation', link: '/getting-started/install' },
            { text: 'Quickstart', link: '/getting-started/quickstart' },
          ]
        }
      ],

      '/architecture/': [
        {
          text: 'Architecture',
          items: [
            { text: 'Overview', link: '/architecture/overview' },
            { text: 'Diagrams', link: '/architecture/diagrams' },
          ]
        }
      ],

      '/concepts/': [
        {
          text: 'Concepts',
          items: [
            { text: 'Capabilities', link: '/concepts/capabilities' },
            { text: 'Discovery Manifest', link: '/concepts/manifest' },
            { text: 'Glossary', link: '/concepts/glossary' },
          ]
        }
      ],

      '/discovery/': [
        {
          text: 'Discovery Engine',
          items: [
            { text: 'How It Works', link: '/discovery/engine' },
            { text: 'Stages', link: '/discovery/stages' },
            { text: 'Validation', link: '/discovery/validation' },
          ]
        }
      ],

      '/platform/': [
        {
          text: 'Platform Abstraction',
          items: [
            { text: 'Overview', link: '/platform/abstraction' },
            { text: 'Providers', link: '/platform/providers' },
            { text: 'Environments', link: '/platform/environments' },
          ]
        }
      ],

      '/runtime/': [
        {
          text: 'Runtime',
          items: [
            { text: 'Overview', link: '/runtime/overview' },
          ]
        }
      ],

      '/sdk/': [
        {
          text: 'Plugin SDK',
          items: [
            { text: 'Overview', link: '/sdk/overview' },
          ]
        }
      ],

      '/plugins/': [
        {
          text: 'Plugins',
          items: [
            { text: 'Architecture', link: '/plugins/overview' },
            { text: 'Supported Plugins', link: '/plugins/supported' },
          ]
        }
      ],

      '/api/': [
        {
          text: 'API Reference',
          items: [
            { text: 'Contracts', link: '/api/contracts' },
            { text: 'Models', link: '/api/models' },
          ]
        }
      ],

      '/cli/': [
        {
          text: 'CLI',
          items: [
            { text: 'Reference', link: '/cli/reference' },
          ]
        }
      ],

      '/testing/': [
        {
          text: 'Testing',
          items: [
            { text: 'Overview', link: '/testing/overview' },
          ]
        }
      ],

      '/development/': [
        {
          text: 'Development',
          items: [
            { text: 'Developer Guide', link: '/development/guide' },
          ]
        }
      ],

      '/deployment/': [
        {
          text: 'Deployment',
          items: [
            { text: 'Guide', link: '/deployment/guide' },
          ]
        }
      ],

      '/adr/': [
        {
          text: 'Architecture Decision Records',
          items: [
            { text: 'ADR 0001 — Why Go', link: '/adr/0001-go' },
            { text: 'ADR 0002 — Pipeline Architecture', link: '/adr/0002-pipeline-architecture' },
            { text: 'ADR 0003 — Platform Abstraction', link: '/adr/0003-platform-abstraction' },
            { text: 'ADR 0004 — Discovery Cache', link: '/adr/0004-discovery-cache' },
          ]
        }
      ],

      '/ai/': [
        {
          text: 'AI Context',
          items: [
            { text: 'Master Context', link: '/ai/AI_CONTEXT' },
            { text: 'Repository State', link: '/ai/REPOSITORY_STATE' },
            { text: 'Engineering Rules', link: '/ai/ENGINEERING_RULES' },
            { text: 'Architecture Index', link: '/ai/ARCHITECTURE_INDEX' },
            { text: 'Agent Guide', link: '/ai/AGENT_GUIDE' },
            { text: 'Known Issues', link: '/ai/KNOWN_ISSUES' },
          ]
        }
      ],
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/Hardik-Sankhla/base-infrastructure' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2025 Hardik Sankhla'
    },

    editLink: {
      pattern: 'https://github.com/Hardik-Sankhla/base-infrastructure/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },

    search: {
      provider: 'local'
    },

    outline: {
      level: [2, 3]
    }
  },

  markdown: {
    theme: {
      light: 'github-light',
      dark: 'github-dark'
    }
  },

  srcDir: '../docs',
})
