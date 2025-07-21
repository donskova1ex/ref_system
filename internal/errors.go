package internal

import "errors"

var (
	ErrRecordNoFound     = errors.New("no record found")
	ErrCodeGenerate      = errors.New("error generating code")
	ErrOwnerNotFound     = errors.New("owner not found")
	ErrOwnerUUIDRequired = errors.New("uuid is required")
)
