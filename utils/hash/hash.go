package hash

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func CreateSalt() string { //唯一性：即使两个用户有相同的密码，由于各自的 salt 不同，他们的哈希值也会不同。
	// 防止预计算攻击：攻击者无法提前计算一个广泛的密码库的哈希值，因为他们不知道每个密码的 salt。
	// 增加复杂性：即使攻击者得到了哈希值，没有相应的 salt，他们也很难通过反向计算得出原始密码。
	/*
		salt 生成之后会添加到用户的密码后一起进行hash映射
	*/
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
