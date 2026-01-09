# Zenit (巔峰) Framework
A high-performance full-stack framework designed specifically for enterprise applications. The backend is based on the industry's most stable Gin (Go), and the frontend is deeply integrated with SolidStart.

Zenit 是一個專為企業級應用設計的高性能全棧 (Full-stack) 框架。後端基於業界最穩定的 Gin (Go)，前端深度整合 SolidStart。

Zenit 的核心理念在於結合 Go 標準庫的「穩定禪意」與 Solid.js 的「細粒度效能」。它不僅擁有類似 Laravel 的開發架構，更針對 2026 年的 AI 輔助編碼 (Vibe Coding) 與雲原生 Kubernetes 環境進行了深度優化。

---

## ✨ 全棧核心特色
### 🛡️ 企業級核心 (Gin):
* **後端:** 基於 Gin Web Framework，完全相容 Go 標準庫 net/http，原生支援 HTTP/2 與 HTTP/3。
* **穩定性:** 擁有 Go 社群最強大的中間件生態與安全修補支援，適合長期維護的商業專案。

### ⚡️ 前端巔峰 (SolidStart):
整合 SolidStart，利用細粒度響應式更新 (Fine-grained reactivity)，徹底告別 Virtual DOM 的效能開銷。
🏗️ 穩固架構 (Laravel-inspired): 採用標準化的 Service-Repository 模式，內建依賴注入 (DI) 機制，確保複雜業務邏輯清晰、易測試。

### 🤖 AI 編碼友好 (Vibe Coding):
標準化結構: 嚴格的層次劃分讓 Cursor, GitHub Copilot 等 AI 工具能極速理解程式碼上下文。
強型別同步: 自動將後端 Go 結構體同步為前端 TypeScript 類型，實現端到端類型安全。

### ☁️ 雲原生 K8s Ready: 
內建支援 K8s 健康檢查、Prometheus 監控指標、OpenTelemetry 鏈路追蹤以及 Graceful Shutdown 機制。

---

## 🚀 快速開始 (2026 旗艦版)
#### 環境要求
* Go 1.24+
* Node.js 22+ (推薦使用 pnpm)

#### 初始化專案
```bash
# 使用 Zenit CLI 建立全棧專案
zenit new my-enterprise-app --fullstack --core=gin

cd my-enterprise-app
```

開發模式
Zenit 將同步啟動 Gin API 伺服器與 SolidStart 開發環境：
```bash
zenit dev
```

## 📂 專案結構 (Gin-based)
Zenit 採用現代化的全棧一體化結構，專為多人協作設計：
* `cmd/`: 後端啟動入口。
* `internal/`:
* `api/`: Gin 路由與 Handler 定義。
* `service/`: 核心業務邏輯層。
* `repository/`: 資料庫持久層 (支援 GORM / Ent)。
* `src/`: SolidStart 前端源碼 (Routes, Components, Signals)。
* `platform/`: Kubernetes Helm Charts、Terraform 與 Dockerfile 定義。

---

## ☁️ 為什麼在 2026 年選擇 Gin 核心？
雖然 Fiber 在基準測試中極快，但 Zenit (巔峰) 選擇 Gin 作為核心，是因為企業級後端更看重：
* **標準相容性:** 與雲端負載平衡器、WAF 及服務網格 (Service Mesh) 的無縫對接。
* **安全性:** 基於 net/http 的 Gin 擁有更成熟的記憶體安全管理與更嚴格的 RFC 規範實作。
* **人才生態:** 全球有超過 48% 的 Go 開發者熟悉 Gin，大幅降低企業的招聘與上手成本。

## 🛠️ 技術棧
* **Backend:** Go (Gin)
* **Frontend:** SolidStart (Solid.js)
* **Database:** PostgreSQL / MySQL
* **DevOps:** Docker, Kubernetes, GitHub Actions

---

## 🤝 參與貢獻
Zenit 代表著技術的巔峰。我們歡迎開發者加入，共同打造 2026 年最具生產力的 Go 全棧框架。

## 📄 授權
MIT License
