VERSION=0.0.1
BINARY=Manhattan
MANINSTALLDIR=${PREFIX}/share/man
BASHINSTALLDIR=${PREFIX}/share/bash-completion/completions
all:
	binary
binary:
	go build -ldflags "-X main.version=${VERSION}"
install:
	install-binary
	install -m 644 manhattan.1 ${MANINSTALLDIR}/man1/
	install -m completion ${BASHINSTALLDIR}
install-binary:
	install -d -m 0755 ${INSTALLDIR}
	install -m 755 manhattan ${INSTALLDIR}
clean:
	rm -f Manhattan
test:
	cd tests && go test
