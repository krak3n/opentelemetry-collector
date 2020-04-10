// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "internal/data_generator/main.go". DO NOT EDIT.
// To regenerate this file run "go run internal/data_generator/main.go".

package data

import (
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
)

// ResourceSpansSlice logically represents a slice of ResourceSpans.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpansSlice struct {
	// orig points to the slice otlptrace.ResourceSpans field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.ResourceSpans
}

func newResourceSpansSlice(orig *[]*otlptrace.ResourceSpans) ResourceSpansSlice {
	return ResourceSpansSlice{orig}
}

// NewResourceSpansSlice creates a ResourceSpansSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewResourceSpansSlice() ResourceSpansSlice {
	orig := []*otlptrace.ResourceSpans(nil)
	return ResourceSpansSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceSpansSlice()".
func (es ResourceSpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es ResourceSpansSlice) At(ix int) ResourceSpans {
	return newResourceSpans(&(*es.orig)[ix])
}

// MoveTo moves all elements from the current slice to the dest. The current slice will be cleared.
func (es ResourceSpansSlice) MoveTo(dest ResourceSpansSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen < len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new ResourceSpansSlice can be initialized:
// es := NewResourceSpansSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es ResourceSpansSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*otlptrace.ResourceSpans(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen < oldLen {
		(*es.orig) = (*es.orig)[0:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]otlptrace.ResourceSpans, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpans struct {
	// orig points to the pointer otlptrace.ResourceSpans field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.ResourceSpans
}

func newResourceSpans(orig **otlptrace.ResourceSpans) ResourceSpans {
	return ResourceSpans{orig}
}

// NewResourceSpans creates a new "nil" ResourceSpans.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewResourceSpans() ResourceSpans {
	orig := (*otlptrace.ResourceSpans)(nil)
	return newResourceSpans(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms ResourceSpans) InitEmpty() {
	*ms.orig = &otlptrace.ResourceSpans{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms ResourceSpans) IsNil() bool {
	return *ms.orig == nil
}

// Resource returns the resource associated with this ResourceSpans.
// If no resource available, it creates an empty message and associates it with this ResourceSpans.
//
//  Empty initialized ResourceSpans will return "nil" Resource.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms ResourceSpans) Resource() Resource {
	return newResource(&(*ms.orig).Resource)
}

// InstrumentationLibrarySpans returns the InstrumentationLibrarySpans associated with this ResourceSpans.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms ResourceSpans) InstrumentationLibrarySpans() InstrumentationLibrarySpansSlice {
	return newInstrumentationLibrarySpansSlice(&(*ms.orig).InstrumentationLibrarySpans)
}

// InstrumentationLibrarySpansSlice logically represents a slice of InstrumentationLibrarySpans.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpansSlice struct {
	// orig points to the slice otlptrace.InstrumentationLibrarySpans field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.InstrumentationLibrarySpans
}

func newInstrumentationLibrarySpansSlice(orig *[]*otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpansSlice {
	return InstrumentationLibrarySpansSlice{orig}
}

// NewInstrumentationLibrarySpansSlice creates a InstrumentationLibrarySpansSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewInstrumentationLibrarySpansSlice() InstrumentationLibrarySpansSlice {
	orig := []*otlptrace.InstrumentationLibrarySpans(nil)
	return InstrumentationLibrarySpansSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewInstrumentationLibrarySpansSlice()".
func (es InstrumentationLibrarySpansSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es InstrumentationLibrarySpansSlice) At(ix int) InstrumentationLibrarySpans {
	return newInstrumentationLibrarySpans(&(*es.orig)[ix])
}

// MoveTo moves all elements from the current slice to the dest. The current slice will be cleared.
func (es InstrumentationLibrarySpansSlice) MoveTo(dest InstrumentationLibrarySpansSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen < len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new InstrumentationLibrarySpansSlice can be initialized:
// es := NewInstrumentationLibrarySpansSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es InstrumentationLibrarySpansSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*otlptrace.InstrumentationLibrarySpans(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen < oldLen {
		(*es.orig) = (*es.orig)[0:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]otlptrace.InstrumentationLibrarySpans, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// InstrumentationLibrarySpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewInstrumentationLibrarySpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type InstrumentationLibrarySpans struct {
	// orig points to the pointer otlptrace.InstrumentationLibrarySpans field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.InstrumentationLibrarySpans
}

func newInstrumentationLibrarySpans(orig **otlptrace.InstrumentationLibrarySpans) InstrumentationLibrarySpans {
	return InstrumentationLibrarySpans{orig}
}

// NewInstrumentationLibrarySpans creates a new "nil" InstrumentationLibrarySpans.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewInstrumentationLibrarySpans() InstrumentationLibrarySpans {
	orig := (*otlptrace.InstrumentationLibrarySpans)(nil)
	return newInstrumentationLibrarySpans(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms InstrumentationLibrarySpans) InitEmpty() {
	*ms.orig = &otlptrace.InstrumentationLibrarySpans{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms InstrumentationLibrarySpans) IsNil() bool {
	return *ms.orig == nil
}

// InstrumentationLibrary returns the instrumentationlibrary associated with this InstrumentationLibrarySpans.
// If no instrumentationlibrary available, it creates an empty message and associates it with this InstrumentationLibrarySpans.
//
//  Empty initialized InstrumentationLibrarySpans will return "nil" InstrumentationLibrary.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms InstrumentationLibrarySpans) InstrumentationLibrary() InstrumentationLibrary {
	return newInstrumentationLibrary(&(*ms.orig).InstrumentationLibrary)
}

// Spans returns the Spans associated with this InstrumentationLibrarySpans.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms InstrumentationLibrarySpans) Spans() SpanSlice {
	return newSpanSlice(&(*ms.orig).Spans)
}

