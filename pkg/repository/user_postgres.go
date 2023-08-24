package repository

//func (r *AuthPostgres) GetUser(username, password string) (project.User, error) {
//	var user project.User
//	query := "SELECT id FROM user_all WHERE username=$1 AND password_hash=$2"
//	err := r.db.Get(&user, query, username, password)
//	if err != nil {
//		logrus.Errorf("error in auth_postres getUser, err: %s", err.Error())
//	}
//
//	return user, nil
//}
