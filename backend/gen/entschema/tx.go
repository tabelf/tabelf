// Code generated by entc, DO NOT EDIT.

package entschema

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Admin is the client for interacting with the Admin builders.
	Admin *AdminClient
	// Collaboration is the client for interacting with the Collaboration builders.
	Collaboration *CollaborationClient
	// Community is the client for interacting with the Community builders.
	Community *CommunityClient
	// CommunityCategory is the client for interacting with the CommunityCategory builders.
	CommunityCategory *CommunityCategoryClient
	// CommunityMeta is the client for interacting with the CommunityMeta builders.
	CommunityMeta *CommunityMetaClient
	// Feedback is the client for interacting with the Feedback builders.
	Feedback *FeedbackClient
	// Focus is the client for interacting with the Focus builders.
	Focus *FocusClient
	// GoodArticle is the client for interacting with the GoodArticle builders.
	GoodArticle *GoodArticleClient
	// GoodArticleCategory is the client for interacting with the GoodArticleCategory builders.
	GoodArticleCategory *GoodArticleCategoryClient
	// GoodArticleHot is the client for interacting with the GoodArticleHot builders.
	GoodArticleHot *GoodArticleHotClient
	// GoodArticleMeta is the client for interacting with the GoodArticleMeta builders.
	GoodArticleMeta *GoodArticleMetaClient
	// Invite is the client for interacting with the Invite builders.
	Invite *InviteClient
	// Message is the client for interacting with the Message builders.
	Message *MessageClient
	// Notice is the client for interacting with the Notice builders.
	Notice *NoticeClient
	// PayOrder is the client for interacting with the PayOrder builders.
	PayOrder *PayOrderClient
	// PayOrderFeedback is the client for interacting with the PayOrderFeedback builders.
	PayOrderFeedback *PayOrderFeedbackClient
	// PersonalFolder is the client for interacting with the PersonalFolder builders.
	PersonalFolder *PersonalFolderClient
	// Recharge is the client for interacting with the Recharge builders.
	Recharge *RechargeClient
	// ShareLink is the client for interacting with the ShareLink builders.
	ShareLink *ShareLinkClient
	// Station is the client for interacting with the Station builders.
	Station *StationClient
	// StationCategory is the client for interacting with the StationCategory builders.
	StationCategory *StationCategoryClient
	// StationMeta is the client for interacting with the StationMeta builders.
	StationMeta *StationMetaClient
	// Team is the client for interacting with the Team builders.
	Team *TeamClient
	// TeamFolder is the client for interacting with the TeamFolder builders.
	TeamFolder *TeamFolderClient
	// TeamGroup is the client for interacting with the TeamGroup builders.
	TeamGroup *TeamGroupClient
	// UrlCrawl is the client for interacting with the UrlCrawl builders.
	UrlCrawl *UrlCrawlClient
	// WebLink is the client for interacting with the WebLink builders.
	WebLink *WebLinkClient
	// Workspace is the client for interacting with the Workspace builders.
	Workspace *WorkspaceClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.Account = NewAccountClient(tx.config)
	tx.Admin = NewAdminClient(tx.config)
	tx.Collaboration = NewCollaborationClient(tx.config)
	tx.Community = NewCommunityClient(tx.config)
	tx.CommunityCategory = NewCommunityCategoryClient(tx.config)
	tx.CommunityMeta = NewCommunityMetaClient(tx.config)
	tx.Feedback = NewFeedbackClient(tx.config)
	tx.Focus = NewFocusClient(tx.config)
	tx.GoodArticle = NewGoodArticleClient(tx.config)
	tx.GoodArticleCategory = NewGoodArticleCategoryClient(tx.config)
	tx.GoodArticleHot = NewGoodArticleHotClient(tx.config)
	tx.GoodArticleMeta = NewGoodArticleMetaClient(tx.config)
	tx.Invite = NewInviteClient(tx.config)
	tx.Message = NewMessageClient(tx.config)
	tx.Notice = NewNoticeClient(tx.config)
	tx.PayOrder = NewPayOrderClient(tx.config)
	tx.PayOrderFeedback = NewPayOrderFeedbackClient(tx.config)
	tx.PersonalFolder = NewPersonalFolderClient(tx.config)
	tx.Recharge = NewRechargeClient(tx.config)
	tx.ShareLink = NewShareLinkClient(tx.config)
	tx.Station = NewStationClient(tx.config)
	tx.StationCategory = NewStationCategoryClient(tx.config)
	tx.StationMeta = NewStationMetaClient(tx.config)
	tx.Team = NewTeamClient(tx.config)
	tx.TeamFolder = NewTeamFolderClient(tx.config)
	tx.TeamGroup = NewTeamGroupClient(tx.config)
	tx.UrlCrawl = NewUrlCrawlClient(tx.config)
	tx.WebLink = NewWebLinkClient(tx.config)
	tx.Workspace = NewWorkspaceClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: Account.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)
