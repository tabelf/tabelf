package tasks

import (
	"entgo.io/ent/dialect/sql"
	"tabelf/backend/gen/entschema"
	entadmin "tabelf/backend/gen/entschema/admin"
	entstation "tabelf/backend/gen/entschema/station"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/app"
)

// AutoCreateStation 收录管理员的weblink的链接到好站推荐.
func AutoCreateStation(jobCtx JobContext, config app.Config) {
	ctx := jobCtx.Context

	admins, err := app.EntClient.Admin.Query().Where(
		entadmin.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	for _, admin := range admins {
		offset, limit := 1, 20
		skipUID := make([]string, 0)
		for {
			links, err := app.EntClient.WebLink.Query().Where(
				entweblink.UserUID(admin.UserUID),
				entweblink.DeactivatedAtIsNil(),
				func(s *sql.Selector) {
					if len(skipUID) != 0 {
						entweblink.UIDNotIn(skipUID...)
					}
				},
				entweblink.FileType(app.URLFileType),
			).Offset((offset - 1) * limit).Limit(limit).All(ctx)
			if err != nil {
				return
			}
			stationBulk := make([]*entschema.StationCreate, 0)
			for _, link := range links {
				exist, err := app.EntClient.Station.Query().Where(
					func(s *sql.Selector) {
						s.Where(sql.Like(entstation.FieldLink, "%"+link.Link+"%"))
					},
				).Exist(ctx)
				if err != nil || exist {
					skipUID = append(skipUID, link.UID)
					continue
				}
				stationCreate := app.EntClient.Station.Create().
					SetTitle(link.Title).
					SetDescription(link.Description).
					SetLink(link.Link).
					SetIcon(link.Image).
					SetImage(link.Image).
					SetSource(link.Title).
					SetUserUID(link.UserUID).
					SetStatus(app.Hidden).
					SetCategoryUID("").
					SetTags(make([]string, 0))
				stationBulk = append(stationBulk, stationCreate)
			}
			if err = app.EntClient.Station.
				CreateBulk(stationBulk...).
				Exec(ctx); err != nil {
				return
			}
			if len(links) != limit {
				break
			}
			offset++
		}
	}
}
