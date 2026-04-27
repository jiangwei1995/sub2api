ALTER TABLE user_affiliates
    ADD COLUMN IF NOT EXISTS aff_rebate_rate_level2_percent DECIMAL(5,2);

ALTER TABLE user_affiliate_ledger
    ADD COLUMN IF NOT EXISTS level SMALLINT NOT NULL DEFAULT 1;
