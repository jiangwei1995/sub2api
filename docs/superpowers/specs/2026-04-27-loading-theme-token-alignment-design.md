# Loading Theme Token Alignment

**Date**: 2026-04-27  
**Status**: Approved

## Goal

将登录、注册、用户管理台、管理员管理台，以及其它依赖全局 `primary` 主题 token 的界面，统一到现有 loading 风格所使用的青绿色主题语言。

本次改动以“全局主题收敛”为目标，不重做页面结构，不重写业务组件，不引入新的设计系统。

## Scope

- 调整 `frontend/tailwind.config.js` 中的 `primary` 色阶定义
- 调整与 `primary` 绑定的全局渐变、光晕、mesh 背景和共享按钮/输入框视觉
- 保持现有浅色 / 深色结构不变，只统一主色表达
- 不改动 landing 页的琥珀主题
- 不主动重构业务页面，只让它们通过共享 token 自动继承新主题

## Non-Goals

- 不修改路由、业务逻辑、接口请求
- 不对所有硬编码的语义色（例如错误红、成功绿）做全量替换
- 不重新设计管理台的数据卡片结构或信息层级

## Implementation Plan

### 1. Theme Tokens

将 `primary` 色阶统一为更贴近 loading 风格的青绿色序列，并同步更新：

- `gradient-primary`
- `mesh-gradient`
- `glow` / `glow-lg` 阴影色

### 2. Shared Global Styles

在 `frontend/src/style.css` 中保持现有组件 API 不变，更新以下共享样式的视觉输出：

- `.btn-primary`
- `.input` focus ring / border
- `.text-gradient`
- 选中态、进度条、主色徽标

### 3. Page Inheritance

登录页、注册页、用户管理台、管理员管理台继续使用现有布局组件：

- `AuthLayout`
- `AppLayout`
- `AppHeader`
- `AppSidebar`

这些页面不做结构性调整，只通过全局 token 变化自动完成风格对齐。

## Verification

- 补测试锁定新的 `primary` token、背景渐变和 glow 配置
- 运行主题相关测试
- 运行 `pnpm typecheck`
- 运行 `pnpm build`

## Risks

- 任何使用 `primary-*` 的页面都会一起变色，这是有意的全局主题调整
- 少量使用硬编码蓝 / 紫 / 绿的局部卡片不会完全随 token 改变，但不影响本次“主色统一”目标
