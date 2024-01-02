package repo_transfers

import (
	"banktransfer/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type I_Repo_Transfers interface {
	// func ที่ต้องรับ struct เพราะต้องใช้งาน DB ไม่ใช้ใช้งาน interface
	Transaction_Postgres(func(*repo_Transfers) error) error

	// โอน
	Create_Transfer(dataTransfer *models.CreateTransferParams) error
	//ฝาก
	Create_Deposit(dataDeposit *models.Create_Deposit_and_Withdraw) error
	// ถอน
	Create_Withdraw(dataWithdraw *models.Create_Deposit_and_Withdraw) error

	// เช็คการโอนทั้งหมด ของบัญชีนั้นๆ (ตาม account) และเลือกช่วงเวลา
	GetTransfer_ById(idAccount int64, startTime, endTime string) ([]models.ListTransfers_ForOwner, error)

	// เช็คการโอนทั้งหมด ของเราทุกบัญชี (ตาม owner) และเลือกช่วงเวลา
	GetTransfer_ByOwner(ownerAccount, startTime, endTime string) ([]models.ListTransfers_ForOwner, error)

	// เอาไว้ดึงข้อมูลที่จะเช็คเป็นรายบัญชี เช่น อยากรู้ว่า
	// เราโอนเข้าบัญชี นาย ก ทั้งหมดกี่ครั้ง
	// และเลือกช่วงเวลา
	ListStatement(accountId int64) ([]models.Transfer, error)

	// หบัวจาก ฝาก,ถอน, โอน ต้องมาอัพเดท
	UpdateAccount_Blance([]models.UpdateAccountParams) error

	// **************************** Check Data  Core ****************
	CheckDataFor_Transfer(dataCheck *models.CreateTransferParams) error
	CheckDataFor_Deposit_and_Withdraw(dataCheck *models.Create_Deposit_and_Withdraw) error

	// Check data inside ****************

	// เช็ค account คนที่จะโอน หรือ เช็ค account คนจะฝากและถอน
	checkAccountOwner(owner string, id int64) error
	// เช็ค account คนที่จะรับโอน
	checkAccountTo(id int64) error
	// เช็ค สกุลเงิน
	checkCurrency(fromAccount, toAccount string) error
	// เช็ค จำนนเงินคนที่จะโอน
	checkAccountBalance(amount, balbnce int64) error
}

type repo_Transfers struct {
	db *gorm.DB
}

func New_Repo_User(db *gorm.DB) *repo_Transfers {
	return &repo_Transfers{
		db: db,
	}
}

// ฝากเงิน *******************
func (ra *repo_Transfers) Create_Deposit(dataDeposit *models.Create_Deposit_and_Withdraw) error {
	time.Now()
	// check data for Deposit
	erCheckData := ra.CheckDataFor_Deposit_and_Withdraw(dataDeposit)
	if erCheckData != nil {
		return erCheckData // ปกติจะ เป็น tx.Error แต่อันนี้ ได้ค่ามาเป็น tx.Error แล้ว
	}

	// ปั้นข้อมูลใหม่ สำหรับ Update balance
	new_dataDeposit := []models.UpdateAccountParams{
		{
			ID:      dataDeposit.AccountID,
			Balance: dataDeposit.Amount,
		},
	}

	//************* ทำ transaction *********************
	return ra.Transaction_Postgres(func(p *repo_Transfers) error {
		//create entries ********************
		dataEntries := models.Entry{
			// FromAccountID *********************
			AccountID: dataDeposit.AccountID,
			Amount:    dataDeposit.Amount,
			// CreatedAt: time.Now(), // มันใสค่าให้เองถึงเราๆม่ใสค่าให้ ***
			Entries_type: "deposit",
		}

		tx := p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			fmt.Println("entries naja")
			return tx.Error
		}

		// update accounts ****************
		errUpdate := ra.UpdateAccount_Blance(new_dataDeposit)
		if errUpdate != nil {
			return errUpdate
		}

		fmt.Println("Deposit Complete")
		return nil
	})

}

