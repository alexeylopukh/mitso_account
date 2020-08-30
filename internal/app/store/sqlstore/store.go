package sqlstore

import (
	"database/sql"
	"mitsoСhat/internal/app/model"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {

	sqlStore := &Store{
		db: db,
	}
	return sqlStore
}

func (s *Store) GetStudentInstallmentPlan(login string, pass string) (error, []*model.InstallmentPlan) {
	plans := []*model.InstallmentPlan{}
	rows, err := s.db.Query("select date, sum_to_pay, participation_fee from installment_plan where login = ?", login)
	if err != nil {
		return err, nil
	}
	defer rows.Close()


	for rows.Next() {
		var plan model.InstallmentPlan
		if err := rows.Scan(&plan.Date, &plan.SumToPay, &plan.ParticipationFee); err != nil {
			return err, nil
		}
		plans = append(plans, &plan)
	}
	return nil, plans
}

func (s *Store) GetStudent(login string, pass string) (error, *model.Student) {
	var student = model.Student{}
	err := s.db.QueryRow("select Login, Student from students where Login = ?", login).Scan(
		&student.Login,
		&student.Fio,
	)
	if err != nil {
		return err, nil
	}

	er := s.db.QueryRow("select balance, principal, fine, prolongation, date from data where ID = ?", login).Scan(
		&student.Balance,
		&student.Principal,
		&student.Fine,
		&student.Prolongation,
		&student.LastUpdate,
	)
	if er != nil {
		return er, nil
	}

	e, palans := s.GetStudentInstallmentPlan(login, pass)
	if e != nil {
		return nil, &student
	}

	student.Plans = palans

	return nil, &student
}
