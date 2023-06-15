package formatter

import "api-dunia-coding/entity"

type MemberFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type MemberLoginFormatter struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func FormatMember(member entity.Member, token string) MemberFormatter {
	formatter := MemberFormatter{
		ID:    member.ID,
		Name:  member.Name,
		Email: member.Email,
		Token: token,
	}

	return formatter
}

func FormatMemberLogin(message string, token string) MemberLoginFormatter {
	formatter := MemberLoginFormatter{
		Message: message,
		Token:   token,
	}

	return formatter
}