// ถอนเงิน *****************
func (ra *repo_Transfers) Create_Withdraw(dataWithdraw *models.Create_Deposit_and_Withdraw) error {
	// check data for Deposit
	erCheckData := ra.CheckDataFor_Deposit_and_Withdraw(dataWithdraw)
	if erCheckData != nil {
		return erCheckData // ปกติจะ เป็น tx.Error แต่อันนี้ ได้ค่ามาเป็น tx.Error แล้ว
	}

	// // ปั้นข้อมูลใหม่ สำหรับ Update balance
	new_dataDeposit := []models.UpdateAccountParams{
		{
			ID:      dataWithdraw.AccountID,
			Balance: -dataWithdraw.Amount,
		},
	}

	//************* ทำ transaction *********************
	return ra.Transaction_Postgres(func(p *repo_Transfers) error {
		//create entries ********************
		dataEntries := models.Entry{
			AccountID:    dataWithdraw.AccountID,
			Amount:       -dataWithdraw.Amount,
			CreatedAt:    time.Now(),
			Entries_type: "withdraw",
		}

		// create entries
		tx := p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			fmt.Println("entries naja")
			return tx.Error
		}

		// update accounts ****************
		errUpdate := ra.UpdateAccount_Blance(new_dataDeposit)
		if errUpdate != nil {
			return errUpdate
		}
		fmt.Println("Withdraw Complete")
		return nil
	})

}

func (ra *repo_Transfers) CheckDataFor_Deposit_and_Withdraw(dataCheck *models.Create_Deposit_and_Withdraw) error {

	// เช็ค account ที่จะฝากหรือ ถอนว่ามีหรือไม่
	// เช็ค owner จาก token
	// accountParams := new(models.CreateAccountParams)
	// tx := ra.db.Table("accounts").Where("id =?", dataCheck.AccountID).Where("owner =?", dataCheck.Owner).First(&accountParams)
	// if tx.Error != nil {
	// 	return tx.Error
	// }
	// return nil

	tx := ra.checkAccountOwner(dataCheck.Owner, dataCheck.AccountID)
	if tx != nil {
		return tx
	}
	return nil
}

