// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
)

func newProblemCategory(db *gorm.DB, opts ...gen.DOOption) problemCategory {
	_problemCategory := problemCategory{}

	_problemCategory.problemCategoryDo.UseDB(db, opts...)
	_problemCategory.problemCategoryDo.UseModel(&model.ProblemCategory{})

	tableName := _problemCategory.problemCategoryDo.TableName()
	_problemCategory.ALL = field.NewAsterisk(tableName)
	_problemCategory.ID = field.NewInt32(tableName, "id")
	_problemCategory.ProblemID = field.NewInt32(tableName, "problem_id")
	_problemCategory.CategoryID = field.NewInt32(tableName, "category_id")
	_problemCategory.CreatedAt = field.NewTime(tableName, "created_at")
	_problemCategory.UpdatedAt = field.NewTime(tableName, "updated_at")
	_problemCategory.DeletedAt = field.NewField(tableName, "deleted_at")

	_problemCategory.fillFieldMap()

	return _problemCategory
}

type problemCategory struct {
	problemCategoryDo

	ALL        field.Asterisk
	ID         field.Int32
	ProblemID  field.Int32
	CategoryID field.Int32
	CreatedAt  field.Time  // 创建时间
	UpdatedAt  field.Time  // 修改时间
	DeletedAt  field.Field // 删除时间(软删除)

	fieldMap map[string]field.Expr
}

func (p problemCategory) Table(newTableName string) *problemCategory {
	p.problemCategoryDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p problemCategory) As(alias string) *problemCategory {
	p.problemCategoryDo.DO = *(p.problemCategoryDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *problemCategory) updateTableName(table string) *problemCategory {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.ProblemID = field.NewInt32(table, "problem_id")
	p.CategoryID = field.NewInt32(table, "category_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *problemCategory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *problemCategory) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["problem_id"] = p.ProblemID
	p.fieldMap["category_id"] = p.CategoryID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p problemCategory) clone(db *gorm.DB) problemCategory {
	p.problemCategoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p problemCategory) replaceDB(db *gorm.DB) problemCategory {
	p.problemCategoryDo.ReplaceDB(db)
	return p
}

type problemCategoryDo struct{ gen.DO }

type IProblemCategoryDo interface {
	gen.SubQuery
	Debug() IProblemCategoryDo
	WithContext(ctx context.Context) IProblemCategoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProblemCategoryDo
	WriteDB() IProblemCategoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProblemCategoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProblemCategoryDo
	Not(conds ...gen.Condition) IProblemCategoryDo
	Or(conds ...gen.Condition) IProblemCategoryDo
	Select(conds ...field.Expr) IProblemCategoryDo
	Where(conds ...gen.Condition) IProblemCategoryDo
	Order(conds ...field.Expr) IProblemCategoryDo
	Distinct(cols ...field.Expr) IProblemCategoryDo
	Omit(cols ...field.Expr) IProblemCategoryDo
	Join(table schema.Tabler, on ...field.Expr) IProblemCategoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProblemCategoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProblemCategoryDo
	Group(cols ...field.Expr) IProblemCategoryDo
	Having(conds ...gen.Condition) IProblemCategoryDo
	Limit(limit int) IProblemCategoryDo
	Offset(offset int) IProblemCategoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProblemCategoryDo
	Unscoped() IProblemCategoryDo
	Create(values ...*model.ProblemCategory) error
	CreateInBatches(values []*model.ProblemCategory, batchSize int) error
	Save(values ...*model.ProblemCategory) error
	First() (*model.ProblemCategory, error)
	Take() (*model.ProblemCategory, error)
	Last() (*model.ProblemCategory, error)
	Find() ([]*model.ProblemCategory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProblemCategory, err error)
	FindInBatches(result *[]*model.ProblemCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProblemCategory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProblemCategoryDo
	Assign(attrs ...field.AssignExpr) IProblemCategoryDo
	Joins(fields ...field.RelationField) IProblemCategoryDo
	Preload(fields ...field.RelationField) IProblemCategoryDo
	FirstOrInit() (*model.ProblemCategory, error)
	FirstOrCreate() (*model.ProblemCategory, error)
	FindByPage(offset int, limit int) (result []*model.ProblemCategory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProblemCategoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p problemCategoryDo) Debug() IProblemCategoryDo {
	return p.withDO(p.DO.Debug())
}

func (p problemCategoryDo) WithContext(ctx context.Context) IProblemCategoryDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p problemCategoryDo) ReadDB() IProblemCategoryDo {
	return p.Clauses(dbresolver.Read)
}

func (p problemCategoryDo) WriteDB() IProblemCategoryDo {
	return p.Clauses(dbresolver.Write)
}

func (p problemCategoryDo) Session(config *gorm.Session) IProblemCategoryDo {
	return p.withDO(p.DO.Session(config))
}

func (p problemCategoryDo) Clauses(conds ...clause.Expression) IProblemCategoryDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p problemCategoryDo) Returning(value interface{}, columns ...string) IProblemCategoryDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p problemCategoryDo) Not(conds ...gen.Condition) IProblemCategoryDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p problemCategoryDo) Or(conds ...gen.Condition) IProblemCategoryDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p problemCategoryDo) Select(conds ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p problemCategoryDo) Where(conds ...gen.Condition) IProblemCategoryDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p problemCategoryDo) Order(conds ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p problemCategoryDo) Distinct(cols ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p problemCategoryDo) Omit(cols ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p problemCategoryDo) Join(table schema.Tabler, on ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p problemCategoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p problemCategoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p problemCategoryDo) Group(cols ...field.Expr) IProblemCategoryDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p problemCategoryDo) Having(conds ...gen.Condition) IProblemCategoryDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p problemCategoryDo) Limit(limit int) IProblemCategoryDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p problemCategoryDo) Offset(offset int) IProblemCategoryDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p problemCategoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProblemCategoryDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p problemCategoryDo) Unscoped() IProblemCategoryDo {
	return p.withDO(p.DO.Unscoped())
}

