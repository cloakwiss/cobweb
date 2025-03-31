#include "tidy.h"
#include "tidybuffio.h"
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *tidy_html(char *buff, uint64_t buff_size) {
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
            // printf("\nDiagnostics:\n\n%s", errbuf.bp);
        // printf("\nAnd here is the result:\n\n%s", output.bp);
        ;

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

    strncpy(output_buffer, (const char *)output.bp, output.size);

    tidyBufFree(&output);
    tidyBufFree(&errbuf);
    tidyRelease(tdoc);
    // ---------------------------------------------------------------------------

    return output_buffer;
}
