VERSION=0.0.1
BINARY=Manhattan
MANINSTALLDIR=/usr/share/man
BASHINSTALLDIR=/etc/bash_completion.d
all: binary
binary:
	go build -ldflags "-X main.version=${VERSION}"
install: install-binary
	sudo install -m 644 manhattan.1 ${MANINSTALLDIR}/man1/
	sudo cp bash_autocomplete ${BASHINSTALLDIR}/manhattan
install-binary:
	install -d -m 0755 ${INSTALLDIR}
	sudo install -m 755 Manhattan ${INSTALLDIR}
clean:
	rm -f Manhattan
test:
	cd tests && go test
