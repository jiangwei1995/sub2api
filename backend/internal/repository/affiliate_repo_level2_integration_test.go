//go:build integration

package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/stretchr/testify/require"
)

func TestAffiliateRepository_Level2AccrualAndInviteeList(t *testing.T) {
	ctx := context.Background()
	tx := testEntTx(t)
	txCtx := dbent.NewTxContext(ctx, tx)
	client := tx.Client()

	repo := NewAffiliateRepository(client, integrationDB)

	l2Inviter := mustCreateUser(t, client, &service.User{
		Email:        fmt.Sprintf("affiliate-l2-%d-a@example.com", time.Now().UnixNano()),
		PasswordHash: "hash",
		Role:         service.RoleUser,
		Status:       service.StatusActive,
	})
	l1Inviter := mustCreateUser(t, client, &service.User{
		Email:        fmt.Sprintf("affiliate-l2-%d-b@example.com", time.Now().UnixNano()+1),
		PasswordHash: "hash",
		Role:         service.RoleUser,
		Status:       service.StatusActive,
	})
	invitee := mustCreateUser(t, client, &service.User{
		Email:        fmt.Sprintf("affiliate-l2-%d-c@example.com", time.Now().UnixNano()+2),
		PasswordHash: "hash",
		Role:         service.RoleUser,
		Status:       service.StatusActive,
	})

	_, err := repo.EnsureUserAffiliate(txCtx, l2Inviter.ID)
	require.NoError(t, err)
	_, err = repo.EnsureUserAffiliate(txCtx, l1Inviter.ID)
	require.NoError(t, err)
	_, err = repo.EnsureUserAffiliate(txCtx, invitee.ID)
	require.NoError(t, err)

	bound, err := repo.BindInviter(txCtx, l1Inviter.ID, l2Inviter.ID)
	require.NoError(t, err)
	require.True(t, bound)

	bound, err = repo.BindInviter(txCtx, invitee.ID, l1Inviter.ID)
	require.NoError(t, err)
	require.True(t, bound)

	applied, err := repo.AccrueQuota(txCtx, l2Inviter.ID, l1Inviter.ID, 5, 0)
	require.NoError(t, err)
	require.True(t, applied)

	applied, err = repo.AccrueQuotaAtLevel(txCtx, l1Inviter.ID, invitee.ID, 20, 0, 1)
	require.NoError(t, err)
	require.True(t, applied)

	applied, err = repo.AccrueQuotaAtLevel(txCtx, l2Inviter.ID, invitee.ID, 10, 0, 2)
	require.NoError(t, err)
	require.True(t, applied)

	level1Count := querySingleInt(t, txCtx, client,
		`SELECT COUNT(*) FROM user_affiliate_ledger WHERE user_id = $1 AND source_user_id = $2 AND action = 'accrue' AND level = 1`,
		l2Inviter.ID, l1Inviter.ID)
	require.Equal(t, 1, level1Count)

	level2Count := querySingleInt(t, txCtx, client,
		`SELECT COUNT(*) FROM user_affiliate_ledger WHERE user_id = $1 AND source_user_id = $2 AND action = 'accrue' AND level = 2`,
		l2Inviter.ID, invitee.ID)
	require.Equal(t, 1, level2Count)

	invitees, err := repo.ListInvitees(txCtx, l2Inviter.ID, 10)
	require.NoError(t, err)
	require.Len(t, invitees, 2)
	require.Equal(t, l1Inviter.ID, invitees[0].UserID)
	require.Equal(t, 1, invitees[0].Level)
	require.InDelta(t, 5.0, invitees[0].TotalRebate, 1e-9)
	require.Equal(t, invitee.ID, invitees[1].UserID)
	require.Equal(t, 2, invitees[1].Level)
	require.InDelta(t, 10.0, invitees[1].TotalRebate, 1e-9)
}

func TestAffiliateRepository_ListUsersWithCustomSettings_IncludesLevel2Override(t *testing.T) {
	ctx := context.Background()
	tx := testEntTx(t)
	txCtx := dbent.NewTxContext(ctx, tx)
	client := tx.Client()

	repo := NewAffiliateRepository(client, integrationDB)

	u := mustCreateUser(t, client, &service.User{
		Email:        fmt.Sprintf("affiliate-l2-custom-%d@example.com", time.Now().UnixNano()),
		PasswordHash: "hash",
		Role:         service.RoleUser,
		Status:       service.StatusActive,
	})

	rate := 7.5
	require.NoError(t, repo.SetUserRebateRateLevel2(txCtx, u.ID, &rate))

	summary, err := repo.EnsureUserAffiliate(txCtx, u.ID)
	require.NoError(t, err)
	require.NotNil(t, summary.AffRebateRateLevel2Percent)
	require.InDelta(t, 7.5, *summary.AffRebateRateLevel2Percent, 1e-9)

	entries, total, err := repo.ListUsersWithCustomSettings(txCtx, service.AffiliateAdminFilter{
		Page:     1,
		PageSize: 20,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), total)
	require.Len(t, entries, 1)
	require.Equal(t, u.ID, entries[0].UserID)
	require.NotNil(t, entries[0].AffRebateRateLevel2Percent)
	require.InDelta(t, 7.5, *entries[0].AffRebateRateLevel2Percent, 1e-9)

	require.NoError(t, repo.SetUserRebateRateLevel2(txCtx, u.ID, nil))
	cleared, err := repo.EnsureUserAffiliate(txCtx, u.ID)
	require.NoError(t, err)
	require.Nil(t, cleared.AffRebateRateLevel2Percent)
}
