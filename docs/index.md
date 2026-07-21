---
layout: home

hero:
  name: "Base Infrastructure"
  text: "Capability-Aware Bootstrap Framework"
  tagline: A single static binary that interrogates any host environment and maps it to a structured capability graph — Linux, Windows, macOS, Android, and BSD.
  image:
    src: /logo.svg
    alt: Base Infrastructure
  actions:
    - theme: brand
      text: Get Started
      link: /getting-started/install
    - theme: alt
      text: Architecture
      link: /architecture/overview
    - theme: alt
      text: View on GitHub
      link: https://github.com/Hardik-Sankhla/base-infrastructure

features:
  - icon: 🔍
    title: Cross-Platform Discovery
    details: Discovers OS, hardware, network, software, filesystem, and environment on Linux, Windows, macOS, Android (Termux), and BSD — without requiring Python, Ruby, or any other runtime.

  - icon: 🧩
    title: Capability Graph
    details: Every discovery result is translated into a structured capability (e.g., network.connectivity, runtime.docker). Capabilities become the input to the planner — not raw system data.

  - icon: 🔌
    title: Plugin Ecosystem
    details: Bundled plugins for Git, Docker, Python, and Pocketbase. Each plugin is a set of detect/install scripts executed via a STDIN/STDOUT JSON-RPC boundary — no Go required to write one.

  - icon: 🏗️
    title: Clean Architecture
    details: Full platform abstraction layer isolates all OS-specific syscalls. The discovery engine never knows what OS it is running on — it only talks to Provider interfaces.

  - icon: 🤖
    title: AI-Native Documentation
    details: The docs/ai/ directory contains machine-readable context files designed for AI coding agents — architecture index, engineering rules, known issues, and certification guides.

  - icon: 🚀
    title: Zero Dependencies at Runtime
    details: Compiles to a single static binary. No Go, no pip, no npm required on the target host. Cross-compile for Linux, Windows, and macOS from a single CI pipeline.
---
