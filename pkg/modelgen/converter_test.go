package modelgen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var createTable = `
CREATE TABLE IF NOT EXISTS oauth2_third_parties (
    id bigint unsigned NOT NULL AUTO_INCREMENT,
    identity_tag varchar(50) NOT NULL DEFAULT '' UNIQUE,
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 COMMENT='oauth2 第三方平台列表' DEFAULT CHARSET=utf8;
 `

func TestConverter(t *testing.T) {
	_, err := ParseDDL(createTable)
	require.NoError(t, err, "parse failed")
	// assert.Equal(t, ddl.TableName, "point_third_orders")
	// assert.Len(t, ddl.Columns, 16)
	// cols, err := Converter(ddl.Columns, &setting.Settings{})
	// require.NoError(t, err, "converter failed")
	// t.Log(cols[0].String())
	// assert.Len(t, cols, 16)
	// assert.NotEmpty(t, cols[0].String())
}
