package result

type Result1[T1 any] struct {
	v1  T1
	err error
}

func (r *Result1[T1]) Unwrap() T1 {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1
}

func (r *Result1[T1]) Err() error {
	return r.err
}

func (r *Result1[T1]) IsErr() bool {
	return r.err != nil
}

func Ok1[T1 any](v1 T1) Result1[T1] {
	return Result1[T1]{v1: v1, err: nil}
}

func Err1[T1 any](err error) Result1[T1] {
	return Result1[T1]{err: err}
}

func (r *Result1[T1]) V1() T1 {
	return r.v1
}

type Result2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
}

func (r *Result2[T1, T2]) Unwrap() (T1, T2) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2
}

func (r *Result2[T1, T2]) Err() error {
	return r.err
}

func (r *Result2[T1, T2]) IsErr() bool {
	return r.err != nil
}

func Ok2[T1, T2 any](v1 T1, v2 T2) Result2[T1, T2] {
	return Result2[T1, T2]{v1: v1, v2: v2, err: nil}
}

func Err2[T1, T2 any](err error) Result2[T1, T2] {
	return Result2[T1, T2]{err: err}
}

func (r *Result2[T1, T2]) V1() T1 {
	return r.v1
}

func (r *Result2[T1, T2]) V2() T2 {
	return r.v2
}

type Result3[T1, T2, T3 any] struct {
	v1  T1
	v2  T2
	v3  T3
	err error
}

func (r *Result3[T1, T2, T3]) Unwrap() (T1, T2, T3) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3
}

func (r *Result3[T1, T2, T3]) Err() error {
	return r.err
}

func (r *Result3[T1, T2, T3]) IsErr() bool {
	return r.err != nil
}

func Ok3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Result3[T1, T2, T3] {
	return Result3[T1, T2, T3]{v1: v1, v2: v2, v3: v3, err: nil}
}

func Err3[T1, T2, T3 any](err error) Result3[T1, T2, T3] {
	return Result3[T1, T2, T3]{err: err}
}

func (r *Result3[T1, T2, T3]) V1() T1 {
	return r.v1
}

func (r *Result3[T1, T2, T3]) V2() T2 {
	return r.v2
}

func (r *Result3[T1, T2, T3]) V3() T3 {
	return r.v3
}

type Result4[T1, T2, T3, T4 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	err error
}

func (r *Result4[T1, T2, T3, T4]) Unwrap() (T1, T2, T3, T4) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4
}

func (r *Result4[T1, T2, T3, T4]) Err() error {
	return r.err
}

func (r *Result4[T1, T2, T3, T4]) IsErr() bool {
	return r.err != nil
}

func Ok4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Result4[T1, T2, T3, T4] {
	return Result4[T1, T2, T3, T4]{v1: v1, v2: v2, v3: v3, v4: v4, err: nil}
}

func Err4[T1, T2, T3, T4 any](err error) Result4[T1, T2, T3, T4] {
	return Result4[T1, T2, T3, T4]{err: err}
}

func (r *Result4[T1, T2, T3, T4]) V1() T1 {
	return r.v1
}

func (r *Result4[T1, T2, T3, T4]) V2() T2 {
	return r.v2
}

func (r *Result4[T1, T2, T3, T4]) V3() T3 {
	return r.v3
}

func (r *Result4[T1, T2, T3, T4]) V4() T4 {
	return r.v4
}

type Result5[T1, T2, T3, T4, T5 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	err error
}

func (r *Result5[T1, T2, T3, T4, T5]) Unwrap() (T1, T2, T3, T4, T5) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5
}

func (r *Result5[T1, T2, T3, T4, T5]) Err() error {
	return r.err
}

func (r *Result5[T1, T2, T3, T4, T5]) IsErr() bool {
	return r.err != nil
}

func Ok5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) Result5[T1, T2, T3, T4, T5] {
	return Result5[T1, T2, T3, T4, T5]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, err: nil}
}

func Err5[T1, T2, T3, T4, T5 any](err error) Result5[T1, T2, T3, T4, T5] {
	return Result5[T1, T2, T3, T4, T5]{err: err}
}

func (r *Result5[T1, T2, T3, T4, T5]) V1() T1 {
	return r.v1
}

func (r *Result5[T1, T2, T3, T4, T5]) V2() T2 {
	return r.v2
}

func (r *Result5[T1, T2, T3, T4, T5]) V3() T3 {
	return r.v3
}

func (r *Result5[T1, T2, T3, T4, T5]) V4() T4 {
	return r.v4
}

