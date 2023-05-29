package tasks

import (
	"io"
	"log"
	"net/http"
	"strings"

	"tabelf/backend/gen/entschema/urlcrawl"
	"tabelf/backend/service/app"
)

// BaiDuUrlCrawl 百度url链接提交.
func BaiDuUrlCrawl(jobCtx JobContext, config app.Config) {
	ctx := jobCtx.Context

	crawls, err := app.EntClient.UrlCrawl.Query().Where(
		urlcrawl.HasComplete(false),
		urlcrawl.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	if len(crawls) == 0 || app.Basic.Env != "prod" {
		return
	}

	app.Log.UnionLogger.Info("百度url链接提交.")
	urls := make([]string, 0)
	for _, crawl := range crawls {
		urls = append(urls, crawl.URL)
	}

	api := "http://data.zz.baidu.com/urls?site=https://www.tabelf.com&token=a7mYgScJ9x4AYEA7"
	reqBody := strings.Join(urls, "\n")
	reqContentType := "text/plain"

	req, err := http.NewRequest("POST", api, strings.NewReader(reqBody))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", reqContentType)

	count := 0

	for i := 0; i < 2000; i++ {
		func() {
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			count++
			log.Printf("crawl resp body %+v", string(body))
		}()
	}
	if err = app.EntClient.UrlCrawl.Update().AddCount(count).Where(
		urlcrawl.HasComplete(false),
		urlcrawl.DeactivatedAtIsNil(),
	).Exec(ctx); err != nil {
		return
	}
}
