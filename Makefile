VERSION=0.0.1
BINARY=manhattan
MANINSTALLDIR=/usr/share/man
BASHINSTALLDIR=/etc/bash_completion.d
all: binary
binary:
	go build -ldflags "-X main.version=${VERSION}" -o manhattan
install: install-binary
	sudo install -m 644 manhattan.1 ${MANINSTALLDIR}/man1/
	sudo cp bash_autocomplete ${BASHINSTALLDIR}/manhattan
	sudo cp manhattan.json /etc/sysconfig/manhattan.json
install-binary:
	install -d -m 0755 ${INSTALLDIR}
	sudo install -m 755 manhattan ${INSTALLDIR}
clean:
	rm -f manhattan
test:
	go test oci-seccomp-gen/*