func (r *Result5[T1, T2, T3, T4, T5]) V5() T5 {
	return r.v5
}

type Result6[T1, T2, T3, T4, T5, T6 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	err error
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) Unwrap() (T1, T2, T3, T4, T5, T6) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) Err() error {
	return r.err
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) IsErr() bool {
	return r.err != nil
}

func Ok6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) Result6[T1, T2, T3, T4, T5, T6] {
	return Result6[T1, T2, T3, T4, T5, T6]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, err: nil}
}

func Err6[T1, T2, T3, T4, T5, T6 any](err error) Result6[T1, T2, T3, T4, T5, T6] {
	return Result6[T1, T2, T3, T4, T5, T6]{err: err}
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V1() T1 {
	return r.v1
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V2() T2 {
	return r.v2
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V3() T3 {
	return r.v3
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V4() T4 {
	return r.v4
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V5() T5 {
	return r.v5
}

func (r *Result6[T1, T2, T3, T4, T5, T6]) V6() T6 {
	return r.v6
}

type Result7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	err error
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Unwrap() (T1, T2, T3, T4, T5, T6, T7) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) Err() error {
	return r.err
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) IsErr() bool {
	return r.err != nil
}

func Ok7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) Result7[T1, T2, T3, T4, T5, T6, T7] {
	return Result7[T1, T2, T3, T4, T5, T6, T7]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, err: nil}
}

func Err7[T1, T2, T3, T4, T5, T6, T7 any](err error) Result7[T1, T2, T3, T4, T5, T6, T7] {
	return Result7[T1, T2, T3, T4, T5, T6, T7]{err: err}
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V1() T1 {
	return r.v1
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V2() T2 {
	return r.v2
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V3() T3 {
	return r.v3
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V4() T4 {
	return r.v4
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V5() T5 {
	return r.v5
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V6() T6 {
	return r.v6
}

func (r *Result7[T1, T2, T3, T4, T5, T6, T7]) V7() T7 {
	return r.v7
}

type Result8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	err error
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) Err() error {
	return r.err
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) IsErr() bool {
	return r.err != nil
}

func Ok8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) Result8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Result8[T1, T2, T3, T4, T5, T6, T7, T8]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, v8: v8, err: nil}
}

func Err8[T1, T2, T3, T4, T5, T6, T7, T8 any](err error) Result8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return Result8[T1, T2, T3, T4, T5, T6, T7, T8]{err: err}
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V1() T1 {
	return r.v1
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V2() T2 {
	return r.v2
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V3() T3 {
	return r.v3
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V4() T4 {
	return r.v4
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V5() T5 {
	return r.v5
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V6() T6 {
	return r.v6
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V7() T7 {
	return r.v7
}

func (r *Result8[T1, T2, T3, T4, T5, T6, T7, T8]) V8() T8 {
	return r.v8
}

type Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	v9  T9
	err error
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.v9
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Err() error {
	return r.err
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) IsErr() bool {
	return r.err != nil
}

func Ok9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, v8: v8, v9: v9, err: nil}
}

func Err9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](err error) Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{err: err}
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V1() T1 {
	return r.v1
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V2() T2 {
	return r.v2
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V3() T3 {
	return r.v3
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V4() T4 {
	return r.v4
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V5() T5 {
	return r.v5
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V6() T6 {
	return r.v6
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V7() T7 {
	return r.v7
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V8() T8 {
	return r.v8
}

func (r *Result9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) V9() T9 {
	return r.v9
}

type Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	v1  T1
	v2  T2
	v3  T3
	v4  T4
	v5  T5
	v6  T6
	v7  T7
	v8  T8
	v9  T9
	v10 T10
	err error
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Unwrap() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	if r.err != nil {
		panic(r.err)
	}
	return r.v1, r.v2, r.v3, r.v4, r.v5, r.v6, r.v7, r.v8, r.v9, r.v10
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Err() error {
	return r.err
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) IsErr() bool {
	return r.err != nil
}

func Ok10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{v1: v1, v2: v2, v3: v3, v4: v4, v5: v5, v6: v6, v7: v7, v8: v8, v9: v9, v10: v10, err: nil}
}

func Err10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](err error) Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{err: err}
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V1() T1 {
	return r.v1
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V2() T2 {
	return r.v2
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V3() T3 {
	return r.v3
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V4() T4 {
	return r.v4
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V5() T5 {
	return r.v5
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V6() T6 {
	return r.v6
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V7() T7 {
	return r.v7
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V8() T8 {
	return r.v8
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V9() T9 {
	return r.v9
}

func (r *Result10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) V10() T10 {
	return r.v10
}
