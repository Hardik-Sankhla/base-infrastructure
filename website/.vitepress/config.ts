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
      { text: 'Development', link: '/development/guide' },
      { text: 'Engineering', link: '/engineering/overview' },
      { text: 'Reference', link: '/cli/reference' },
      {
        text: 'Project',
        items: [
          { text: 'Roadmap', link: 'https://github.com/Hardik-Sankhla/base-infrastructure/blob/main/ROADMAP.md' },
          { text: 'Contributing', link: 'https://github.com/Hardik-Sankhla/base-infrastructure/blob/main/CONTRIBUTING.md' },
          { text: 'Release Notes', link: 'https://github.com/Hardik-Sankhla/base-infrastructure/blob/main/CHANGELOG.md' },
          { text: 'GitHub', link: 'https://github.com/Hardik-Sankhla/base-infrastructure' },
        ]
      }
    ],

    sidebar: {
      '/getting-started/': [
        {
          text: 'Getting Started',
          items: [
            { text: 'Introduction', link: '/introduction' },
            { text: 'Installation', link: '/getting-started/install' },
            { text: 'Quickstart', link: '/getting-started/quickstart' },
            { text: 'Usage Guide', link: '/getting-started/usage' },
          ]
        }
      ],

      '/architecture/': [
        {
          text: 'Architecture',
          items: [
            { text: 'Overview', link: '/architecture/overview' },
            { text: 'Platform Design', link: '/platform/abstraction' },
            { text: 'Discovery Pipeline', link: '/discovery/engine' },
            { text: 'Diagrams', link: '/architecture/diagrams' },
            { text: 'Decision Log', link: '/architecture/DECISION_LOG' },
          ]
        }
      ],

      '/development/': [
        {
          text: 'Development',
          items: [
            { text: 'Developer Guide', link: '/development/guide' },
            { text: 'Testing', link: '/testing/overview' },
            { text: 'Deployment', link: '/deployment/guide' },
          ]
        }
      ],

      '/engineering/': [
        {
          text: 'Engineering OS',
          items: [
            { text: 'Overview', link: '/engineering/overview' },
          ]
        }
      ],

      '/guides/': [
        {
          text: 'Guides',
          items: [
            { text: 'Engineering Rules', link: '/guides/ENGINEERING_RULES' },
            { text: 'Known Issues', link: '/guides/KNOWN_ISSUES' },
          ]
        }
      ],

      '/cli/': [
        {
          text: 'Reference',
          items: [
            { text: 'CLI Reference', link: '/cli/reference' },
            { text: 'Capabilities', link: '/concepts/capabilities' },
            { text: 'Discovery Manifest', link: '/concepts/manifest' },
            { text: 'Glossary', link: '/concepts/glossary' },
            { text: 'API Contracts', link: '/api/contracts' },
            { text: 'API Models', link: '/api/models' },
            { text: 'Plugin SDK', link: '/sdk/overview' },
            { text: 'Plugin Architecture', link: '/plugins/overview' },
            { text: 'Supported Plugins', link: '/plugins/supported' },
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
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/Hardik-Sankhla/base-infrastructure' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2026 Hardik Sankhla'
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
