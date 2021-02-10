run: nginx-1.18.0/objs/nginx
	mkdir -p logs
	go build -v -buildmode c-shared -o main.so . && ./nginx-1.18.0/objs/nginx -p . -c nginx.conf

nginx-1.18.0/objs/nginx: nginx-1.18.0
	(cd nginx-1.18.0 && ./configure --with-compat --with-debug --with-cc-opt='-O0 -g' && make)

nginx-1.18.0: nginx-1.18.0.tar.gz
	tar xfvz nginx-1.18.0.tar.gz

nginx-1.18.0.tar.gz:
	wget https://nginx.org/download/nginx-1.18.0.tar.gz
