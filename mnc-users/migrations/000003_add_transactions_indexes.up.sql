CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_target_user ON transactions(target_user);
CREATE INDEX idx_transactions_transaction_type ON transactions(transaction_type);