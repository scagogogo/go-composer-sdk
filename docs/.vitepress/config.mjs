import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Go Composer SDK',
  description: 'A comprehensive Go library for PHP Composer package manager',

  // GitHub Pages configuration
  base: '/go-composer-sdk/',

  // Multi-language support
  locales: {
    root: {
      label: 'English',
      lang: 'en',
      title: 'Go Composer SDK',
      description: 'A comprehensive Go library for PHP Composer package manager',
      themeConfig: {
        nav: [
          { text: 'Home', link: '/' },
          { text: 'Guide', link: '/guide/getting-started' },
          { text: 'API Reference', link: '/api/' },
          { text: 'Examples', link: '/examples/' },
          { text: 'GitHub', link: 'https://github.com/scagogogo/go-composer-sdk' }
        ],
        sidebar: {
          '/guide/': [
            {
              text: 'Guide',
              items: [
                { text: 'Getting Started', link: '/guide/getting-started' },
                { text: 'Installation', link: '/guide/installation' },
                { text: 'Configuration', link: '/guide/configuration' },
                { text: 'Basic Usage', link: '/guide/basic-usage' }
              ]
            }
          ],
          '/api/': [
            {
              text: 'API Reference',
              items: [
                { text: 'Overview', link: '/api/' },
                { text: 'Core', link: '/api/core' },
                { text: 'Package Management', link: '/api/package-management' },
                { text: 'Project Management', link: '/api/project-management' },
                { text: 'Security & Audit', link: '/api/security-audit' },
                { text: 'Platform & Environment', link: '/api/platform-environment' },
                { text: 'Utilities', link: '/api/utilities' },
                { text: 'Detector', link: '/api/detector' },
                { text: 'Installer', link: '/api/installer' }
              ]
            }
          ],
          '/examples/': [
            {
              text: 'Examples',
              items: [
                { text: 'Overview', link: '/examples/' },
                { text: 'Basic Operations', link: '/examples/basic-operations' },
                { text: 'Package Management', link: '/examples/package-management' },
                { text: 'Project Setup', link: '/examples/project-setup' },
                { text: 'Security Audit', link: '/examples/security-audit' },
                { text: 'Advanced Usage', link: '/examples/advanced-usage' }
              ]
            }
          ]
        },
        socialLinks: [
          { icon: 'github', link: 'https://github.com/scagogogo/go-composer-sdk' }
        ],
        footer: {
          message: 'Released under the MIT License.',
          copyright: 'Copyright © 2024 Go Composer SDK'
        }
      }
    },
    zh: {
      label: '简体中文',
      lang: 'zh-CN',
      title: 'Go Composer SDK',
      description: '全面的 PHP Composer 包管理器 Go 语言库',
      themeConfig: {
        nav: [
          { text: '首页', link: '/zh/' },
          { text: '指南', link: '/zh/guide/getting-started' },
          { text: 'API 参考', link: '/zh/api/' },
          { text: '示例', link: '/zh/examples/' },
          { text: 'GitHub', link: 'https://github.com/scagogogo/go-composer-sdk' }
        ],
        sidebar: {
          '/zh/guide/': [
            {
              text: '指南',
              items: [
                { text: '快速开始', link: '/zh/guide/getting-started' },
                { text: '安装', link: '/zh/guide/installation' },
                { text: '配置', link: '/zh/guide/configuration' },
                { text: '基本用法', link: '/zh/guide/basic-usage' }
              ]
            }
          ],
          '/zh/api/': [
            {
              text: 'API 参考',
              items: [
                { text: '概览', link: '/zh/api/' },
                { text: '核心功能', link: '/zh/api/core' },
                { text: '包管理', link: '/zh/api/package-management' },
                { text: '项目管理', link: '/zh/api/project-management' },
                { text: '安全审计', link: '/zh/api/security-audit' },
                { text: '平台环境', link: '/zh/api/platform-environment' },
                { text: '工具函数', link: '/zh/api/utilities' },
                { text: '检测器', link: '/zh/api/detector' },
                { text: '安装器', link: '/zh/api/installer' }
              ]
            }
          ],
          '/zh/examples/': [
            {
              text: '示例',
              items: [
                { text: '概览', link: '/zh/examples/' },
                { text: '基本操作', link: '/zh/examples/basic-operations' },
                { text: '包管理', link: '/zh/examples/package-management' },
                { text: '项目设置', link: '/zh/examples/project-setup' },
                { text: '安全审计', link: '/zh/examples/security-audit' },
                { text: '高级用法', link: '/zh/examples/advanced-usage' }
              ]
            }
          ]
        },
        socialLinks: [
          { icon: 'github', link: 'https://github.com/scagogogo/go-composer-sdk' }
        ],
        footer: {
          message: '基于 MIT 许可证发布。',
          copyright: 'Copyright © 2024 Go Composer SDK'
        }
      }
    }
  },

  themeConfig: {
    search: {
      provider: 'local'
    }
  },

  // Ignore dead links for now during development
  ignoreDeadLinks: true
})
