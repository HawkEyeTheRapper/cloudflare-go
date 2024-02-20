// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
)

type V4PagePaginationResult struct {
	Items []T                        `json:"items"`
	JSON  v4PagePaginationResultJSON `json:"-"`
}

// v4PagePaginationResultJSON contains the JSON metadata for the struct
// [V4PagePaginationResult]
type v4PagePaginationResultJSON struct {
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type V4PagePaginationResultInfo struct {
	Page    int64                          `json:"page"`
	PerPage int64                          `json:"per_page"`
	JSON    v4PagePaginationResultInfoJSON `json:"-"`
}

// v4PagePaginationResultInfoJSON contains the JSON metadata for the struct
// [V4PagePaginationResultInfo]
type v4PagePaginationResultInfoJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type V4PagePagination[T any] struct {
	Result     V4PagePaginationResult     `json:"result"`
	ResultInfo V4PagePaginationResultInfo `json:"result_info"`
	JSON       v4PagePaginationJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// v4PagePaginationJSON contains the JSON metadata for the struct
// [V4PagePagination[T]]
type v4PagePaginationJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *V4PagePagination[T]) GetNextPage() (res *V4PagePagination[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *V4PagePagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type V4PagePaginationAutoPager[T any] struct {
	page *V4PagePagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewV4PagePaginationAutoPager[T any](page *V4PagePagination[T], err error) *V4PagePaginationAutoPager[T] {
	return &V4PagePaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *V4PagePaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Result.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *V4PagePaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *V4PagePaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *V4PagePaginationAutoPager[T]) Index() int {
	return r.run
}

type V4PagePaginationArrayResultInfo struct {
	Page    int64                               `json:"page"`
	PerPage int64                               `json:"per_page"`
	JSON    v4PagePaginationArrayResultInfoJSON `json:"-"`
}

// v4PagePaginationArrayResultInfoJSON contains the JSON metadata for the struct
// [V4PagePaginationArrayResultInfo]
type v4PagePaginationArrayResultInfoJSON struct {
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationArrayResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type V4PagePaginationArray[T any] struct {
	Result     []T                             `json:"result"`
	ResultInfo V4PagePaginationArrayResultInfo `json:"result_info"`
	JSON       v4PagePaginationArrayJSON       `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// v4PagePaginationArrayJSON contains the JSON metadata for the struct
// [V4PagePaginationArray[T]]
type v4PagePaginationArrayJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V4PagePaginationArray[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *V4PagePaginationArray[T]) GetNextPage() (res *V4PagePaginationArray[T], err error) {
	u := r.cfg.Request.URL
	currentPage, err := strconv.Atoi(u.Query().Get("page"))
	if err != nil {
		currentPage = 1
	}
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *V4PagePaginationArray[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type V4PagePaginationArrayAutoPager[T any] struct {
	page *V4PagePaginationArray[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewV4PagePaginationArrayAutoPager[T any](page *V4PagePaginationArray[T], err error) *V4PagePaginationArrayAutoPager[T] {
	return &V4PagePaginationArrayAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *V4PagePaginationArrayAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Result) == 0 {
		return false
	}
	if r.idx >= len(r.page.Result) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Result) == 0 {
			return false
		}
	}
	r.cur = r.page.Result[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *V4PagePaginationArrayAutoPager[T]) Current() T {
	return r.cur
}

func (r *V4PagePaginationArrayAutoPager[T]) Err() error {
	return r.err
}

func (r *V4PagePaginationArrayAutoPager[T]) Index() int {
	return r.run
}