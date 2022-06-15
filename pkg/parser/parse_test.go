package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var createTable = `
-- +goose Up
CREATE TABLE point_third_orders (
	id bigint(20) NOT NULL AUTO_INCREMENT,
	serial varchar(50) NOT NULL DEFAULT '',
	source_kind int NOT NULL DEFAULT 0 COMMENT 'bitmap for 全家 or BitoWallet',
	order_kind tinyint NOT NULL DEFAULT 0 COMMENT '1:限價單 2:市價單', 
	order_id bigint(20) NOT NULL DEFAULT 0 COMMENT '外站 order id',
	provider varchar(50) NOT NULL DEFAULT '' COMMENT '外站名稱 e.g bitopro, ftx',
	provider_key varchar(100) NOT NULL DEFAULT '' COMMENT '外站用的 api key',
	order_state tinyint NOT NULL DEFAULT 0 COMMENT '外站訂單狀態 1:處理中 2:完全成交 3:取消',
	ordered_samount Decimal(65,0) NOT NULL DEFAULT 0 COMMENT '下單時的顆數',
	ordered_sprice Decimal(65,0) NOT NULL DEFAULT 0 COMMENT '下單的價格',
	traded_samount Decimal(65,0) NOT NULL DEFAULT 0 COMMENT '成交的顆數',
	avg_sprice Decimal(65,0) NOT NULL DEFAULT 0 COMMENT '成交均價',
	twd_srate Decimal(65,0) NOT NULL DEFAULT 0 COMMENT '下單時匯率',
	created_ati INT NOT NULL DEFAULT 0,
	updated_ati INT NOT NULL DEFAULT 0,
	completed_ati INT NOT NULL DEFAULT 0 COMMENT '訂單完成時間'
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
 `

func TestParse(t *testing.T) {
	ddl, err := Parse(createTable)
	require.NoError(t, err, "parse failed")
	t.Log(ddl)
}
