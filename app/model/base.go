package model

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/util/gconv"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// sql build where

func TestWhereBuild(t *testing.T) {
	cond, vals, err := whereBuild(map[string]interface{}{
		"name":   "ttt",
		"age in": []int{20, 19, 18},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cond)
	t.Log(vals)

	err = GetDB().Table("test").Where(cond, vals...).Struct(&Tag{})
	if err != nil && err != sql.ErrNoRows {
		t.Fatal(err)
	}
}

func ModelSearch(mod modelType, maps map[string]interface{}) (query *gdb.Model, err error) {
	query = Model(mod)
	fields, fidelOk := maps["field"]
	page, pageOk := maps["page"]
	size, sizeOk := maps["size"]
	order, orderOk := maps["order"].(string)
	delete(maps, "page")
	delete(maps, "size")
	delete(maps, "order")

	if fidelOk {
		query = query.Fields(gconv.String(fields))
	}

	if len(maps) > 0 {
		cond, values, err := whereBuild(maps)
		if err != nil {
			return nil, err
		}

		query = query.Where(cond, values...)
	}

	if pageOk && sizeOk {
		query = query.Offset(gconv.Int(page)).Limit(gconv.Int(size))
	}

	if orderOk {
		query = query.OrderBy(order)
	}

	return query, nil
}

func whereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}
		strings.Join(ks, ",")
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
		}
	}
	return
}
