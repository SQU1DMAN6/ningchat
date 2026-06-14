dev:
	./bin/ningchat local

production:
	./bin/ningchat production

build:
	go build -o ./bin/ningchat ./cmd
	chmod 755 ./bin/ningchat

reload:
	sudo systemctl daemon-reload

start:
	sudo systemctl start ningchat

enable:
	sudo systemctl enable ningchat

stop:
	sudo systemctl stop ningchat

restart:
	sudo sytemctl restart ningchat

status:
	sudo systemctl status ningchat

log:
	journalctl -u ningchat -f

pull:
	git pull

update: pull build reload restart status log
