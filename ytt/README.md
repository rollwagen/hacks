# YouTube Transcript Go (ytt)

A command-line tool written in Go to download and extract transcripts from YouTube videos.

## Installation

```bash
# Clone the repository
git clone https://github.com/rollwagen/hacks && cd youtube-transcript-go

# Go install
go install -ldflags="-s -w"  .
```

## Building

```bash
# Clone the repository
git clone https://github.com/rollwagen/hacks && cd youtube-transcript-go

# Build the project
make build
```

## Usage

```bash
ytt "https://www.youtube.com/watch?v=VIDEO_ID"
```
