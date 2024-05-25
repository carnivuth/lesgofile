prefix = /usr/local
unit_path = /etc/systemd/system

all: build

build: clean
	mkdir build
	go build -o build/lesgofile

install: build
	cp build/lesgofile $(DESTDIR)$(prefix)/bin/lesgofile
	cp settings.example.json $(DESTDIR)$(prefix)/bin/settings.json
	cp lesgofile.service $(unit_path)/lesgofile.service
	sed -i "s|\[\[BIN_PATH\]\]|$(DESTDIR)$(prefix)|g" "/etc/systemd/system/lesgofile.service"


uninstall:
	mkdir -p $(DESTDIR)$(prefix)/bin
	rm -f $(DESTDIR)$(prefix)/bin/lesgofile
	rm -f $(DESTDIR)$(prefix)/bin/settings.json
	rm -f $(unit_path)/lesgofile.service

clean:
	rm -fr build

distclean: clean

.PHONY: all install clean distclean uninstall
