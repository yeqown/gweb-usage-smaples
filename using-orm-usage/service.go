package main

// NewUser to create a record in Mysql
func NewUser(ui *User) error {
	uic := NewUserColl()
	if err := uic.Create(ui).Error; err != nil {
		return err
	}
	return nil
}

// NewRecipe to create a record in Mongo
func NewRecipe(r *Recipe) error {
	rc := NewRecipeDetailColl()
	if err := rc.Insert(r); err != nil {
		return err
	}
	return nil
}
