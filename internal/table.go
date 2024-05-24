package internal

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type TableWriter[T any] struct {
	columnHeaders []string
	columnChooser func(item T) []string
	writer        *tabwriter.Writer
	headerWritten bool
}

func NewTableWriter[T any](columnHeaders []string, columnChooser func(item T) []string) TableWriter[T] {
	return TableWriter[T]{
		columnHeaders: columnHeaders,
		columnChooser: columnChooser,
		writer:        tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0),
	}
}

func (tw *TableWriter[T]) WriteHeader() error {
	if !tw.headerWritten {
		_, err := fmt.Fprintln(tw.writer, strings.Join(tw.columnHeaders, "\t"))
		if err == nil {
			tw.headerWritten = true
		}
		return err
	}
	return nil
}

func (tw *TableWriter[T]) Write(item T) error {
	err := tw.WriteHeader()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(tw.writer, strings.Join(tw.columnChooser(item), "\t"))
	return err
}

func (tw *TableWriter[T]) Flush() error {
	return tw.writer.Flush()
}
