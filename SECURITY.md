# Security Policy

## GeoVision3D

GeoVision3D takes security seriously and follows best practices for protecting API credentials, infrastructure, and user data.

---

# Supported Versions

| Version | Supported |
|---|---|
| Latest | ✅ |
| Older Releases | ❌ |

---

# Reporting a Vulnerability

If you discover a security vulnerability, please report it responsibly.

## Contact

**Project Maintainer:**  
Seli Kasela

📧 Email: shellybogel@gmail.com

---

# Recommended Report Information

Please include:

- vulnerability description
- affected component
- reproduction steps
- screenshots/logs (if applicable)
- potential impact

---

# Security Best Practices

GeoVision3D recommends:

- storing API keys in environment variables
- restricting Google Maps API domains
- enabling HTTPS in production
- avoiding hardcoded secrets
- using secure deployment environments
- limiting API permissions
- applying rate limiting for public endpoints

---

# API Key Security

Never expose:

- Google Maps API keys
- OpenCage API keys
- private credentials
- server environment files

Recommended protections:

- IP restrictions
- domain restrictions
- API quotas
- key rotation

---

# Infrastructure Security

Recommended deployment stack:

| Component | Recommendation |
|---|---|
| Reverse Proxy | NGINX |
| SSL | Cloudflare |
| Hosting | Railway / Render |
| Firewall | Cloudflare WAF |
| Environment Variables | Secure Secret Storage |

---

# Dependencies

Keep all dependencies updated regularly:

```bash
go mod tidy
go get -u ./...
```

---

# Production Recommendations

For production environments:

- enable HTTPS
- disable debug logs
- use secure CORS configuration
- monitor API usage
- apply request rate limiting

---

# Responsible Disclosure

Please do not publicly disclose security vulnerabilities until they have been reviewed and resolved.

---

# License

This project is released under the Apache License 2.0.
