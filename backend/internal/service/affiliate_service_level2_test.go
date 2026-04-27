//go:build unit

package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type affiliateAccrualCall struct {
	inviterID     int64
	inviteeUserID int64
	amount        float64
	freezeHours   int
	level         int
}

type affiliateRepoLevel2Stub struct {
	summaries         map[int64]*AffiliateSummary
	accruedByLevel    map[string]float64
	accrualCalls      []affiliateAccrualCall
	accrueErrByLevel  map[int]error
	listCustomEntries []AffiliateAdminEntry
}

func affiliateLevelKey(inviterID, inviteeUserID int64, level int) string {
	return fmt.Sprintf("%d:%d:%d", inviterID, inviteeUserID, level)
}

func cloneAffiliateSummary(in *AffiliateSummary) *AffiliateSummary {
	if in == nil {
		return nil
	}
	out := *in
	if in.InviterID != nil {
		v := *in.InviterID
		out.InviterID = &v
	}
	if in.AffRebateRatePercent != nil {
		v := *in.AffRebateRatePercent
		out.AffRebateRatePercent = &v
	}
	if in.AffRebateRateLevel2Percent != nil {
		v := *in.AffRebateRateLevel2Percent
		out.AffRebateRateLevel2Percent = &v
	}
	return &out
}

func (r *affiliateRepoLevel2Stub) EnsureUserAffiliate(_ context.Context, userID int64) (*AffiliateSummary, error) {
	if summary, ok := r.summaries[userID]; ok {
		return cloneAffiliateSummary(summary), nil
	}
	return nil, ErrAffiliateProfileNotFound
}

func (r *affiliateRepoLevel2Stub) GetAffiliateByCode(_ context.Context, _ string) (*AffiliateSummary, error) {
	return nil, ErrAffiliateProfileNotFound
}

func (r *affiliateRepoLevel2Stub) BindInviter(_ context.Context, _, _ int64) (bool, error) {
	return false, nil
}

func (r *affiliateRepoLevel2Stub) AccrueQuota(ctx context.Context, inviterID, inviteeUserID int64, amount float64, freezeHours int) (bool, error) {
	return r.AccrueQuotaAtLevel(ctx, inviterID, inviteeUserID, amount, freezeHours, 1)
}

func (r *affiliateRepoLevel2Stub) AccrueQuotaAtLevel(_ context.Context, inviterID, inviteeUserID int64, amount float64, freezeHours int, level int) (bool, error) {
	r.accrualCalls = append(r.accrualCalls, affiliateAccrualCall{
		inviterID:     inviterID,
		inviteeUserID: inviteeUserID,
		amount:        amount,
		freezeHours:   freezeHours,
		level:         level,
	})
	if err := r.accrueErrByLevel[level]; err != nil {
		return false, err
	}
	return true, nil
}

func (r *affiliateRepoLevel2Stub) GetAccruedRebateFromInvitee(_ context.Context, inviterID, inviteeUserID int64) (float64, error) {
	return r.GetAccruedRebateFromInviteeByLevel(context.Background(), inviterID, inviteeUserID, 1)
}

func (r *affiliateRepoLevel2Stub) GetAccruedRebateFromInviteeByLevel(_ context.Context, inviterID, inviteeUserID int64, level int) (float64, error) {
	if r.accruedByLevel == nil {
		return 0, nil
	}
	return r.accruedByLevel[affiliateLevelKey(inviterID, inviteeUserID, level)], nil
}

func (r *affiliateRepoLevel2Stub) ThawFrozenQuota(_ context.Context, _ int64) (float64, error) {
	return 0, nil
}

func (r *affiliateRepoLevel2Stub) TransferQuotaToBalance(_ context.Context, _ int64) (float64, float64, error) {
	return 0, 0, nil
}

func (r *affiliateRepoLevel2Stub) ListInvitees(_ context.Context, _ int64, _ int) ([]AffiliateInvitee, error) {
	return nil, nil
}

func (r *affiliateRepoLevel2Stub) UpdateUserAffCode(_ context.Context, _ int64, _ string) error {
	return nil
}

func (r *affiliateRepoLevel2Stub) ResetUserAffCode(_ context.Context, _ int64) (string, error) {
	return "", nil
}

func (r *affiliateRepoLevel2Stub) SetUserRebateRate(_ context.Context, _ int64, _ *float64) error {
	return nil
}

func (r *affiliateRepoLevel2Stub) SetUserRebateRateLevel2(_ context.Context, _ int64, _ *float64) error {
	return nil
}

func (r *affiliateRepoLevel2Stub) BatchSetUserRebateRate(_ context.Context, _ []int64, _ *float64) error {
	return nil
}

func (r *affiliateRepoLevel2Stub) ListUsersWithCustomSettings(_ context.Context, _ AffiliateAdminFilter) ([]AffiliateAdminEntry, int64, error) {
	return r.listCustomEntries, int64(len(r.listCustomEntries)), nil
}

func affiliateTestSettingService(values map[string]string) *SettingService {
	repo := newMockSettingRepo()
	for key, value := range values {
		_ = repo.Set(context.Background(), key, value)
	}
	return &SettingService{settingRepo: repo}
}

func affiliateInt64Ptr(v int64) *int64 {
	return &v
}

func affiliateFloat64Ptr(v float64) *float64 {
	return &v
}

