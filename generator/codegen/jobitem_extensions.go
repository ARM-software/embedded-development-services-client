package codegen

import (
	"slices"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

type JobItem struct {
	JobItemSchema string
}

type JobItems = []JobItem

type JobItemsParams = struct {
	JobItems
}

func AddJobItemsToParams(d *Data) (err error) {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return GetJobItems(swagger) }, "jobs.go")
}

func GetJobItems(swagger *openapi3.T) (jobItems JobItemsParams, err error) {
	var jobs JobItems
	for schemaName, schema := range swagger.Components.Schemas {
		schemaVal := schema.Value
		if schemaVal == nil {
			err = commonerrors.Newf(commonerrors.ErrUndefined, "the schema response for '%s' exists but has no value", schemaName)
			return
		}

		isRedacted, subErr := isExtensionFlagSet(schemaVal.ExtensionProps, redactFlag)
		if subErr != nil {
			err = subErr
			return
		}
		if isRedacted {
			continue
		}

		isJobItem, subErr := isExtensionFlagSet(schemaVal.ExtensionProps, jobFlag)
		if subErr != nil {
			err = subErr
			return
		}
		if isJobItem {
			jobs = append(jobs, JobItem{schemaName})
		}
	}

	slices.SortFunc(jobs, func(a, b JobItem) int {
		return strings.Compare(a.JobItemSchema, b.JobItemSchema)
	})

	jobItems = JobItemsParams{
		jobs,
	}

	return
}
