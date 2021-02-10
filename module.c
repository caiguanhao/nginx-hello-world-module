#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>
#include <main.h>

static ngx_http_module_t ngx_http_hello_world_module_ctx = {
	NULL,
	NULL,

	NULL,
	NULL,

	NULL,
	NULL,

	NULL,
	NULL
};

static ngx_int_t ngx_http_hello_world_handler(ngx_http_request_t *r)
{
	return Handler(r);
}

static char *ngx_http_hello_world(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
{
	ngx_http_core_loc_conf_t *clcf;


	clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
	clcf->handler = ngx_http_hello_world_handler;

	return NGX_CONF_OK;
}

static ngx_command_t ngx_http_hello_world_commands[] = {

	{ ngx_string("hello_world"),
		NGX_HTTP_LOC_CONF|NGX_CONF_NOARGS,

		ngx_http_hello_world,
		0,
		0,
		NULL},

	ngx_null_command
};

ngx_module_t ngx_http_hello_world_module = {
	NGX_MODULE_V1,
	&ngx_http_hello_world_module_ctx,
	ngx_http_hello_world_commands,
	NGX_HTTP_MODULE,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NGX_MODULE_V1_PADDING
};

extern ngx_module_t  ngx_http_hello_world_module;

ngx_module_t *ngx_modules[] = {
	&ngx_http_hello_world_module,
	NULL
};

char *ngx_module_names[] = {
	"ngx_http_hello_world_module",
	NULL
};

char *ngx_module_order[] = {
	NULL
};
