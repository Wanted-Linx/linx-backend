package repository

import (
	"context"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/Wanted-Linx/linx-backend/api/ent/company"
	"github.com/Wanted-Linx/linx-backend/api/ent/tasktype"
	"github.com/pkg/errors"
)

type companyRepository struct {
	db *ent.Client
}

func NewCompanyRepository(db *ent.Client) domain.CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Save(reqCompany *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.Create().
		SetID(reqCompany.ID).
		SetName(reqCompany.Name).
		SetBusinessNumber(reqCompany.BusinessNumber).
		SetUser(reqCompany.Edges.User).
		SetNillableProfileImage(reqCompany.ProfileImage).
		SetNillableDescription(reqCompany.Description).
		SetNillableAddress(reqCompany.Address).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	user, err := c.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	c.Edges.User = user
	return c, nil
}

func (r *companyRepository) GetByID(companyID int, reqCompany *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.Query().
		Where(company.ID(companyID)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *companyRepository) UploadProfileImage(reqCompany *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.UpdateOneID(reqCompany.ID).
		SetNillableProfileImage(reqCompany.ProfileImage).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func (r *companyRepository) UpdateProfile(reqCompany *ent.Company) (*ent.Company, error) {
	c, err := r.db.Company.UpdateOneID(reqCompany.ID).
		SetNillableAddress(reqCompany.Address).
		SetNillableHompage(reqCompany.Hompage).
		SetNillableDescription(reqCompany.Description).Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user, err := c.QueryUser().First(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	c.Edges.User = user

	return c, nil
}

func (r *companyRepository) GetAll(limit, offset int) ([]*ent.Company, error) {
	c, err := r.db.Company.Query().WithUser().
		Limit(limit).Offset(offset).All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}

func (r *companyRepository) GetAllTasks(companyID int) ([]*ent.TaskType, error) {
	tasks, err := r.db.TaskType.Query().
		Where(tasktype.HasCompanyWith(company.ID(companyID))).All(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tasks, nil
}

func (r *companyRepository) SaveTasks(c *ent.Company, taskType *ent.TaskType) (*ent.TaskType, error) {
	task, err := r.db.TaskType.Create().
		SetType(taskType.Type).
		SetCompany(c).
		Save(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return task, nil
}
