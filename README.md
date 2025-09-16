# Cobweb

A lightweight and minimal application that saves webpages as EPUB documents, prioritizing good defaults. It uses a basic web scraper, so it may not work with pages that rely on JavaScript for dynamic content.

## What it does

- **Save webpages to EPUB**: Follows links up to a chosen depth and bundles pages and assets into a single `.epub`.
- **Simple by default**: Fetches HTML, normalizes it to XHTML, and generates basic `content.opf`, `nav.xhtml` (TOC), and `META-INF/container.xml`.
- **Two ways to use**:
  - **CLI**: `cobweb <url>` with flags for depth and excluding assets.
  - **Web UI**: run without args to open a small server for archiving.

## Status and limitations

- Focused on static HTML. **JavaScript‑heavy pages may not work**.
- EPUB 3 first. Metadata is minimal; TOC is derived from file paths.
- Asset filtering is filename/extension based; HTTP header filtering is not yet wired.

## Install

Prerequisites:

- Go 1.24+
- A `libtidy` build available for CGO (used to convert HTML → XHTML)

Build:

```bash
go build -o cobweb ./src
```

Run tests (optional):

```bash
go test ./src/...
```

## Usage

### CLI

```bash
./cobweb https://example.com/article \
  -d 1 \
  -O example-article \
  -j -c -i -f
```

- **positional**: the target URL
- **-O, --output**: output file name (without extension)
- **-d, --depth**: crawl depth (default 0)
- **-j, --no-js**: exclude JavaScript files
- **-c, --no-css**: exclude CSS files
- **-i, --no-images**: exclude images
- **-f, --no-fonts**: exclude fonts
- **-a, --no-audio**: exclude audio
- **-V, --no-video**: exclude video
- **-A, --allow-domain**: whitelist additional domains (repeatable)
- **-D, --block-domains**: blacklist domains (repeatable)
- **-C, --cookies**: cookie file (not yet used end‑to‑end)
- **-T, --timeout**: request timeout (default 60s)

Output: creates `<output>.epub` in the current directory (defaults to page title when `-O` is omitted).

### Web UI

Run without arguments:

```bash
./cobweb
```

- Serves on `http://localhost:8080/`
- UI lives in `src/public` and lets you submit a URL to archive
- Resultant `.epub` will be available to download from the server

## How it works (high‑level)

- `fetch`: Uses Colly to crawl the target URL and collect pages/assets.
- `tidy`: Calls a tiny CGO wrapper around `libtidy` to produce XHTML.
- `epub/process`: Normalizes paths, sorts assets, separates XHTML pages from other assets.
- `epub/manifests`: Emits `content.opf`, `nav.xhtml` (TOC), and `container.xml`.
- `epub/zip`: Writes everything into a single EPUB archive.

## Notes for development

- Module path: `github.com/cloakwiss/cobweb`
- Main entry: `src/main.go`
- CLI flags defined in `src/app/options.go`
- Web server in `src/web_ui`

### Building libtidy

You need a `libtidy` that your linker can find. On many systems:

- Linux: install `tidy`/`libtidy-dev` via your package manager
- macOS: `brew install tidy-html5`
- Windows: use a prebuilt `tidy` DLL or build from source and ensure the library is on PATH for CGO

The project links using `-ltidy`, see `src/tidy/tidy.go`.
