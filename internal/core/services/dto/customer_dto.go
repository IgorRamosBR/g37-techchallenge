package dto

import (
	"fmt"
	"g37-lanchonete/internal/core/domain"
	"strings"

	"github.com/asaskevich/govalidator"
)

type CustomerDTO struct {
	Name  string `json:"name" valid:"length(0|100)~Name length should be less than 100 characters"`
	Email string `json:"email" valid:"email,length(5|100)~Email length should be between 5 and 100 characters"`
	CPF   string `json:"cpf" valid:"cpf"`
}

func (c CustomerDTO) ToCustomer() domain.Customer {
	return domain.Customer{
		Name:  c.Name,
		Cpf:   c.CPF,
		Email: c.Email,
	}
}

func (c CustomerDTO) ValidateCustomer() (bool, error) {
	if _, err := govalidator.ValidateStruct(c); err != nil {
		return false, err
	}

	// Validate CPF using a custom function
	if !isValidCPF(c.CPF) {
		return false, fmt.Errorf("invalid CPF [%s]", c.CPF)
	}

	return true, nil
}

func isValidCPF(cpf string) bool {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if len(cpf) != 11 {
		return false
	}

	if strings.Count(cpf, string(cpf[0])) == 11 {
		return false
	}

	// Check if all digits are the same
	if cpf == "00000000000" {
		return false
	}

	// Validate CPF using the standard algorithm
	var sum1, sum2 int
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum1 += digit * (10 - i)
		sum2 += digit * (11 - i)
	}

	sum1 %= 11
	if sum1 < 2 {
		sum1 = 0
	} else {
		sum1 = 11 - sum1
	}

	sum2 += sum1 * 2
	sum2 %= 11
	if sum2 < 2 {
		sum2 = 0
	} else {
		sum2 = 11 - sum2
	}

	return cpf[9]-'0' == byte(sum1) && cpf[10]-'0' == byte(sum2)
}