// SpanSlice logically represents a slice of Span.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanSlice struct {
	// orig points to the slice otlptrace.Span field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span
}

func newSpanSlice(orig *[]*otlptrace.Span) SpanSlice {
	return SpanSlice{orig}
}

// NewSpanSlice creates a SpanSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanSlice() SpanSlice {
	orig := []*otlptrace.Span(nil)
	return SpanSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanSlice()".
func (es SpanSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanSlice) At(ix int) Span {
	return newSpan(&(*es.orig)[ix])
}

// MoveTo moves all elements from the current slice to the dest. The current slice will be cleared.
func (es SpanSlice) MoveTo(dest SpanSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen < len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new SpanSlice can be initialized:
// es := NewSpanSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*otlptrace.Span(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen < oldLen {
		(*es.orig) = (*es.orig)[0:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]otlptrace.Span, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// Span represents a single operation within a trace.
// See Span definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/master/opentelemetry/proto/trace/v1/trace.proto#L37
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpan function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Span struct {
	// orig points to the pointer otlptrace.Span field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.Span
}

func newSpan(orig **otlptrace.Span) Span {
	return Span{orig}
}

// NewSpan creates a new "nil" Span.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewSpan() Span {
	orig := (*otlptrace.Span)(nil)
	return newSpan(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms Span) InitEmpty() {
	*ms.orig = &otlptrace.Span{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms Span) IsNil() bool {
	return *ms.orig == nil
}

// TraceID returns the traceid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = []byte(v)
}

// SpanID returns the spanid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = []byte(v)
}

// TraceState returns the tracestate associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) TraceState() TraceState {
	return TraceState((*ms.orig).TraceState)
}

// SetTraceState replaces the tracestate associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetTraceState(v TraceState) {
	(*ms.orig).TraceState = string(v)
}

// ParentSpanID returns the parentspanid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) ParentSpanID() SpanID {
	return SpanID((*ms.orig).ParentSpanId)
}

// SetParentSpanID replaces the parentspanid associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetParentSpanID(v SpanID) {
	(*ms.orig).ParentSpanId = []byte(v)
}

// Name returns the name associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Name() string {
	return (*ms.orig).Name
}

// SetName replaces the name associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetName(v string) {
	(*ms.orig).Name = v
}

// Kind returns the kind associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Kind() SpanKind {
	return SpanKind((*ms.orig).Kind)
}

// SetKind replaces the kind associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetKind(v SpanKind) {
	(*ms.orig).Kind = otlptrace.Span_SpanKind(v)
}

// StartTime returns the starttime associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) StartTime() TimestampUnixNano {
	return TimestampUnixNano((*ms.orig).StartTimeUnixNano)
}

// SetStartTime replaces the starttime associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetStartTime(v TimestampUnixNano) {
	(*ms.orig).StartTimeUnixNano = uint64(v)
}

// EndTime returns the endtime associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) EndTime() TimestampUnixNano {
	return TimestampUnixNano((*ms.orig).EndTimeUnixNano)
}

// SetEndTime replaces the endtime associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetEndTime(v TimestampUnixNano) {
	(*ms.orig).EndTimeUnixNano = uint64(v)
}

// Attributes returns the Attributes associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// Events returns the Events associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Events() SpanEventSlice {
	return newSpanEventSlice(&(*ms.orig).Events)
}

// DroppedEventsCount returns the droppedeventscount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) DroppedEventsCount() uint32 {
	return (*ms.orig).DroppedEventsCount
}

