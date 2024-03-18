# trpc-go 博客系统

欢迎来到基于trpc-go实现的博客系统，这是一个使用Go语言构建的全栈项目，充分利用了Go的简洁性、高性能和并发处理能力。本系统采用了trpc-go框架和北极星进行分布式开发，展示了如何在现代微服务架构中实现高效的服务通信和管理。

## 🌟 特性

- **Go语言实现**: 全部服务使用Go语言编写，继承了Go的优势，如简洁的语法、出色的性能、原生的并发支持和跨平台部署能力。
- **分布式开发**: 利用trpc-go框架和北极星实现了高效的分布式服务开发。
- **腾讯开源技术**: 项目基于腾讯的开源技术栈，保证了代码的现代性和高性能。
- **多样化代码架构**: 项目中的每个微服务采用了不同的代码组织方式，包括：
    - `go_clean_template` ( Clean Architecture)
    - `kratos-layout` (DDD架构)
    - `food-app-server` (DDD架构)
    - `project-layout` (Go项目标准布局)
- **全面的服务支持**: 包含用户服务、互动服务、博客服务、三方对接服务、权限服务和网关服务。
- **前端技术**: 前端使用Vue 3构建，提供了现代化的用户界面和交互体验。
- **安全性**: 采用网关鉴权机制，确保系统的安全性和稳定性。
- **在线演示**: 项目已经部署并上线，可以通过以下链接访问：[http://www.baker-yuan.cn/](http://www.baker-yuan.cn/)

## 🛠️ 使用的框架和布局

- **trpc-go**: [GitHub - trpc-group/trpc-go](https://github.com/trpc-group/trpc-go)
- **北极星**: [GitHub - polarismesh/polaris](https://github.com/polarismesh/polaris)
- **go-clean-template**: [GitHub - evrone/go-clean-template](https://github.com/evrone/go-clean-template)
- **kratos-layout**: [GitHub - go-kratos/kratos-layout](https://github.com/go-kratos/kratos-layout)
- **food-app-server**: [GitHub - victorsteven/food-app-server](https://github.com/victorsteven/food-app-server)
- **project-layout**: [GitHub - golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## 🚀 快速开始

请确保您已经安装了所有必要的依赖，包括Go环境、trpc-go框架和北极星服务。

1. 克隆仓库到本地：
   ```sh
   git clone https://github.com/baker-yuan/go_blog.git
   ```

2. 进入项目目录：
   ```sh
   cd go_blog
   ```

3. 安装依赖并启动各个微服务。

4. 启动前端项目。

## 💡 贡献

如果您对改进项目有任何建议或想要贡献代码，请随时提交Pull Request或创建Issue。

## ⚖️ 许可证

本项目采用 [MIT 许可证](LICENSE)。

---

感谢您对trpc-go博客系统的兴趣，我们期待您的参与和反馈，共同推动项目向前发展。