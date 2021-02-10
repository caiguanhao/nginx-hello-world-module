package main

/*
#cgo CFLAGS: -I./nginx-1.18.0/src/core -I./nginx-1.18.0/objs -I./nginx-1.18.0/src/os/unix -I./nginx-1.18.0/src/http -I./nginx-1.18.0/src/event -I./nginx-1.18.0/src/http/modules -I./nginx-1.18.0/src/event/modules
#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
#include <ngx_core.h>
#include <ngx_http.h>

static ngx_chain_t chain_from_string(ngx_http_request_t *r, u_char *string, size_t len) {
	ngx_buf_t *b;
	b = ngx_pcalloc(r->pool, sizeof(ngx_buf_t));
	b->pos = string;
	b->last = string + len;
	b->memory = 1;
	b->last_buf = 1;

	ngx_chain_t out;
	out.buf = b;
	out.next = NULL;
	return out;
}
*/
import "C"

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
	"unsafe"
)

// example
func GetDate() string {
	out, _ := exec.Command("date").Output()
	return string(out)
}

func stringToUcharLen(input string) (*C.uchar, C.ulong, func()) {
	bytes := C.CBytes([]byte(input))
	return (*C.uchar)(bytes), C.ulong(len(input)), func() { C.free(bytes) }
}

//export Handler
func Handler(r *C.ngx_http_request_t) C.ngx_int_t {
	uri := C.GoStringN((*C.char)(unsafe.Pointer(r.uri.data)), C.int(r.uri.len))
	fmt.Println(uri)

	var respBody *C.uchar
	var respBodyLen C.ulong
	var free func()
	if uri == "/date" {
		respBody, respBodyLen, free = stringToUcharLen(GetDate())
	} else {
		rf := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
		respBody, respBodyLen, free = stringToUcharLen(fmt.Sprintf("%f", rf))
	}
	defer free()
	out := C.chain_from_string(r, respBody, respBodyLen)

	contentType, contentTypeLen, free := stringToUcharLen("text/plain")
	defer free()

	r.headers_out.content_type.data = contentType
	r.headers_out.content_type.len = contentTypeLen
	r.headers_out.status = 200
	r.headers_out.content_length_n = C.longlong(respBodyLen)
	C.ngx_http_send_header(r)
	return C.ngx_http_output_filter(r, &out)
}

func main() {
}
