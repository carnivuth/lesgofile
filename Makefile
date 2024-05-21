prefix = /usr/local

all: build

build:
	mkdir build
	go build -o build/lesgofile

install: build
	cp build/lesgofile $(DESTDIR)$(prefix)/bin/lesgofile
	cp settings.example.json $(DESTDIR)$(prefix)/bin/settings.json

uninstall:
	mkdir -p $(DESTDIR)$(prefix)/bin
	-rm -f $(DESTDIR)$(prefix)/bin/lesgofile
	-rm -f $(DESTDIR)$(prefix)/bin/settings.json

clean:
	rm -fr build

distclean: clean

.PHONY: all install clean distclean uninstall
