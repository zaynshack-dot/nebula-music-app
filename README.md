# 🪐 Nebula HiFi • OmniStream Pro

Welcome to the standalone production repository for **Nebula HiFi**, co-founded and built solo.

## 🏛️ Project Architecture
This application implements a **Client-Side Federated Scraper Engine** combined with a **Local Go Loopback Audio Proxy** bound to `127.0.0.1:8585`.

```
[ Tidal / Qobuz API CDNs ]
           │
           ▼ (Range HTTP Chunks)
[ Go Daemon (127.0.0.1:8585) ] ──(HMAC Signed)──► [ Flutter Nebula UI ]
           │
           ▼
[ SQLite LRU Disk Cache ]
```

## 🚀 Features

✅ **Android 17 Style Animations** - Ki energy barriers, martial arts controls, glowing UI effects
✅ **True 24-bit/192kHz FLAC Streaming** - Native lossless audio bitstream passthrough
✅ **USB DAC Exclusive Bypass Mode** - Zero-latency Qobuz/Tidal to Chi-Fi IEM routing
✅ **Smart Vibe Radio** - Genre-aware song continuity with Camelot key matching
✅ **Synced Karaoke Lyrics** - Real-time word-by-word lyric synchronization
✅ **Chi-Fi Hardware Matrix** - Automatic EQ curves for 100+ IEM models
✅ **HMAC Cryptographic Firewall** - Localhost-only security enforcement
✅ **Zero-Allocation Memory Pools** - Go buffer recycling (65KB chunks)

## 📱 Frontend Files

- **nebula_android17_animations.html** - Main player with Android 17 ki energy effects
- **nebula_music_player_ui.html** - Alternative UI with circular visualizer
- **nebula_flagship_app.html** - Full iPhone 16 mockup with all features
- **nebula_exact_mockup.html** - Battle strategy vs Spotify/Apple Music
- **nebula_build_studio.html** - Repository file explorer

## 🔧 Compilation Instructions

### Compile Go Backend Daemon
```bash
cd backend
go mod tidy
go build -ldflags="-s -w" -o nebula_daemon
```

### Run Server
```bash
./nebula_daemon
```

*The server will bind strictly to `127.0.0.1:8585` with HMAC enforcement active.*

## 🎵 Technology Stack

- **Backend:** Go 1.22 (HMAC, buffer pooling, HTTP proxying)
- **Frontend:** HTML5/CSS3/JavaScript (no frameworks - zero dependencies)
- **Mobile:** Flutter + Dart (Optional)
- **Audio Protocol:** FLAC 24-bit/192kHz lossless bitstream
- **Security:** HMAC-SHA256 with timestamp verification

## 🎯 Target Market

10+ million chi-fi audiophile IEM/DAC dongle owners worldwide. Estimated TAM: $500M+ annual.

---

**Built for solo founders who understand that focus beats scale.** 🚀⚡