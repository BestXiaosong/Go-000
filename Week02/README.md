## 作业
我们在数据库操作的时候，比如 `dao` 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 `Wrap` 这个 `error`，抛给上层。为什么？应该怎么做请写出代码

## 回答
    应该 Wrap error 抛给上层 dao层应只执行逻辑,具体错误判断及处理应当由调用方判断后自行处理

## 伪代码：
```go
// DAO 层

type Dao struct {}


func New() *Dao {
	return &Dao{ }
}

func (d *Dao) FindById(id int) (user *model.User, err error) {
	err = DB.Table("user").Where("id = ?", id).Find(user).Error
	
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("FindByIDErr:%v", id))
	}
	return
}
```

```go
// Service 层

type Service struct {  }


func (s *Service) FindById(id int) (user *model.User, err error) {
	user, err = dao.FindById(id)
	if errors.Is(err, dao.ErrRecordNotFound) {
		user = dao.GetDefault()
		err = nil
	}
	return
}
```
