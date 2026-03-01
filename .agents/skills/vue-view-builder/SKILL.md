---
name: vue-view-builder
description: 在本仓库中新增或重构 Vue 3 页面（SFC + script setup + Element Plus）并同步路由配置。用户提出“新建页面、改造页面结构、补充交互、接入后台视图”时使用。
---

# Vue View Builder

## 执行目标
- 生成可直接运行的页面代码，而不是示例片段。
- 保持与现有目录和路由风格一致。
- 让页面在桌面端后台场景下具备可用的基础交互。

## 执行步骤
1. 先定位页面归属模块，优先放入 `src/views/<module>/`。
2. 使用单文件组件结构：`<script setup>`、`<template>`、`<style scoped lang="scss">`。
3. 使用 Composition API（`ref`/`reactive`/`computed`）组织状态与方法。
4. 使用 Element Plus 组件搭建页面框架，避免引入额外 UI 库。
5. 若是新页面，更新 `src/router/index.js` 对应路由与子路由。
6. 保持命名清晰：页面文件使用小驼峰或已有风格，变量与方法语义化。

## 代码约束
- 优先复用现有布局入口：`src/views/layout/layoutContainer.vue`。
- 需要接口数据时，不在页面内直接写 axios；通过 API 模块调用。
- 页面级样式仅处理当前页面布局和视觉；通用样式放到可复用位置。
- 不为“可能用到”预置复杂抽象；先满足当前页面需求。

## 完成检查
- 检查路由可达并可渲染。
- 检查页面在无数据和有数据时均不报错。
- 检查交互事件（按钮、分页、切换）具备基础反馈。
- 检查样式未污染其它页面（`scoped` 生效）。
