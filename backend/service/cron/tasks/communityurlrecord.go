package tasks

import (
	"fmt"
	"log"
	"tabelf/backend/gen/entschema"
	entcommunity "tabelf/backend/gen/entschema/community"
	"tabelf/backend/gen/entschema/urlcrawl"
	"tabelf/backend/service/app"
	"time"
)

// CommunityUrlRecord 收录社区url链接供爬虫爬取.
func CommunityUrlRecord(jobCtx JobContext, config app.Config) {
	current := time.Now()
	ctx := jobCtx.Context

	log.Print("community url crawls start.")
	// 收录30天以内的
	communities, err := app.EntClient.Community.Query().Where(
		entcommunity.Status(app.PassStatus),
		entcommunity.DeactivatedAtIsNil(),
		entcommunity.CreatedAtGTE(current.Add(-45*24*time.Hour)),
		entcommunity.CreatedAtLTE(current),
	).All(ctx)
	if err != nil {
		return
	}
	repeatURL := make([]string, 0)
	communityMap := make(map[string]*entschema.Community)
	for _, community := range communities {
		repeatURL = append(repeatURL, community.UID)
		communityMap[community.UID] = community
	}
	crawls, err := app.EntClient.UrlCrawl.Query().Where(
		urlcrawl.CommunityUIDIn(repeatURL...),
		urlcrawl.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	for _, crawl := range crawls {
		delete(communityMap, crawl.CommunityUID)
	}
	for _, community := range communityMap {
		if err = app.EntClient.UrlCrawl.Create().
			SetCommunityUID(community.UID).
			SetURL(fmt.Sprintf("https://www.tabelf.com/detail/%s", community.UID)).
			Exec(ctx); err != nil {
			continue
		}
	}
}