// SetDroppedEventsCount replaces the droppedeventscount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetDroppedEventsCount(v uint32) {
	(*ms.orig).DroppedEventsCount = v
}

// Links returns the Links associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Links() SpanLinkSlice {
	return newSpanLinkSlice(&(*ms.orig).Links)
}

// DroppedLinksCount returns the droppedlinkscount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) DroppedLinksCount() uint32 {
	return (*ms.orig).DroppedLinksCount
}

// SetDroppedLinksCount replaces the droppedlinkscount associated with this Span.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) SetDroppedLinksCount(v uint32) {
	(*ms.orig).DroppedLinksCount = v
}

// Status returns the status associated with this Span.
// If no status available, it creates an empty message and associates it with this Span.
//
//  Empty initialized Span will return "nil" SpanStatus.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms Span) Status() SpanStatus {
	return newSpanStatus(&(*ms.orig).Status)
}

// SpanEventSlice logically represents a slice of SpanEvent.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEventSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEventSlice struct {
	// orig points to the slice otlptrace.Span_Event field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span_Event
}

func newSpanEventSlice(orig *[]*otlptrace.Span_Event) SpanEventSlice {
	return SpanEventSlice{orig}
}

// NewSpanEventSlice creates a SpanEventSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanEventSlice() SpanEventSlice {
	orig := []*otlptrace.Span_Event(nil)
	return SpanEventSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanEventSlice()".
func (es SpanEventSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanEventSlice) At(ix int) SpanEvent {
	return newSpanEvent(&(*es.orig)[ix])
}

// MoveTo moves all elements from the current slice to the dest. The current slice will be cleared.
func (es SpanEventSlice) MoveTo(dest SpanEventSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen < len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new SpanEventSlice can be initialized:
// es := NewSpanEventSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanEventSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*otlptrace.Span_Event(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen < oldLen {
		(*es.orig) = (*es.orig)[0:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]otlptrace.Span_Event, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// SpanEvent is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs. See OTLP for event definition.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEvent function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEvent struct {
	// orig points to the pointer otlptrace.Span_Event field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.Span_Event
}

func newSpanEvent(orig **otlptrace.Span_Event) SpanEvent {
	return SpanEvent{orig}
}

// NewSpanEvent creates a new "nil" SpanEvent.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewSpanEvent() SpanEvent {
	orig := (*otlptrace.Span_Event)(nil)
	return newSpanEvent(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms SpanEvent) InitEmpty() {
	*ms.orig = &otlptrace.Span_Event{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms SpanEvent) IsNil() bool {
	return *ms.orig == nil
}

// Timestamp returns the timestamp associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) Timestamp() TimestampUnixNano {
	return TimestampUnixNano((*ms.orig).TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) SetTimestamp(v TimestampUnixNano) {
	(*ms.orig).TimeUnixNano = uint64(v)
}

// Name returns the name associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) Name() string {
	return (*ms.orig).Name
}

// SetName replaces the name associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) SetName(v string) {
	(*ms.orig).Name = v
}

// Attributes returns the Attributes associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanEvent.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanEvent) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// SpanLinkSlice logically represents a slice of SpanLink.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLinkSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLinkSlice struct {
	// orig points to the slice otlptrace.Span_Link field contained somewhere else.
	// We use pointer-to-slice to be able to modify it in functions like Resize.
	orig *[]*otlptrace.Span_Link
}

func newSpanLinkSlice(orig *[]*otlptrace.Span_Link) SpanLinkSlice {
	return SpanLinkSlice{orig}
}

