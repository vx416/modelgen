package modelgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vx416/modelgen/pkg/parser"
	"github.com/vx416/modelgen/pkg/setting"
)

var createTable = `
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
 `

func TestConverter(t *testing.T) {
	ddl, err := parser.ParseDDL(createTable)
	require.NoError(t, err, "parse failed")
	assert.Equal(t, ddl.TableName, "point_third_orders")
	assert.Len(t, ddl.Columns, 16)
	cols, err := Converter(ddl.Columns, &setting.Settings{})
	require.NoError(t, err, "converter failed")
	t.Log(cols[0].String())
	assert.Len(t, cols, 16)
	assert.NotEmpty(t, cols[0].String())
}
