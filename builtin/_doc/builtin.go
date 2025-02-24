package builtin

// echo formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Echo(a ...any) (n int, err error)
func Blines(r io.Reader) BLineReader

// type returns the reflection [Type] that represents the dynamic type of i.
// If i is a nil interface value, type returns nil.
func Type(i any) Type
func NewRange(start, end, step int) *IntRange

// print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...any) (n int, err error)

// println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...any) (n int, err error)

// printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) (n int, err error)

// errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand.
// If there is more than one %w verb, the returned error will implement an
// Unwrap method returning a []error containing all the %w operands in the
// order they appear in the arguments.
// It is invalid to supply the %w verb with an operand that does not implement
// the error interface. The %w verb is otherwise a synonym for %v.
func Errorf(format string, a ...any) error

// fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...any) (n int, err error)

// fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, a ...any) (n int, err error)

// fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)

// sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...any) string

// sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(a ...any) string

// sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...any) string

// open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(name string) (*File, error)

// create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*File, error)
func Lines(r io.Reader) LineReader
func (io.Reader) Gop_Enum() LineIter

// String converts the floating-point number f to a string,
// according to the format fmt and precision prec. It rounds the
// result assuming that the original was obtained from a floating-point
// value of bitSize bits (32 for float32, 64 for float64).
//
// The format fmt is one of
// 'b' (-ddddp±ddd, a binary exponent),
// 'e' (-d.dddde±dd, a decimal exponent),
// 'E' (-d.ddddE±dd, a decimal exponent),
// 'f' (-ddd.dddd, no exponent),
// 'g' ('e' for large exponents, 'f' otherwise),
// 'G' ('E' for large exponents, 'f' otherwise),
// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
// For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
func (float64) String(fmt byte, prec, bitSize int) string

// String is equivalent to FormatInt(int64(i), 10).
func (int) String() string

// String returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func (int64) String(base int) string

// String returns the string representation of i in the given base,
// for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
// for digit values >= 10.
func (uint64) String(base int) string

// The len built-in function returns the length of v, according to its type:
//
//	Array: the number of elements in v.
//	Pointer to array: the number of elements in *v (even if v is nil).
//	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
//	String: the number of bytes in v.
//	Channel: the number of elements queued (unread) in the channel buffer;
//	         if v is nil, len(v) is zero.
//
// For some arguments, such as a string literal or a simple array expression, the
// result can be a constant. See the Go language specification's "Length and
// capacity" section for details.
func (string) Len() int

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func (string) Count(substr string) int

// Int is equivalent to ParseInt(s, 10, 0), converted to type int.
func (string) Int() (int, error)

// Int64 interprets a string s in the given base (0, 2 to 36) and
// bit size (0 to 64) and returns the corresponding value i.
//
// The string may begin with a leading sign: "+" or "-".
//
// If the base argument is 0, the true base is implied by the string's
// prefix following the sign (if present): 2 for "0b", 8 for "0" or "0o",
// 16 for "0x", and 10 otherwise. Also, for argument base 0 only,
// underscore characters are permitted as defined by the Go syntax for
// [integer literals].
//
// The bitSize argument specifies the integer type
// that the result must fit into. Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
// If bitSize is below 0 or above 64, an error is returned.
//
// The errors that Int64 returns have concrete type *NumError
// and include err.Num = s. If s is empty or contains invalid
// digits, err.Err = ErrSyntax and the returned value is 0;
// if the value corresponding to s cannot be represented by a
// signed integer of the given size, err.Err = ErrRange and the
// returned value is the maximum magnitude integer of the
// appropriate bitSize and sign.
//
// [integer literals]: https://go.dev/ref/spec#Integer_literals
func (string) Int64(base int, bitSize int) (i int64, err error)

// Uint64 is like ParseInt but for unsigned numbers.
//
// A sign prefix is not permitted.
func (string) Uint64(base int, bitSize int) (uint64, error)

// Float converts the string s to a floating-point number
// with the precision specified by bitSize: 32 for float32, or 64 for float64.
// When bitSize=32, the result still has type float64, but it will be
// convertible to float32 without changing its value.
//
// Float accepts decimal and hexadecimal floating-point numbers
// as defined by the Go syntax for [floating-point literals].
// If s is well-formed and near a valid floating-point number,
// Float returns the nearest floating-point number rounded
// using IEEE754 unbiased rounding.
// (Parsing a hexadecimal floating-point value only rounds when
// there are more bits in the hexadecimal representation than
// will fit in the mantissa.)
//
// The errors that Float returns have concrete type *NumError
// and include err.Num = s.
//
// If s is not syntactically well-formed, Float returns err.Err = ErrSyntax.
//
// If s is syntactically well-formed but is more than 1/2 ULP
// away from the largest floating point number of the given size,
// Float returns f = ±Inf, err.Err = ErrRange.
//
// Float recognizes the string "NaN", and the (possibly signed) strings "Inf" and "Infinity"
// as their respective special floating point values. It ignores case when matching.
//
// [floating-point literals]: https://go.dev/ref/spec#Floating-point_literals
func (string) Float(bitSize int) (float64, error)

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func (string) Index(substr string) int

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func (string) IndexAny(chars string) int

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func (string) IndexByte(c byte) int

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func (string) IndexRune(r rune) int

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func (string) LastIndex(substr string) int

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func (string) LastIndexAny(chars string) int

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func (string) LastIndexByte(c byte) int

