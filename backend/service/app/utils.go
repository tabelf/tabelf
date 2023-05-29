package app

import (
	"context"
	cryptorand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/status"
	"tabelf/backend/gen/entschema"
)

var ratio = decimal.NewFromInt(100)

// PriceFromString string-元转换为Int64-分.
func PriceFromString(price string) (int64, error) {
	p, err := decimal.NewFromString(price)
	if err != nil {
		return 0, err
	}
	return p.Mul(ratio).IntPart(), nil
}

func RPCError(err error, prefix string) error {
	if rpcErr, ok := status.FromError(err); ok {
		return errors.Errorf("%s: %s", prefix, rpcErr.Message())
	}
	return err
}

func HasString(list []string, value string) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}

// WithTx 事务快捷方式.
// https://entgo.io/docs/transactions/#best-practices
func WithTx(ctx context.Context, client *entschema.Client, fn func(tx *entschema.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			err = tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if r := tx.Rollback(); r != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", r)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func GetMilliTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func GetTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(DateWithTimeLayout)
}

// RandomString returns a random string with a fixed length.
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func IsNotBlank(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}

func IntSliceContains(source []int, target int) bool {
	for _, item := range source {
		if item == target {
			return true
		}
	}
	return false
}

func StringSliceContains(source []string, target string) bool {
	for _, item := range source {
		if item == target {
			return true
		}
	}
	return false
}

func ParseTime(timestamp string) time.Time {
	t, err := time.ParseInLocation(DateWithTimeLayout, timestamp, time.Local)
	if err != nil {
		return time.Now().UTC()
	}
	return t
}

func ParseDecimal(str interface{}) decimal.Decimal {
	var res decimal.Decimal
	var err error
	switch val := str.(type) {
	case int:
		res = decimal.NewFromInt(int64(val))
	case int8:
		res = decimal.NewFromInt(int64(val))
	case int16:
		res = decimal.NewFromInt(int64(val))
	case int32:
		res = decimal.NewFromInt(int64(val))
	case int64:
		res = decimal.NewFromInt(val)
	case float32:
		res = decimal.NewFromFloat32(val)
	case float64:
		res = decimal.NewFromFloat(val)
	case string:
		res, err = decimal.NewFromString(val)
	case decimal.Decimal:
		res = val
	default:
		res = decimal.Zero
	}
	if err != nil {
		return decimal.Zero
	}
	return res
}

func CompareDecimal(c1, c2 interface{}) int {
	d1 := ParseDecimal(c1)
	d2 := ParseDecimal(c2)
	return d1.Cmp(d2)
}

func AddDecimal(args ...interface{}) decimal.Decimal {
	res := decimal.Zero
	for _, d := range args {
		res = res.Add(ParseDecimal(d))
	}
	return res
}

func SubDecimal(c1, c2 interface{}) decimal.Decimal {
	d1 := ParseDecimal(c1)
	d2 := ParseDecimal(c2)
	return d1.Sub(d2)
}

func DivDecimal(c1, c2 interface{}) decimal.Decimal {
	d1 := ParseDecimal(c1)
	d2 := ParseDecimal(c2)
	return d1.Div(d2)
}

func MulDecimal(c1, c2 interface{}) decimal.Decimal {
	d1 := ParseDecimal(c1)
	d2 := ParseDecimal(c2)
	return d1.Mul(d2)
}

func GetYMD(time time.Time) string {
	return time.Format(YMDLayout)
}

func GetOrderNo() string {
	random := 0
	orderTime := fmt.Sprintf("%d", time.Now().UnixNano())[:17]
	n, err := cryptorand.Int(cryptorand.Reader, big.NewInt(10000000))
	if err != nil {
		random = int(time.Now().Unix())
	} else {
		random = int(n.Int64())
	}
	return orderTime + fmt.Sprintf("%07d", random)
}

func Then(flag bool, t, f any) any {
	if flag {
		return t
	}
	return f
}

func HtmlPlainText(html string) string {
	htmlRegex := regexp.MustCompile("<[^>]*>")
	return htmlRegex.ReplaceAllString(html, "")
}

func TruncateString(str string, maxLength int) string {
	length := utf8.RuneCountInString(str)
	if length <= maxLength {
		return str
	}
	runes := []rune(str)
	truncatedRunes := runes[:maxLength]
	return string(truncatedRunes)
}

/*
	HotArticleWeight 计算文章的权重.
	浏览量: 20
	使用量: 30
	收藏量: 20
	发布时间: 30
*/
var weights = map[string]float64{
	"views":      0.25,
	"used":       0.4,
	"collection": 0.25,
	"time":       0.1,
}

func HotArticleWeight(article *entschema.GoodArticle) float64 {
	now := time.Now()

	age := now.Sub(article.CreatedAt).Hours() / 24.0 // 计算文章年龄（天数）
	bv := math.Log(float64(article.View) + 1)        // 对浏览量取对数
	uv := math.Log(float64(article.Used) + 1)        // 对使用量取对数
	fv := math.Sqrt(float64(article.Star))           // 对收藏量开方
	wAge := math.Exp(-age / 7.0)                     // 根据文章年龄计算权重分数
	return weights["views"]*wAge*bv +
		weights["used"]*wAge*uv +
		weights["collection"]*fv +
		weights["time"]*wAge // 计算总得分
}