// NewSpanLinkSlice creates a SpanLinkSlice with 0 elements.
// Can use "Resize" to initialize with a given length.
func NewSpanLinkSlice() SpanLinkSlice {
	orig := []*otlptrace.Span_Link(nil)
	return SpanLinkSlice{&orig}
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewSpanLinkSlice()".
func (es SpanLinkSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     ... // Do something with the element
// }
func (es SpanLinkSlice) At(ix int) SpanLink {
	return newSpanLink(&(*es.orig)[ix])
}

// MoveTo moves all elements from the current slice to the dest. The current slice will be cleared.
func (es SpanLinkSlice) MoveTo(dest SpanLinkSlice) {
	if es.Len() == 0 {
		// Just to ensure that we always return a Slice with nil elements.
		*es.orig = nil
		return
	}
	if dest.Len() == 0 {
		*dest.orig = *es.orig
		*es.orig = nil
		return
	}
	*dest.orig = append(*dest.orig, *es.orig...)
	*es.orig = nil
	return
}

// Resize is an operation that resizes the slice:
// 1. If newLen is 0 then the slice is replaced with a nil slice.
// 2. If the newLen < len then equivalent with slice[0:newLen].
// 3. If the newLen > len then (newLen - len) empty elements will be appended to the slice.
//
// Here is how a new SpanLinkSlice can be initialized:
// es := NewSpanLinkSlice()
// es.Resize(4)
// for i := 0; i < es.Len(); i++ {
//     e := es.At(i)
//     // Here should set all the values for e.
// }
func (es SpanLinkSlice) Resize(newLen int) {
	if newLen == 0 {
		(*es.orig) = []*otlptrace.Span_Link(nil)
		return
	}
	oldLen := len(*es.orig)
	if newLen < oldLen {
		(*es.orig) = (*es.orig)[0:newLen]
		return
	}
	// TODO: Benchmark and optimize this logic.
	extraOrigs := make([]otlptrace.Span_Link, newLen-oldLen)
	oldOrig := (*es.orig)
	for i := range extraOrigs {
		oldOrig = append(oldOrig, &extraOrigs[i])
	}
	(*es.orig) = oldOrig
}

// SpanLink is a pointer from the current span to another span in the same trace or in a
// different trace. See OTLP for link definition.
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLink struct {
	// orig points to the pointer otlptrace.Span_Link field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.Span_Link
}

func newSpanLink(orig **otlptrace.Span_Link) SpanLink {
	return SpanLink{orig}
}

// NewSpanLink creates a new "nil" SpanLink.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewSpanLink() SpanLink {
	orig := (*otlptrace.Span_Link)(nil)
	return newSpanLink(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms SpanLink) InitEmpty() {
	*ms.orig = &otlptrace.Span_Link{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms SpanLink) IsNil() bool {
	return *ms.orig == nil
}

// TraceID returns the traceid associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) TraceID() TraceID {
	return TraceID((*ms.orig).TraceId)
}

// SetTraceID replaces the traceid associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) SetTraceID(v TraceID) {
	(*ms.orig).TraceId = []byte(v)
}

// SpanID returns the spanid associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) SpanID() SpanID {
	return SpanID((*ms.orig).SpanId)
}

// SetSpanID replaces the spanid associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) SetSpanID(v SpanID) {
	(*ms.orig).SpanId = []byte(v)
}

// TraceState returns the tracestate associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) TraceState() TraceState {
	return TraceState((*ms.orig).TraceState)
}

// SetTraceState replaces the tracestate associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) SetTraceState(v TraceState) {
	(*ms.orig).TraceState = string(v)
}

// Attributes returns the Attributes associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) Attributes() AttributeMap {
	return newAttributeMap(&(*ms.orig).Attributes)
}

// DroppedAttributesCount returns the droppedattributescount associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) DroppedAttributesCount() uint32 {
	return (*ms.orig).DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this SpanLink.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanLink) SetDroppedAttributesCount(v uint32) {
	(*ms.orig).DroppedAttributesCount = v
}

// SpanStatus is an optional final status for this span. Semantically when Status wasn't set
// it is means span ended without errors and assume Status.Ok (code = 0).
//
// This is a reference type, if passsed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanStatus function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanStatus struct {
	// orig points to the pointer otlptrace.Status field contained somewhere else.
	// We use pointer-to-pointer to be able to modify it in InitEmpty func.
	orig **otlptrace.Status
}

func newSpanStatus(orig **otlptrace.Status) SpanStatus {
	return SpanStatus{orig}
}

// NewSpanStatus creates a new "nil" SpanStatus.
// To initialize the struct call "InitEmpty".
//
// This must be used only in testing code since no "Set" method available.
func NewSpanStatus() SpanStatus {
	orig := (*otlptrace.Status)(nil)
	return newSpanStatus(&orig)
}

// InitEmpty overwrites the current value with empty.
func (ms SpanStatus) InitEmpty() {
	*ms.orig = &otlptrace.Status{}
}

// IsNil returns true if the underlying data are nil.
//
// Important: All other functions will cause a runtime error if this returns "true".
func (ms SpanStatus) IsNil() bool {
	return *ms.orig == nil
}

// Code returns the code associated with this SpanStatus.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanStatus) Code() StatusCode {
	return StatusCode((*ms.orig).Code)
}

// SetCode replaces the code associated with this SpanStatus.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanStatus) SetCode(v StatusCode) {
	(*ms.orig).Code = otlptrace.Status_StatusCode(v)
}

// Message returns the message associated with this SpanStatus.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanStatus) Message() string {
	return (*ms.orig).Message
}

// SetMessage replaces the message associated with this SpanStatus.
//
// Important: This causes a runtime error if IsNil() returns "true".
func (ms SpanStatus) SetMessage(v string) {
	(*ms.orig).Message = v
}
