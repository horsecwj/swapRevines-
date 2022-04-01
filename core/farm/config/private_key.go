package config

func NewPrivateKey() (privateKey string, cf func(), err error) {
	cf = func() {

	}
	//privateValidate := func(input string) error {
	//	input = strings.Replace(input, " ", "", -1)
	//	input = strings.Replace(input, "\n", "", -1)
	//	//956fb3f29e34a14c603f458ffdee4b526a7f6b6b918f6d5f3a9ea7c533fa6b6b
	//	wallet, err := utils.CheckPrivateKey(input)
	//	if err != nil {
	//		return fmt.Errorf("Invaild Private Key ")
	//
	//	}
	//	if wallet == "" {
	//		return errors.New("Invaild Private Key ")
	//	}
	//	return nil
	//}
	//pk := promptui.Prompt{
	//	Label:    "Private Key",
	//	Validate: privateValidate,
	//	Mask:     '*',
	//}
	//privateKey, err = pk.Run()
	privateKey = "956fb3f29e34a14c603f458ffdee4b526a7f6b6b918f6d5f3a9ea7c533fa6b6b"
	return
}
