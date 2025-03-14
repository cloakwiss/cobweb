package tidy

/*
#cgo LDFLAGS: -L../ -ltidy
#include "headers/tidy.h"
#include "headers/tidybuffio.h"
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *tidy_html(char *buff, uint64_t buff_size) {
    // Tidy Shit
    // ---------------------------------------------------------------------------
    TidyBuffer output = {0};
    TidyBuffer errbuf = {0};
    int rc = -1;
    Bool ok;

    TidyDoc tdoc = tidyCreate(); // Initialize "document"

    ok = tidyOptSetBool(tdoc, TidyXhtmlOut, yes); // Convert to XHTML
                                                  //
    if (ok)
        rc = tidySetErrorBuffer(tdoc, &errbuf); // Capture diagnostics
                                                //
    if (rc >= 0)
        rc = tidyParseString(tdoc, buff); // Parse the input
                                          //
    if (rc >= 0)
        rc = tidyCleanAndRepair(tdoc); // Tidy it up!
                                       //
    if (rc >= 0)
        rc = tidyRunDiagnostics(tdoc); // Kvetch
                                       //
    if (rc > 1)                        // If error, force output.
        rc = (tidyOptSetBool(tdoc, TidyForceOutput, yes) ? rc : -1);

    if (rc >= 0)
        rc = tidySaveBuffer(tdoc, &output); // Pretty Print

    if (rc >= 0) {
        if (rc > 0)
            printf("\nDiagnostics:\n\n%s", errbuf.bp);
        // printf("\nAnd here is the result:\n\n%s", output.bp);

    } else {
        printf("A severe error (%d) occurred.\n", rc);
		return NULL;
	}

    char *output_buffer = calloc(output.size + 2, sizeof(char));
    if (output_buffer == NULL) {
        printf("error allocating output buffer.\n");
        tidyBufFree(&output);
        tidyBufFree(&errbuf);
        tidyRelease(tdoc);
        exit(0);
    }

    if (strncpy_s(output_buffer, output.size + 2, (const char *)output.bp,
                  output.size) != 0) {

        printf("error copying output\n");
        tidyBufFree(&output);
        tidyBufFree(&errbuf);
        tidyRelease(tdoc);
        exit(0);
    }

    tidyBufFree(&output);
    tidyBufFree(&errbuf);
    tidyRelease(tdoc);
    // ---------------------------------------------------------------------------

    return output_buffer;
}
*/
import "C"
import (
	// "fmt"
	// "io"
	// "os"
	"unsafe"
)

func TidyHTML(input []byte) []byte {
	cInput := C.CBytes(input)
	defer C.free(unsafe.Pointer(cInput))

	cOutput := C.tidy_html((*C.char)(cInput), C.uint64_t(len(input)))

	if cOutput == nil {
		return []byte("Error processing HTML")
	}
	defer C.free(unsafe.Pointer(cOutput))

	return C.GoBytes(unsafe.Pointer(cOutput), C.int(C.strlen(cOutput)))
}
