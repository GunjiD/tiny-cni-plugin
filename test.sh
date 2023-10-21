set -e

mkdir -p /tmp/tiny-cni-plugin
(export XDG_RUNTIME_DIR=/tmp/tiny-cni-plugin; unshare -rmn go test)