func (ra *repo_Transfers) UpdateAccount_Blance(accountUpdate []models.UpdateAccountParams) error {

	// logic อันนี้ได้ แต่ดีรึป่าวไม่รู้
	// accountParams.Balance = accountParams.Balance + accountUpdate.Balance
	// tx = ra.db.Table("accounts").Where("id =?", accountUpdate.ID).Updates(&accountParams)
	// if tx.Error != nil {
	// 	return tx.Error
	// }

	// หรือใช้ query update  ไปเลย ใช้  blance = blance+ ค่าใหม่
	myQuery := `UPDATE  accounts SET  balance = balance+? WHERE  id = ?`
	for _, v := range accountUpdate {
		tx := ra.db.Exec(myQuery, v.Balance, v.ID)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func (rt *repo_Transfers) CheckDataFor_Transfer(dataCheck *models.CreateTransferParams) error {
	// ในแต่ละ func มันจะ  query ทุก func ซึ่งมันเยอะไป
	//  query ที่เดียวแล้วส่งค่าเข้าไปเช็คแทน

	// query From Account ******
	fromAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.FromAccountID).Where("owner=?", dataCheck.Owner).First(fromAccount)

	// query To Account ********
	toAccount := new(models.Account)
	rt.db.Table("accounts").Where("id=?", dataCheck.ToAccountID).First(toAccount)
	// เอาข้อมูลที่ได้ส่งเข้าไปใน func ต่างๆ เพื่อเช็ค error

	var transferErr error
	// เช็ค error ต่างๆ ******************
	// เช็ค owner, accoutTo, currency, blance
	if transferErr = rt.checkAccountOwner(fromAccount.Owner, fromAccount.ID); transferErr == nil {
		if transferErr = rt.checkAccountTo(fromAccount.ID); transferErr == nil {
			if transferErr = rt.checkCurrency(fromAccount.Currency, toAccount.Currency); transferErr == nil {
				if transferErr = rt.checkAccountBalance(dataCheck.Amount, fromAccount.Balance); transferErr == nil {
					return nil
				}
			}
		}
	}

	if transferErr != nil {
		errorTese := fmt.Sprintf(transferErr.Error())
		return errors.New(errorTese)
	}

	// ต้องเช็ค  owner  และ account ว่ามีอยู๋จริงไหม  เช็คจาก accounts
	// เอา id กับ owner มาเช็ค
	// อันนี้ใช้รวมกัน การเช็ค CheckDataFor_Deposit_and_Withdraw *************************
	// from_Account := new(models.CreateAccountParams)
	// tx := rt.db.Table("accounts").Where("owner =?", dataCheck.Owner).Where("id =?", dataCheck.FromAccountID).First(&from_Account)
	// if tx.Error != nil {
	// 	// ไม่เจอบัญชี สำหรับโอนเงิน
	// 	errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "From account")
	// 	return errors.New(errorTese)
	// 	// return tx.Error
	// }

	// เช็ค จำนวนเงินที่จะโอน ของ  from_Account *************
	// if from_Account.Balance < dataCheck.Amount {
	// 	errorTese := fmt.Sprintf("money not enough")
	// 	return errors.New(errorTese)
	// }

	// เช็ค  account ปลายทางด้วยว่ามีจริงไหม เช็คจาก accounts
	// เอา id มาเช็ค
	// to_Account := new(models.CreateAccountParams)
	// tx = rt.db.Table("accounts").Where("id =?", dataCheck.ToAccountID).First(&to_Account)
	// if tx.Error != nil {
	// 	// ไม่เจอ บัญชีปลายทางสำหรับโอน
	// 	errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "To account")
	// 	return errors.New(errorTese)
	// 	// return tx.Error
	// }

	// เช็ค current ว่า ปะเภทเดียวกันรึป่าว ********
	// if from_Account.Currency != to_Account.Currency {
	// 	return errors.New("cruuency not match")
	// }

	return nil
}

// CreateTransfer implements I_Repo_Transfers
func (rt *repo_Transfers) Create_Transfer(dataTransfer *models.CreateTransferParams) error {

	// Check Data for Transfers
	errCheck := rt.CheckDataFor_Transfer(dataTransfer)
	if errCheck != nil {
		return errCheck
	}

	//****************** entries ***************
	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว
	dataEntries := []models.Entry{
		{
			// FromAccountID *********************
			AccountID:    dataTransfer.FromAccountID,
			Amount:       -dataTransfer.Amount,
			CreatedAt:    time.Now(),
			Entries_type: "transfer",
		},
		{
			// ToAccountID  *******************
			AccountID:    dataTransfer.ToAccountID,
			Amount:       dataTransfer.Amount,
			CreatedAt:    time.Now(),
			Entries_type: "transfer",
		},
	}

	// *************** accounts *****************
	dataAccount := []models.UpdateAccountParams{
		{
			// FromAccountID  *******************
			ID:      dataTransfer.FromAccountID,
			Balance: -dataTransfer.Amount,
		},
		{
			// ToAccountID  *******************
			ID:      dataTransfer.ToAccountID,
			Balance: dataTransfer.Amount,
		},
	}

	// ใสหรือไม่ใสก็ได้มั้ง ********
	dataTransfer.CreatedAt = time.Now()

	// from Account ต้อง - blance
	// to Account ต้อง + blance
	// ทำเป็น array  จะได้ create  ครั้งเดียว
	//**************** ทำ TRanscation  ****************
	return rt.Transaction_Postgres(func(p *repo_Transfers) error {
		// Create Trasfer to DB ***********************
		tx := p.db.Table("transfers").Create(&dataTransfer)
		if tx.Error != nil {
			fmt.Println("TRnsfer naja")
			return tx.Error
		}

		// create Entries  to DB   **************************
		tx = p.db.Table("entries").Create(&dataEntries)
		if tx.Error != nil {
			fmt.Println("entries naja")
			return tx.Error
		}

		// update Accounts to DB  ***************************
		errUpdate := p.UpdateAccount_Blance(dataAccount)
		if errUpdate != nil {
			fmt.Println("accounts naja")
			return errUpdate
		}
		fmt.Println("Transfer is complete")
		return nil
	})

}

// Transaction_Postgres implements I_Repo_Transfers
func (rt *repo_Transfers) Transaction_Postgres(fn func(*repo_Transfers) error) error {
	fmt.Println("")
	dbPostgres := rt.db.Begin()

	newRepo := New_Repo_User(dbPostgres)

	// func ต้องได้ struct ถึงจะเรียกใช้ recive Func ของมันได้
	// ภายใน func นี้ถ้ามีการเรียกใช้ query แล้วมี  error ออกมามันจะไม่ commit ให้
	// transcation ทั้งหมดเลยถูกยกเลิก
	err := fn(newRepo)

	// คีร์หลักคือตรงนี้ ถ้าใน การทำ transcation มี error มันจะมาหยุดตรงนี้ และไม่มีถึ
	// การทำงานของ commit
	if err != nil {
		fmt.Println("Error in Tracscation")
		dbPostgres.Rollback()
		return err
	}

	dbPostgres.Commit()
	return nil
}

// GetTransfer implements I_Repo_Transfers
// เช็คการโอนทั้งหมด ของบัญชีนั้นๆ (ตาม account) และเลือกช่วงเวลา
func (rt *repo_Transfers) GetTransfer_ById(idAccount int64, startTime, endTime string) (dataTransfers []models.ListTransfers_ForOwner, err error) {

	// dataTransfers := []models.ListTransfers_ForOwner{}
	tx := rt.db.Table("transfers").Where("from_account_id=?", idAccount).
		Where("created_at between ? and ?", startTime, endTime).Find(&dataTransfers)
	if tx.Error != nil {
		return dataTransfers, tx.Error
	}
	return dataTransfers, nil
}

// เช็คการโอนทั้งหมด ของเราทุกบัญชี (ตาม owner) และเลือกช่วงเวลา
func (rt *repo_Transfers) GetTransfer_ByOwner(ownerAccount, startTime, endTime string) (dataTransfers []models.ListTransfers_ForOwner, err error) {
	tx := rt.db.Table("transfers").Where("owner = ?", ownerAccount).
		Where("created_at between ? and ?", startTime, endTime).
		Find(&dataTransfers)
	if tx.Error != nil {
		return dataTransfers, tx.Error
	}
	return dataTransfers, nil
}

// ListTransfers implements I_Repo_Transfers
func (rt *repo_Transfers) ListStatement(accountId int64) (statement []models.Transfer, err error) {
	tx := rt.db.Table("entries").Where("account_id = ?", accountId).Find(&statement)
	if tx.Error != nil {
		return statement, tx.Error
	}
	return statement, nil
}

// เช็ค account คนที่จะโอน หรือ เช็ค account คนจะฝากและถอน
func (rt *repo_Transfers) checkAccountOwner(owner string, id int64) error {
	from_Account := new(models.CreateAccountParams)
	tx := rt.db.Table("accounts").Where("owner =?", owner).Where("id =?", id).First(&from_Account)
	if tx.Error != nil {
		errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "From account")
		return errors.New(errorTese)
	}
	return nil
}

// เช็ค account คนที่จะรับโอน
func (rt *repo_Transfers) checkAccountTo(id int64) error {
	to_Account := new(models.CreateAccountParams)
	tx := rt.db.Table("accounts").Where("id =?", id).First(&to_Account)
	if tx.Error != nil {
		errorTese := fmt.Sprintf(tx.Error.Error() + " for:" + "To account")
		return errors.New(errorTese)
	}
	return nil
}

// เช็ค สกุลเงิน
func (rt *repo_Transfers) checkCurrency(fromAccount, toAccount string) error {
	// check currency
	if fromAccount != toAccount {
		return errors.New("cruuency not match")
	}
	return nil
}

// เช็ค จำนนเงินคนที่จะโอน
func (rt *repo_Transfers) checkAccountBalance(amount, balance int64) error {
	// เช็คเงินในการโอน
	if balance < amount {
		errorTese := fmt.Sprintf("money not enough")
		return errors.New(errorTese)
	}
	return nil
}
