package model

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/lingfohn/lime/ent"
	"github.com/lingfohn/lime/ent/predicate"
	"github.com/lingfohn/lime/ent/user"
	"github.com/lingfohn/lime/storage"
)

type UserManager struct {
	db *ent.Client
}

func NewUserManager() *UserManager {
	return &UserManager{db: storage.GetDB()}
}

func (um *UserManager) Query(q QueryMap) (count int, users ent.Users, err error) {
	userQuery := um.db.User.
		Query().
		Where(um.predicate(q)...)
	if count, err = userQuery.Count(context.TODO()); err != nil {
		return
	}
	users, err = userQuery.
		Order(q.Order()...).
		Offset(q.OffSet()).
		Limit(q.Limit()).
		All(context.TODO())
	return
}

func (u *UserManager) predicate(qm QueryMap) []predicate.User {
	var query []predicate.User
	if val, ok := qm.GetString(user.FieldName); ok {
		query = append(query, user.NameContains(val))
	}
	if val, ok := qm.GetString(user.FieldEmail); ok {
		query = append(query, user.EmailContains(val))
	}
	if val, ok := qm.GetString(user.FieldTelephone); ok {
		query = append(query, user.TelephoneContains(val))
	}
	if val, ok := qm.GetInt(user.FieldStatus); ok {
		query = append(query, user.StatusEQ(val))
	}
	return query
}

func (um *UserManager) Create(u *ent.User) (user *ent.User, err error) {
	var password string = "1"
	md5sum := md5.Sum([]byte(password))
	user, err = um.db.User.
		Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetStatus(u.Status).
		SetPassword(fmt.Sprintf("%x", md5sum)).
		Save(context.TODO())
	// TODO 这里要加事务 还有其它逻辑
	return
}

func (um *UserManager) GetUserByUsername(username string) (u *ent.User, err error) {
	return um.db.User.Query().Where(user.UsernameEQ(username)).Only(context.TODO())
}

func (um *UserManager) UpdateUser(id int, u ent.User) (*ent.User, error) {
	return um.db.User.UpdateOneID(id).
		SetEmail(u.Email).
		SetName(u.Name).
		SetTelephone(u.Telephone).
		SetStatus(u.Status).
		Save(context.TODO())
}

func (um *UserManager) Delete(id int) error {
	return um.db.User.DeleteOneID(id).Exec(context.TODO())
}
