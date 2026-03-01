---
name: pinia-store-pattern
description: 按本仓库规范设计或重构 Pinia 状态管理，包括持久化和登录态管理。用户提出“新增全局状态、共享业务状态、维护 token/用户信息”时使用。
---

# Pinia Store Pattern

## 执行目标
- 让状态定义清晰、动作职责单一、持久化边界明确。
- 保持与 `src/stores/modules/user.js` 的风格一致。
- 避免把页面私有状态错误放入全局 store。

## 执行步骤
1. 在 `src/stores/modules/` 新建或更新对应 store 文件。
2. 使用 `defineStore('<id>', () => { ... }, { persist })` 定义状态。
3. 用 `ref`/`computed` 管理 state 与派生值，用具名 action 修改状态。
4. 在 `src/stores/index.js` 统一导出新增 store。
5. 在页面中按需引用 store，避免在多个位置重复业务逻辑。

## 状态设计规则
- 将跨页面共享且有复用价值的数据放入 store。
- 将只影响单页渲染的临时状态留在组件内部。
- 将认证相关字段（如 token、用户资料）集中在用户 store。
- 对持久化字段保持最小化，避免缓存大体量瞬时数据。

## 完成检查
- 检查 store id 唯一、命名语义清楚。
- 检查 action 不直接耦合具体页面 DOM 行为。
- 检查持久化配置符合数据生命周期预期。
- 检查页面刷新后关键状态（如登录态）行为正确。