func TestResolveRebateRateL2Percent_PerUserOverride(t *testing.T) {
	t.Parallel()
	svc := &AffiliateService{}

	require.InDelta(t, AffiliateRebateRateL2Default,
		svc.resolveRebateRateL2Percent(context.Background(), &AffiliateSummary{}), 1e-9)

	rate := 12.5
	require.InDelta(t, 12.5,
		svc.resolveRebateRateL2Percent(context.Background(), &AffiliateSummary{
			AffRebateRateLevel2Percent: &rate,
		}), 1e-9)

	zero := 0.0
	require.InDelta(t, 0.0,
		svc.resolveRebateRateL2Percent(context.Background(), &AffiliateSummary{
			AffRebateRateLevel2Percent: &zero,
		}), 1e-9)

	tooHigh := 999.0
	require.InDelta(t, AffiliateRebateRateMax,
		svc.resolveRebateRateL2Percent(context.Background(), &AffiliateSummary{
			AffRebateRateLevel2Percent: &tooHigh,
		}), 1e-9)
}

func TestAccrueInviteRebate_AppliesLevel1AndLevel2(t *testing.T) {
	t.Parallel()

	repo := &affiliateRepoLevel2Stub{
		summaries: map[int64]*AffiliateSummary{
			300: {UserID: 300, InviterID: affiliateInt64Ptr(200), CreatedAt: time.Now()},
			200: {UserID: 200, InviterID: affiliateInt64Ptr(100), CreatedAt: time.Now()},
		},
	}
	svc := &AffiliateService{
		repo: repo,
		settingService: affiliateTestSettingService(map[string]string{
			SettingKeyAffiliateEnabled:           "true",
			SettingKeyAffiliateRebateRate:        "20",
			SettingKeyAffiliateRebateRateL2:      "10",
			SettingKeyAffiliateRebateFreezeHours: "0",
		}),
	}

	total, err := svc.AccrueInviteRebate(context.Background(), 300, 100)
	require.NoError(t, err)
	require.InDelta(t, 30.0, total, 1e-9)
	require.Len(t, repo.accrualCalls, 2)
	require.Equal(t, affiliateAccrualCall{
		inviterID:     200,
		inviteeUserID: 300,
		amount:        20,
		freezeHours:   0,
		level:         1,
	}, repo.accrualCalls[0])
	require.Equal(t, affiliateAccrualCall{
		inviterID:     100,
		inviteeUserID: 300,
		amount:        10,
		freezeHours:   0,
		level:         2,
	}, repo.accrualCalls[1])
}

func TestAccrueInviteRebate_Level2UsesExclusiveOverride(t *testing.T) {
	t.Parallel()

	repo := &affiliateRepoLevel2Stub{
		summaries: map[int64]*AffiliateSummary{
			300: {
				UserID:    300,
				InviterID: affiliateInt64Ptr(200),
				CreatedAt: time.Now(),
			},
			200: {
				UserID:                     200,
				InviterID:                  affiliateInt64Ptr(100),
				AffRebateRateLevel2Percent: affiliateFloat64Ptr(5),
				CreatedAt:                  time.Now(),
			},
		},
	}
	svc := &AffiliateService{
		repo: repo,
		settingService: affiliateTestSettingService(map[string]string{
			SettingKeyAffiliateEnabled:      "true",
			SettingKeyAffiliateRebateRate:   "20",
			SettingKeyAffiliateRebateRateL2: "10",
		}),
	}

	total, err := svc.AccrueInviteRebate(context.Background(), 300, 100)
	require.NoError(t, err)
	require.InDelta(t, 25.0, total, 1e-9)
	require.Len(t, repo.accrualCalls, 2)
	require.Equal(t, 5.0, repo.accrualCalls[1].amount)
}

func TestAccrueInviteRebate_Level2CanApplyWhenLevel1Skips(t *testing.T) {
	t.Parallel()

	repo := &affiliateRepoLevel2Stub{
		summaries: map[int64]*AffiliateSummary{
			300: {UserID: 300, InviterID: affiliateInt64Ptr(200), CreatedAt: time.Now()},
			200: {UserID: 200, InviterID: affiliateInt64Ptr(100), CreatedAt: time.Now()},
		},
	}
	svc := &AffiliateService{
		repo: repo,
		settingService: affiliateTestSettingService(map[string]string{
			SettingKeyAffiliateEnabled:      "true",
			SettingKeyAffiliateRebateRate:   "0",
			SettingKeyAffiliateRebateRateL2: "10",
		}),
	}

	total, err := svc.AccrueInviteRebate(context.Background(), 300, 100)
	require.NoError(t, err)
	require.InDelta(t, 10.0, total, 1e-9)
	require.Len(t, repo.accrualCalls, 1)
	require.Equal(t, 2, repo.accrualCalls[0].level)
}

func TestAccrueInviteRebate_Level2FailureIsNonBlocking(t *testing.T) {
	t.Parallel()

	repo := &affiliateRepoLevel2Stub{
		summaries: map[int64]*AffiliateSummary{
			300: {UserID: 300, InviterID: affiliateInt64Ptr(200), CreatedAt: time.Now()},
			200: {UserID: 200, InviterID: affiliateInt64Ptr(100), CreatedAt: time.Now()},
		},
		accrueErrByLevel: map[int]error{
			2: errors.New("level2 insert failed"),
		},
	}
	svc := &AffiliateService{
		repo: repo,
		settingService: affiliateTestSettingService(map[string]string{
			SettingKeyAffiliateEnabled:      "true",
			SettingKeyAffiliateRebateRate:   "20",
			SettingKeyAffiliateRebateRateL2: "10",
		}),
	}

	total, err := svc.AccrueInviteRebate(context.Background(), 300, 100)
	require.NoError(t, err)
	require.InDelta(t, 20.0, total, 1e-9)
	require.Len(t, repo.accrualCalls, 2)
}
