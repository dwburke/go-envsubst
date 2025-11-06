# Release Process

This project uses GoReleaser with GitHub Actions for automated releases.

## Creating a Release

1. **Update version** (if needed in documentation)

2. **Create and push a tag:**

   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. **GitHub Actions will automatically:**
   - Build binaries for multiple platforms (Linux, macOS, Windows)
   - Create DEB, RPM, and APK packages
   - Generate checksums
   - Create a GitHub Release with all artifacts

## Supported Platforms

- **Linux**: amd64, arm64, armv7
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

## Package Formats

- **DEB** (Debian/Ubuntu)
- **RPM** (RHEL/Fedora/CentOS)
- **APK** (Alpine Linux)
- **Archives** (tar.gz for Linux/macOS, zip for Windows)

## Arch Linux (AUR)

After creating a release:

1. Clone/update the AUR repository
2. Update the `PKGBUILD` file:
   - Update `pkgver` to the new version
   - Download the `checksums.txt` from the GitHub release
   - Update the SHA256 checksums for each architecture

3. Test the build:

   ```bash
   makepkg -si
   ```

4. Push to AUR:

   ```bash
   git add PKGBUILD .SRCINFO
   git commit -m "Update to v1.0.0"
   git push
   ```

## Testing Locally

To test the release process without publishing:

```bash
goreleaser release --snapshot --clean
```

This creates builds in the `dist/` directory without pushing to GitHub.