// Contains reports whether substr is within s.
func (string) Contains(substr string) bool

// ContainsAny reports whether any Unicode code points in chars are within s.
func (string) ContainsAny(chars string) bool

// ContainsRune reports whether the Unicode code point r is within s.
func (string) ContainsRune(r rune) bool

// Compare returns an integer comparing two strings lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
//
// Compare is included only for symmetry with package bytes.
// It is usually clearer and always faster to use the built-in
// string comparison operators ==, <, >, and so on.
func (string) Compare(b string) int

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding, which is a more general
// form of case-insensitivity.
func (string) EqualFold(t string) bool

// HasPrefix reports whether the string s begins with prefix.
func (string) HasPrefix(prefix string) bool

// HasSuffix reports whether the string s ends with suffix.
func (string) HasSuffix(suffix string) bool

// Quote returns a double-quoted Go string literal representing s. The
// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// control characters and non-printable characters as defined by
// IsPrint.
func (string) Quote() string

// Unquote interprets s as a single-quoted, double-quoted,
// or backquoted Go string literal, returning the string value
// that s quotes.  (If s is single-quoted, it would be a Go
// character literal; Unquote returns the corresponding
// one-character string.)
func (string) Unquote() (string, error)

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func (string) ToTitle() string

// ToUpper returns s with all Unicode letters mapped to their upper case.
func (string) ToUpper() string

// ToLower returns s with all Unicode letters mapped to their lower case.
func (string) ToLower() string

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func (string) Fields() []string

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if the result of (len(s) * count)
// overflows.
func (string) Repeat(count int) string

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to [SplitN] with a count of -1.
//
// To split around the first instance of a separator, see Cut.
func (string) Split(sep string) []string

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to [SplitAfterN] with a count of -1.
func (string) SplitAfter(sep string) []string

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [Split].
//
// To split around the first instance of a separator, see Cut.
func (string) SplitN(sep string, n int) []string

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unsplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for SplitAfter.
func (string) SplitAfterN(sep string, n int) []string

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func (string) Replace(old, new string, n int) string

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func (string) ReplaceAll(old, new string) string

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func (string) Trim(cutset string) string

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func (string) TrimSpace() string

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use [TrimPrefix] instead.
func (string) TrimLeft(cutset string) string

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use [TrimSuffix] instead.
func (string) TrimRight(cutset string) string

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func (string) TrimPrefix(prefix string) string

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func (string) TrimSuffix(suffix string) string

// Capitalize returns a copy of the string str with the first letter mapped to
// its upper case.
func (string) Capitalize() string

// The len built-in function returns the length of v, according to its type:
//
//	Array: the number of elements in v.
//	Pointer to array: the number of elements in *v (even if v is nil).
//	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
//	String: the number of bytes in v.
//	Channel: the number of elements queued (unread) in the channel buffer;
//	         if v is nil, len(v) is zero.
//
// For some arguments, such as a string literal or a simple array expression, the
// result can be a constant. See the Go language specification's "Length and
// capacity" section for details.
func ([]string) Len() int

// The cap built-in function returns the capacity of v, according to its type:
//
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.
//
// For some arguments, such as a simple array expression, the result can be a
// constant. See the Go language specification's "Length and capacity" section for
// details.
func ([]string) Cap() int

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func ([]string) Join(sep string) string

// Capitalize capitalizes the first letter of each string in the slice.
func ([]string) Capitalize() []string

// ToTitle title-cases all strings in the slice.
func ([]string) ToTitle() []string

// ToUpper upper-cases all strings in the slice.
func ([]string) ToUpper() []string

// ToLower lower-cases all strings in the slice.
func ([]string) ToLower() []string

// Repeat repeats each string in the slice count times.
func ([]string) Repeat(count int) []string

// Replace replaces all occurrences of old in each string in the slice with new.
func ([]string) Replace(old, new string, n int) []string

// ReplaceAll replaces all occurrences of old in each string in the slice with new.
func ([]string) ReplaceAll(old, new string) []string

// Trim removes leading and trailing white space from each string in the slice.
func ([]string) Trim(cutset string) []string

// TrimSpace removes leading and trailing white space from each string in the slice.
func ([]string) TrimSpace() []string

// TrimLeft removes leading white space from each string in the slice.
func ([]string) TrimLeft(cutset string) []string

// TrimRight removes trailing white space from each string in the slice.
func ([]string) TrimRight(cutset string) []string

// TrimPrefix removes leading prefix from each string in the slice.
func ([]string) TrimPrefix(prefix string) []string

// TrimSuffix removes trailing suffix from each string in the slice.
func ([]string) TrimSuffix(suffix string) []string
