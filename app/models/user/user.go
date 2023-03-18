type UserModel struct {
	dao *UserDAO
}

func NewUserModel(dao *UserDAO) *UserModel {
	return &UserModel{
		dao: dao,
	}
}

func (m *UserModel) CreateUser(user *UserDTO) error {
	_, err := m.dao.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) GetUserByID(id string) (*UserDTO, error) {
	user, err := m.dao.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserModel) UpdateUser(user *UserDTO) error {
	_, err := m.dao.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) DeleteUser(id string) error {
	_, err := m.dao.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}