// Package servicex provides service lifecycle orchestration with configuration-driven init, data, and save pipelines.
// Package servicex 提供服务生命周期编排，支持基于配置的初始化、数据加载与保存流水线。
//
// Created by xuzhuoxi on 2023-06-24.
// @author xuzhuoxi
//
// Capabilities:
//   - Service interfaces (IService, IInitService, IInitDataService, ILoadDataService, ISaveDataService, etc.) and ServiceBase embeddable helper
//   - ServiceInfo / ServiceConfig: register services by name with factory functions and ordered configuration
//   - ServiceManager: awake, init, init-data, and load-data phases with eventx notifications and optional logging
//   - Event name constants and payload types (ServiceStartData, ServiceResultData) for manager and service lifecycle
//
// 主要职责：
//   - 服务接口（IService、IInitService、IInitDataService、ILoadDataService、ISaveDataService 等）及可内嵌的 ServiceBase 辅助类型
//   - ServiceInfo / ServiceConfig：按名称与工厂函数注册服务并维护有序配置列表
//   - ServiceManager：按序执行激活、初始化、数据初始化与数据加载，并通过 eventx 事件通知，可选日志
//   - 管理器与服务生命周期事件名称常量及载荷类型（ServiceStartData、ServiceResultData）
package servicex
