-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id         INT NOT NULL AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    created_ati INT NOT NULL DEFAULT 0,
    updated_ati INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS crypto_balances (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL DEFAULT 0,
    checksum VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用來驗證餘額沒被亂更改',
    currency VARCHAR(100) COMMENT 'ETH, BTC, USDT, MATIC, DOGE, etc...',
    currency_decimal SMALLINT COMMENT 'i_amount 的小數點位數',
    i_amount DECIMAL(65, 0),
    i_locked_amount DECIMAL(65, 0) COMMENT '提款時提前鎖定餘額',
    created_ati INT NOT NULL DEFAULT 0,
    updated_ati INT NOT NULL DEFAULT 0,
    UNIQUE (user_id, currency),
    PRIMARY KEY (id)
) COMMENT='user 虛擬幣餘額';

CREATE TABLE IF NOT EXISTS cybavo_wallets (
    id                INT NOT NULL AUTO_INCREMENT,
    parent_wallet_id  INT,
    name              VARCHAR(200) NOT NULL DEFAULT '',
    currency          VARCHAR(200) NOT NULL DEFAULT '' COMMENT 'ETH, BTC, USDT, MATIC, DOGE, etc...',
    kind              SMALLINT     NOT NULL COMMENT '1:加值, 2:提領, 3:委派',
    currency_id       INT NOT NULL DEFAULT 0,
    contract_address  VARCHAR(200) NOT NULL DEFAULT '',
    wallet_id         INT          NOT NULL DEFAULT 0,
    api_code          VARCHAR(200) NOT NULL DEFAULT '',
    api_secret        VARCHAR(200) NOT NULL DEFAULT '',
    refresh_token     VARCHAR(200) NOT NULL DEFAULT '',
    withdrawal_prefix VARCHAR(200) NOT NULL DEFAULT '',
    updated_ati        INT NOT NULL DEFAULT 0,
    created_ati        INT NOT NULL DEFAULT 0,
    UNIQUE (api_code),
    UNIQUE (currency, kind),
    PRIMARY KEY (id)
) COMMENT='cybavo wallet 設定';

CREATE TABLE IF NOT EXISTS eth_nft_ownerships (
    id INT NOT NULL AUTO_INCREMENT,
    token_id INT NOT NULL DEFAULT 0,
    user_id INT NOT NULL DEFAULT 0,
    state TINYINT NOT NULL DEFAULT 0 COMMENT '1:minting, 2:minted, 3:transfering, 4:tranfered, 5:staking',
    checksum VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用來驗證ownership沒被亂更改',
    contract_address VARCHAR(255) NOT NULL DEFAULT '',
    token_url VARCHAR(255) NOT NULL DEFAULT '',
    updated_ati INT NOT NULL DEFAULT 0,
    created_ati INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) COMMENT='user eth 鏈 nft 擁有權';


CREATE TABLE IF NOT EXISTS eth_transactions (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL DEFAULT 0,
    serial VARCHAR(100) NOT NULL DEFAULT '', 
    state SMALLINT NOT NULL DEFAULT 0 COMMENT '0:初始化 1:驗證中 2:已上鏈 3:已完成 4:失敗 5:人工審查',
    kind SMALLINT NOT NULL DEFAULT 0 COMMENT '1: 虛擬幣, 2: NFT, 3:智能合約invoke',
    sub_kind SMALLINT NOT NULL DEFAULT 0 COMMENT '1: 內轉, 2:Bitopro內轉',
    currency VARCHAR(200) NOT NULL DEFAULT '' COMMENT '交易幣種',
    currency_decimal SMALLINT NOT NULL DEFAULT 0 COMMENT '交易幣種小數點位數',
    i_amount DECIMAL(65, 0) NOT NULL DEFAULT 0 COMMENT '交易金額 (unit: wei)',
    i_twd_amount DECIMAL(65, 0) NOT NULL DEFAULT 0 COMMENT '交易台幣金額 (unit: wei)',
    fee_currency VARCHAR(200) NOT NULL DEFAULT '' COMMENT '手續費幣種',
    fee_currency_decimal SMALLINT NOT NULL DEFAULT 0 COMMENT '手續費幣種小數點位數',
    i_fee DECIMAL(65, 0) NOT NULL DEFAULT 0 COMMENT '站方交易手續費 (unit: wei)',
    i_twd_fee DECIMAL(65, 0) NOT NULL DEFAULT 0 COMMENT '站方交易台幣手續費 (unit: wei)',
    nft_ownership_id INT NOT NULL DEFAULT 0,
    token_id DECIMAL(65, 0) NOT NULL DEFAULT 0,
    contract_address VARCHAR(255) NOT NULL DEFAULT '',
    abi_fnc VARCHAR(255) NOT NULL DEFAULT '',
    from_wallet VARCHAR(255) NOT NULL DEFAULT '',
    to_wallet VARCHAR(255) NOT NULL DEFAULT '',
    tx_id VARCHAR(100) DEFAULT 0,
    tx_index int(11) DEFAULT 0,
    chain_fee VARCHAR(255) NOT NULL DEFAULT 0 COMMENT '鏈上手續費 (unit: wei)',
    chain_i_amount DECIMAL(65, 0) NOT NULL DEFAULT 0 COMMENT '鏈上交易金額',
    confirm_count SMALLINT NOT NULL DEFAULT 0,
    chain_ati INT NOT NULL DEFAULT 0,
    created_ati INT NOT NULL DEFAULT 0,
    updated_ati INT NOT NULL DEFAULT 0,
    source_ip VARCHAR(100) NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) COMMENT='user eth 鏈 transaction 紀錄';

CREATE TABLE IF NOT EXISTS eth_wallets (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL DEFAULT 0,
    address VARCHAR(200) NOT NULL DEFAULT '',
    kind    SMALLINT     NOT NULL COMMENT '1:加值, 3:委派',
    cybavo_address_index INT NOT NULL DEFAULT 0,
    cybavo_wallet_id INT NOT NULL DEFAULT 0,
    state TINYINT NOT NULL DEFAULT 0 COMMENT '0:初始化 1:已配置',
    updated_ati INT NOT NULL DEFAULT 0,
    created_ati INT NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
) COMMENT='user eth 鏈 address';

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS crypto_balances;
DROP TABLE IF EXISTS cybavo_wallets;
DROP TABLE IF EXISTS eth_transactions;
DROP TABLE IF EXISTS eth_wallets;
DROP TABLE IF EXISTS eth_nft_ownerships;