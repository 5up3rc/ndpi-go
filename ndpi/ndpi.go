package ndpi

// XXX we should use something like
// pkg-config --libs libndpi

// #cgo pkg-config: libndpi
/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
//#include <libndpi-1.8.0/libndpi/ndpi_api.h>
#include <libndpi-1.8.0/libndpi/ndpi_main.h>

extern void *malloc_wrapper(unsigned long size);
extern void free_wrapper(void *freeable);



static inline struct ndpi_detection_module_struct* ndpi_init()
{

	set_ndpi_malloc(malloc_wrapper), set_ndpi_free(free_wrapper);
	//set_ndpi_flow_malloc(NULL), set_ndpi_flow_free(NULL);

	struct ndpi_detection_module_struct* my_ndpi_struct = ndpi_init_detection_module();

	if (my_ndpi_struct == NULL) {
		return NULL;
	}

	//    my_ndpi_struct->http_dont_dissect_response=1;

	NDPI_PROTOCOL_BITMASK all;

	NDPI_BITMASK_ADD(all,NDPI_PROTOCOL_HTTP);
	NDPI_BITMASK_ADD(all,NDPI_PROTOCOL_SSL);

	// enable all protocols
	//    NDPI_BITMASK_SET_ALL(all);
        ndpi_set_protocol_detection_bitmask2(my_ndpi_struct, &all);

	return my_ndpi_struct;
}


*/
import "C"

import (
	"errors"
	"log"
	"unsafe"
)

//export malloc_wrapper
func malloc_wrapper(size C.size_t) unsafe.Pointer {

	return unsafe.Pointer(C.malloc(size))
}

//export free_wrapper
func free_wrapper(freeable unsafe.Pointer) {

	C.free(freeable)

}

type NDPIWrapper struct {
	cM (*C.struct_ndpi_detection_module_struct)
}

var ErrInitFailed = errors.New("nDPI: init failed")

func (n *NDPIWrapper) Init() error {
	log.Println("Init nDPI")
	n.cM = C.ndpi_init()
	if n.cM == nil {
		return ErrInitFailed
	}
	return nil

}
