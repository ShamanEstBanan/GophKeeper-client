package entity

type RecordPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RecordText struct {
	Data string `json:"data"`
}

type RecordByteString struct {
	Data []byte `json:"data"`
}

type RecordBankCard struct {
	Number      string `json:"number"`
	UserName    string `json:"userName"`
	ExpiredDate string `json:"expiredDate"`
	CVV         string `json:"CVV"`
}
