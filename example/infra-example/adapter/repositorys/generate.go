package repositorys

import (
	"github.com/8treenet/freedom"
	"github.com/8treenet/freedom/example/infra-example/application/objects"
)

// findGoodsByPrimary .
func findGoodsByPrimary(rep freedom.GORMRepository, primary interface{}) (result objects.Goods, e error) {
	e = rep.DB().Find(&result, primary).Error
	return
}

// findGoodssByPrimarys .
func findGoodssByPrimarys(rep freedom.GORMRepository, primarys ...interface{}) (results []objects.Goods, e error) {
	e = rep.DB().Find(&results, primarys).Error
	return
}

// findGoods .
func findGoods(rep freedom.GORMRepository, query objects.Goods, builders ...freedom.QueryBuilder) (result objects.Goods, e error) {
	db := rep.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findGoodsByWhere .
func findGoodsByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (result objects.Goods, e error) {
	db := rep.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findGoodsByMap .
func findGoodsByMap(rep freedom.GORMRepository, query map[string]interface{}, builders ...freedom.QueryBuilder) (result objects.Goods, e error) {
	db := rep.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findGoodss .
func findGoodss(rep freedom.GORMRepository, query objects.Goods, builders ...freedom.QueryBuilder) (results []objects.Goods, e error) {
	db := rep.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findGoodssByWhere .
func findGoodssByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (results []objects.Goods, e error) {
	db := rep.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findGoodssByMap .
func findGoodssByMap(rep freedom.GORMRepository, query map[string]interface{}, builders ...freedom.QueryBuilder) (results []objects.Goods, e error) {
	db := rep.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// createGoods .
func createGoods(rep freedom.GORMRepository, object *objects.Goods) (rowsAffected int64, e error) {
	db := rep.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	return
}

// updateGoods .
func updateGoods(rep freedom.GORMRepository, object *objects.Goods) (affected int64, e error) {
	db := rep.DB().Model(object).Updates(object.Save())
	e = db.Error
	affected = db.RowsAffected
	return
}

// findOrderByPrimary .
func findOrderByPrimary(rep freedom.GORMRepository, primary interface{}) (result objects.Order, e error) {
	e = rep.DB().Find(&result, primary).Error
	return
}

// findOrdersByPrimarys .
func findOrdersByPrimarys(rep freedom.GORMRepository, primarys ...interface{}) (results []objects.Order, e error) {
	e = rep.DB().Find(&results, primarys).Error
	return
}

// findOrder .
func findOrder(rep freedom.GORMRepository, query objects.Order, builders ...freedom.QueryBuilder) (result objects.Order, e error) {
	db := rep.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findOrderByWhere .
func findOrderByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (result objects.Order, e error) {
	db := rep.DB()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findOrderByMap .
func findOrderByMap(rep freedom.GORMRepository, query map[string]interface{}, builders ...freedom.QueryBuilder) (result objects.Order, e error) {
	db := rep.DB().Where(query)
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = db.Limit(1).Order(builders[0].Order()).Find(&result).Error
	return
}

// findOrders .
func findOrders(rep freedom.GORMRepository, query objects.Order, builders ...freedom.QueryBuilder) (results []objects.Order, e error) {
	db := rep.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findOrdersByWhere .
func findOrdersByWhere(rep freedom.GORMRepository, query string, args []interface{}, builders ...freedom.QueryBuilder) (results []objects.Order, e error) {
	db := rep.DB()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findOrdersByMap .
func findOrdersByMap(rep freedom.GORMRepository, query map[string]interface{}, builders ...freedom.QueryBuilder) (results []objects.Order, e error) {
	db := rep.DB().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// createOrder .
func createOrder(rep freedom.GORMRepository, object *objects.Order) (rowsAffected int64, e error) {
	db := rep.DB().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	return
}

// updateOrder .
func updateOrder(rep freedom.GORMRepository, object *objects.Order) (affected int64, e error) {
	db := rep.DB().Model(object).Updates(object.Save())
	e = db.Error
	affected = db.RowsAffected
	return
}