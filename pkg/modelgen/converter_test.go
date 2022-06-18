package modelgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vx416/modelgen/pkg/setting"
)

var createTable = `
CREATE TABLE crypto_balances (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL DEFAULT 0,
    checksum VARCHAR(255) NOT NULL DEFAULT '',
    currency VARCHAR(100),
    currency_decimal SMALLINT,
    i_amount DECIMAL(65, 0),
    i_locked_amount DECIMAL(65, 0),
    created_ati INT NOT NULL DEFAULT 0,
    updated_ati INT NOT NULL DEFAULT 0,
    UNIQUE KEY idx_user_id_currency (user_id, currency),
    PRIMARY KEY (id)
) COMMENT='user 虛擬幣餘額';
 `

func TestConverter(t *testing.T) {
	ddl, err := ParseDDL(createTable)
	require.NoError(t, err, "parse failed")
	assert.Equal(t, ddl.TableName, "point_third_orders")
	assert.Len(t, ddl.Columns, 16)
	cols, err := Converter(ddl.Columns, &setting.Settings{})
	require.NoError(t, err, "converter failed")
	t.Log(cols[0].String())
	assert.Len(t, cols, 16)
	assert.NotEmpty(t, cols[0].String())
}
