# 🧠 GoMind
[![Go](https://github.com/osesantos/gomind/actions/workflows/go.yml/badge.svg)](https://github.com/osesantos/gomind/actions/workflows/go.yml)

> A lightweight, modular MCP (Multi-Component Protocol) server written in Go, focused on **Private RAG (Retrieval-Augmented Generation)** powered by local LLMs and your personal data.

---

## ✨ What is GoMind?

**GoMind** is a **brain orchestrator** that sits between a user interface (like LibreChat) and multiple intelligent agents or plugins. It:
- Receives natural language questions
- Orchestrates calls to various data sources (Obsidian notes, local vector DBs, etc.)
- Uses **Hermes** as a message bus to communicate with agents
- Assembles the full context and sends it to a local LLM (via Ollama)
- Returns the response to the user

Think of it as a **LangChain killer** — 100% private, no dependencies, fast and fully owned by you.

---

## 🧩 Architecture Overview

```text
[LibreChat or CLI]
        │
        ▼
     [ GoMind ]  ←←←←←←←←←←←←←←←←←←←←←←←←
        │                               │
        │                               │
        ├──→ Receives natural language question (e.g. "What is GoMind?")
        │                               │
        ├──→ Parses question and identifies required data sources
        │                               │
        │                               ├──→ [Obsidian Reader Agent] (MCP Core)
        │                               │
        │                               ├──→ [Vector Search Agent] (MCP Core)
        │                               │
        │                               └──→ [Other Agents]
     (MCP Core)                         │
        │                               │
        ├──→ Publishes requests via Hermes ("query.obsidian", "query.search", ...)
        │                               │
        ├──← Receives responses via Hermes (with correlation ID)
        │                               │
        └──→ Assembles context + sends prompt to LLM (Ollama)
        │
        ▼
    [ Local Response ]
```

---

## 🔧 Tech Stack

- **Go 1.22+**
- **Hermes** (lightweight pub/sub message bus)
- **Ollama** (local LLM runner, e.g. Mistral)
- **Meilisearch or Chroma** (for vector search)
- **Markdown file support** (e.g. Obsidian vaults)
- **JSON-based message protocol** with correlation ID

---

## 📦 Features

- 🧠 Private RAG from your local knowledge base
- 🪝 Plugin-based architecture using Hermes
- ⚡ Fast and lightweight (no LangChain or Python overhead)
- 🧰 Extensible with `hermes-go-sdk` for writing new agents
- ⏱️ Timeouts, fan-out/fan-in orchestration, modular pipeline

---

## 🚧 Status

> MVP in progress — building core functionality first:  
> ✅ Ollama connector  
> ✅ Obsidian reader agent  
> ✅ Basic message bus with Hermes  
> 🔜 Plugin protocol schema + timeout management  
> 🔜 Agent discovery + CLI mode

---

## 🚀 Getting Started

```bash
git clone https://github.com/osesantos/gomind
cd gomind
go run src/main.go
```

```bash
# Test with curl
curl -X POST http://localhost:4433/ask \
  -H "Content-Type: application/json" \
  -d '{"question": "What is GoMind?"}'

```

Make sure you have:
- Hermes running (or embedded mode)
- Ollama installed and serving a model (e.g. `ollama run mistral`)
- Some markdown notes to test against

---

## 📜 License

MIT — made with ❤️ and caffeine.

---

## 🙌 Credits

Inspired by:
- LangChain (but better)
- Personal RAG workflows
- Modular AI design

Built by [osesantos](https://github.com/osesantos)
