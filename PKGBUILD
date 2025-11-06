# Maintainer: Your Name <your.email@example.com>
pkgname=go-envsubst-bin
pkgver=1.0.0
pkgrel=1
pkgdesc="Go-powered envsubst with full template logic"
arch=('x86_64' 'aarch64' 'armv7h')
url="https://github.com/dwburke/go-envsubst"
license=('MIT')
provides=('go-envsubst')
conflicts=('go-envsubst')
source_x86_64=("${url}/releases/download/v${pkgver}/go-envsubst_Linux_x86_64.tar.gz")
source_aarch64=("${url}/releases/download/v${pkgver}/go-envsubst_Linux_arm64.tar.gz")
source_armv7h=("${url}/releases/download/v${pkgver}/go-envsubst_Linux_armv7.tar.gz")

# Update these checksums after each release
# Download the checksums.txt from the GitHub release and copy the appropriate SHA256 values
sha256sums_x86_64=('SKIP')
sha256sums_aarch64=('SKIP')
sha256sums_armv7h=('SKIP')

package() {
    install -Dm755 go-envsubst "${pkgdir}/usr/bin/go-envsubst"
    install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
    install -Dm644 README.md "${pkgdir}/usr/share/doc/${pkgname}/README.md"
}
