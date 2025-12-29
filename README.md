# 🛡️ Companion Sentinel: Public Intelligence Hub
> **Next-gen security command plane for proactive origin forensics and AI-native infrastructure hardening.**

[![Status: Operational](https://img.shields.io/badge/System-Operational-00ff00?style=for-the-badge&logo=statuspage)](https://the-companion.com)
[![Build: v1.0.0 Elite](https://img.shields.io/badge/Build-v1.0.0_Elite-blue?style=for-the-badge)](https://the-companion.com)
[![Architecture: Zero--Knowledge](https://img.shields.io/badge/Architecture-Zero--Knowledge-purple?style=for-the-badge)](https://the-companion.com/docs)

---

## 🚀 Overview
**Companion Sentinel** is an elite infrastructure security ecosystem designed for modern developers and enterprise operators. By combining lightweight OS-level agents with an advanced AI orchestration layer, we provide 360° forensic visibility and automated remediation for mission-critical Linux environments.

This repository serves as the **Public Operator Interface**. While the platform's core "Brain" remains private to ensure absolute security, we maintain this space for technical transparency, community intelligence gathering, and open-source forensic collaboration.

---

## 🧠 The Intelligence Layer
Our proprietary agent executes deep technical audits that go beyond standard application-level security:

### 1. WordPress Origin Forensics
*   **Zero-Trust Malware Discovery**: Proactively scans the `wp-content` hierarchy for high-risk PHP patterns (`eval`, `base64_decode`, `passthru`) hidden within security plugins or themes.
*   **Hidden Backdoor Detection**: Audits the filesystem for suspicious hidden files (`.log.php`, `.backdoor.php`) used for persistent web-shell access.
*   **Origin Hardening**: Automated verification of `wp-config.php` permissions, XML-RPC DDoS vectors, and sensitive `debug.log` exposure.

### 2. CIS Benchmark Auditing
*   **Industry Standard Hardening**: Automated verification of **Center for Internet Security (CIS)** Linux benchmarks.
*   **Asset Isolation**: Audits partition isolation for `/tmp` and enforces world-writable "Sticky Bit" standards to prevent unauthorized data manipulation.
*   **SSH Protocol Enforcement**: Strictly monitors for SSH Protocol 2 compliance and direct root-login restrictions.

### 3. File Integrity Monitoring (FIM)
*   **Cryptographic Auditing**: Calculates MD5/SHA-256 hashes for high-stakes system files (`/etc/shadow`, `/etc/ssh/sshd_config`) every 60 seconds.
*   **Tamper Detection**: Instantly alerts on unauthorized configuration changes or user-account manipulation.

### 4. Multimodal AI Interrogation
*   **Forensic Vision**: Paste terminal screenshots or log artifacts directly into the Mission Control center for instant visual analysis.
*   **Interrogative Chat**: A multi-turn AI protocol allowing operators to iteratively refine remediation scripts and investigate complex attack chains.

---

## 🔒 Security & Sovereignty
Built on a **Zero-Knowledge Mesh**, Companion Sentinel ensures your technical secrets are impenetrable:
*   **Encrypted Tunnels**: All telemetry is pushed via **TLS 1.3**; no inbound ports are required on your nodes.
*   **Secrets Isolation**: User-provided AI keys and financial PII are cryptographically locked at rest using **AES-256-GCM**.
*   **One-Way Identification**: Personal Access Tokens (PAT) and provisioning keys are stored exclusively as secure **SHA-256 hashes**.

---

## ⚡ Community Interaction
Use this repository's **Issue Tracker** to help refine the sentinel intelligence feed:

*   **[Log a Technical Issue](https://github.com/BrandenNaidoo/Companion-Sentinel-Public/issues/new)**: Found a bug in the agent or dashboard? Let us know.
*   **[Request Forensic Signature](https://github.com/BrandenNaidoo/Companion-Sentinel-Public/issues/new?title=[Signature+Request]+)**: Want the agent to detect a specific CVE, malware pattern, or framework vulnerability? Submit your request for engineer review.
*   **[Technical Wiki](https://the-companion.com/docs)**: Read the full Technical Manual.

---

## 🌎 Enterprise Marketplace
Ready to build a proactive security mesh?
**[Initiate Production Deployment →](https://the-companion.com)**

&copy; 2025 Companion Sentinel. All Rights Reserved.