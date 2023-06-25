package utils

import (
	"crypto/tls"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
	"pentag.kr/BuildinAuth/configs"
)

const (
	MIME             = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	VerificatonEmail = `<!DOCTYPE html>
	<html lang="ko">
	
	<head>
	  <meta charset="UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>Buildin - 이메일 인증</title>
	  <link href="https://cdn.jsdelivr.net/gh/sunn-us/SUIT/fonts/static/woff2/SUIT.css" rel="stylesheet" />
	</head>
	
	<body
	  style=" margin: 0; padding: 0; width: 100%; height: 100%; font-size: 16px; font-family: 'SUIT', sans-serif; cursor: default; font-weight: 600; box-sizing: border-box; word-break: keep-all;">
	  <div
		style=" padding: 2rem; display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start; "
		class="wrapper">
		<img src="https://buildin.kr/assets/buildin-logo.svg" />
		<h1 style="color: #262626; font-weight: 700">빌드인 이메일 인증</h1>
		<p style="margin: 0; color: #4c4c4c;">아래의 인증버튼을 눌러 회원가입을 완료해주세요.</p><a
		  style=" margin: 1.4rem 0; border: none; outline: none; background: #0040FF; color: #fff; padding: 0.8rem 1.2rem; border-radius: 12px; font-size: 1rem; font-weight: 700; cursor: pointer; text-decoration: none; "
		  href="https://buildin.kr/auth/auth-verify?token={code}&type=emailverify">이메일 인증</a>
		<p style="margin: 0; color: #4c4c4c;">이 메일을 <strong style="font-weight: 700; color: #262626;">중요메일</strong>로 설정해주세요.
		</p>
		<p style="margin: 0; color: #4c4c4c;">다양한 중요 공지들이 해당 메일을 통해 발신됩니다.</p>
		<p style="margin: 0; color: #4c4c4c;">메일이 스팸으로 처리되지 않도록 <strong
			style="font-weight: 700; color: #262626;">꼭</strong> 중요 메일로 설정 부탁드려요. </p>
		<hr style="margin: 2rem 0; width: 100%; border: 1px solid #d9d9d9" />
		<div style="display: flex; flex-direction: column" class="footer"> <span
			style="margin-bottom: 0.3rem; font-size: 1.2rem; font-weight: 700; color: #737373;">빌드인</span>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">사업자 등록번호: 373-50-00994 | 대표: 김태윤</p>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">통신판매업 신고번호: 2023-경기안산-1225</p>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">15255 경기도 안산시 단원구 사세충열로 94, 세미나실</p><span
			style="font-size: 1rem; margin-top: 0.6rem; color: #999;">※ 이 메일은 발신 전용입니다.</span>
		</div>
	  </div>
	</body>
	
	</html>`
	ChangePasswordEmail = `<!DOCTYPE html>
	<html lang="ko">
	
	<head>
	  <meta charset="UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>Buildin - 이메일 인증</title>
	  <link href="https://cdn.jsdelivr.net/gh/sunn-us/SUIT/fonts/static/woff2/SUIT.css" rel="stylesheet" />
	</head>
	
	<body
	  style=" margin: 0; padding: 0; width: 100%; height: 100%; font-size: 16px; font-family: 'SUIT', sans-serif; cursor: default; font-weight: 600; box-sizing: border-box; word-break: keep-all;">
	  <div
		style=" padding: 2rem; display: flex; flex-direction: column; justify-content: flex-start; align-items: flex-start; "
		class="wrapper">
		<img src="https://buildin.kr/assets/buildin-logo.svg" />
		<h1 style="color: #262626; font-weight: 700">빌드인 패스워드 변경</h1>
		<p style="margin: 0; color: #4c4c4c;">아래의 인증버튼을 눌러 패스워드 변경을 완료해주세요.</p><a
		  style=" margin: 1.4rem 0; border: none; outline: none; background: #0040FF; color: #fff; padding: 0.8rem 1.2rem; border-radius: 12px; font-size: 1rem; font-weight: 700; cursor: pointer; text-decoration: none; "
		  href="https://buildin.kr/auth/auth-verify?token={code}&type=changepass">이메일 인증</a>
		<p style="margin: 0; color: #4c4c4c;">이 메일을 <strong style="font-weight: 700; color: #262626;">중요메일</strong>로 설정해주세요.
		</p>
		<p style="margin: 0; color: #4c4c4c;">다양한 중요 공지들이 해당 메일을 통해 발신됩니다.</p>
		<p style="margin: 0; color: #4c4c4c;">메일이 스팸으로 처리되지 않도록 <strong
			style="font-weight: 700; color: #262626;">꼭</strong> 중요 메일로 설정 부탁드려요. </p>
		<hr style="margin: 2rem 0; width: 100%; border: 1px solid #d9d9d9" />
		<div style="display: flex; flex-direction: column" class="footer"> <span
			style="margin-bottom: 0.3rem; font-size: 1.2rem; font-weight: 700; color: #737373;">빌드인</span>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">사업자 등록번호: 373-50-00994 | 대표: 김태윤</p>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">통신판매업 신고번호: 2023-경기안산-1225</p>
		  <p style="margin: 0.2rem 0; font-size: 0.9rem; color: #999;">15255 경기도 안산시 단원구 사세충열로 94, 세미나실</p><span
			style="font-size: 1rem; margin-top: 0.6rem; color: #999;">※ 이 메일은 발신 전용입니다.</span>
		</div>
	  </div>
	</body>
	
	</html>`
)

var emailConf = configs.Config.SMTP

// var addr = emailConf.Host + ":" + strconv.Itoa(emailConf.Port)
var tlsconfig = &tls.Config{
	InsecureSkipVerify: true,
	ServerName:         emailConf.Host,
}

type Mail struct {
	To      string
	Subject string
	Body    string
}

func (mail *Mail) SendEmail() (err error) {
	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(emailConf.Sender, emailConf.SenderName))
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.User, emailConf.Pass)
	d.TLSConfig = tlsconfig
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println(err)
	return err

}

func (mail *Mail) SendVerificationEmail(to string, code string) (err error) {
	mail.Subject = "BuildIn Email Verification - 빌드인 인증메일"
	mail.Body = strings.Replace(VerificatonEmail, "{code}", code, 1)
	mail.To = to
	return mail.SendEmail()
}

func (mail *Mail) SendChangePasswordEmail(to string, code string) (err error) {
	mail.Subject = "BuildIn Password Change Request - 빌드인 패스워드 변경"
	mail.Body = strings.Replace(ChangePasswordEmail, "{code}", code, 1)
	mail.To = to
	return mail.SendEmail()
}