func (p problemCategoryDo) Create(values ...*model.ProblemCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p problemCategoryDo) CreateInBatches(values []*model.ProblemCategory, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p problemCategoryDo) Save(values ...*model.ProblemCategory) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p problemCategoryDo) First() (*model.ProblemCategory, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProblemCategory), nil
	}
}

func (p problemCategoryDo) Take() (*model.ProblemCategory, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProblemCategory), nil
	}
}

func (p problemCategoryDo) Last() (*model.ProblemCategory, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProblemCategory), nil
	}
}

func (p problemCategoryDo) Find() ([]*model.ProblemCategory, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProblemCategory), err
}

func (p problemCategoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProblemCategory, err error) {
	buf := make([]*model.ProblemCategory, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p problemCategoryDo) FindInBatches(result *[]*model.ProblemCategory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p problemCategoryDo) Attrs(attrs ...field.AssignExpr) IProblemCategoryDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p problemCategoryDo) Assign(attrs ...field.AssignExpr) IProblemCategoryDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p problemCategoryDo) Joins(fields ...field.RelationField) IProblemCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p problemCategoryDo) Preload(fields ...field.RelationField) IProblemCategoryDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p problemCategoryDo) FirstOrInit() (*model.ProblemCategory, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProblemCategory), nil
	}
}

func (p problemCategoryDo) FirstOrCreate() (*model.ProblemCategory, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProblemCategory), nil
	}
}

func (p problemCategoryDo) FindByPage(offset int, limit int) (result []*model.ProblemCategory, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p problemCategoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p problemCategoryDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p problemCategoryDo) Delete(models ...*model.ProblemCategory) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *problemCategoryDo) withDO(do gen.Dao) *problemCategoryDo {
	p.DO = *do.(*gen.DO)
	return p
}
